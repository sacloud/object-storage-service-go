// Copyright 2022-2023 The sacloud/object-storage-service-go Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package permission

import (
	"net/http/httptest"
	"testing"

	client "github.com/sacloud/api-client-go"
	objectstorage "github.com/sacloud/object-storage-api-go"
	v1 "github.com/sacloud/object-storage-api-go/apis/v1"
	"github.com/sacloud/object-storage-api-go/fake"
	"github.com/sacloud/object-storage-api-go/fake/server"
	service "github.com/sacloud/object-storage-service-go"
	"github.com/sacloud/packages-go/pointer"
	"github.com/sacloud/packages-go/testutil"
	"github.com/stretchr/testify/require"
)

var siteId = "isk01"

func TestPermission_CRUD_plus_L(t *testing.T) {
	server := initFakeServer()
	client := &objectstorage.Client{
		APIRootURL: server.URL,
		Options: &client.Options{
			UserAgent: service.UserAgent,
		},
	}
	svc := New(client)
	var permission *v1.Permission

	t.Run("create", func(t *testing.T) {
		created, err := svc.Create(&CreateRequest{
			SiteId:      siteId,
			DisplayName: testutil.Random(16, testutil.CharSetAlpha),
		})
		require.NoError(t, err)
		require.NotNil(t, created)
		permission = created
	})

	t.Run("read", func(t *testing.T) {
		read, err := svc.Read(&ReadRequest{
			SiteId: siteId,
			Id:     permission.Id.Int64(),
		})
		require.NoError(t, err)
		require.NotNil(t, read)
	})

	t.Run("read return NotFoundError when account is not found", func(t *testing.T) {
		id := int64(123456789012)
		read, err := svc.Read(&ReadRequest{
			SiteId: siteId,
			Id:     id,
		})
		require.Nil(t, read)
		require.Error(t, err)
		require.True(t, v1.IsError404(err))
	})

	t.Run("list", func(t *testing.T) {
		found, err := svc.Find(&FindRequest{
			SiteId: siteId,
		})
		require.NoError(t, err)
		require.Len(t, found, 1)

		require.Equal(t, permission, found[0])
	})

	t.Run("update", func(t *testing.T) {
		updatedName := permission.DisplayName.String() + "-updated"
		updated, err := svc.Update(&UpdateRequest{
			SiteId:      siteId,
			Id:          permission.Id.Int64(),
			DisplayName: pointer.NewString(updatedName),
		})

		require.NoError(t, err)
		require.Equal(t, updatedName, updated.DisplayName.String())
		require.Equal(t, permission.BucketControls, updated.BucketControls)
	})

	t.Run("delete return NotFoundError when account is not found", func(t *testing.T) {
		id := int64(123456789012)
		err := svc.Delete(&DeleteRequest{
			SiteId: siteId,
			Id:     id,
		})

		require.Error(t, err)
		require.True(t, v1.IsError404(err))
	})

	t.Run("delete", func(t *testing.T) {
		err := svc.Delete(&DeleteRequest{
			SiteId: siteId,
			Id:     permission.Id.Int64(),
		})
		require.NoError(t, err)

		found, err := svc.Find(&FindRequest{
			SiteId: siteId,
		})
		require.NoError(t, err)
		require.Len(t, found, 0)
	})
}

func initFakeServer() *httptest.Server {
	fakeServer := &server.Server{
		Engine: &fake.Engine{
			Clusters: []*v1.Cluster{
				{
					Id:              siteId,
					ControlPanelUrl: "https://secure.sakura.ad.jp/objectstorage/",
					DisplayNameEnUs: "Ishikari Site #1",
					DisplayNameJa:   "石狩第1サイト",
					DisplayName:     "石狩第1サイト",
					DisplayOrder:    1,
					EndpointBase:    "isk01.sakurastorage.jp",
				},
			},
		},
	}
	return httptest.NewServer(fakeServer.Handler())
}
