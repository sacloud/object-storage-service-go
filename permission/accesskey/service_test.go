// Copyright 2022 The sacloud/object-storage-service-go Authors
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

package accesskey

import (
	"net/http/httptest"
	"testing"

	objectstorage "github.com/sacloud/object-storage-api-go"
	v1 "github.com/sacloud/object-storage-api-go/apis/v1"
	"github.com/sacloud/object-storage-api-go/fake"
	"github.com/sacloud/object-storage-api-go/fake/server"
	"github.com/stretchr/testify/require"
)

var (
	siteId       = "isk01"
	permissionId = int64(100000000001)
)

func TestPermissionAccessKey_CRUD_plus_L(t *testing.T) {
	fakeServer := initFakeServer()
	client := &objectstorage.Client{
		APIRootURL: fakeServer.URL,
	}

	svc := New(client)
	var permissionKey *v1.PermissionKey

	t.Run("create", func(t *testing.T) {
		created, err := svc.Create(&CreateRequest{
			SiteId:       siteId,
			PermissionId: permissionId,
		})
		require.NoError(t, err)
		require.NotNil(t, created)
		permissionKey = created
	})

	t.Run("read", func(t *testing.T) {
		read, err := svc.Read(&ReadRequest{
			SiteId:       siteId,
			PermissionId: permissionId,
			Id:           permissionKey.Id.String(),
		})
		require.NoError(t, err)
		require.NotNil(t, read)
	})

	t.Run("read return NotFoundError when account is not found", func(t *testing.T) {
		id := "not-exists-account-id"
		read, err := svc.Read(&ReadRequest{
			SiteId:       siteId,
			PermissionId: permissionId,
			Id:           id,
		})
		require.Nil(t, read)
		require.Error(t, err)
		require.True(t, v1.IsError404(err))
	})

	t.Run("list", func(t *testing.T) {
		found, err := svc.Find(&FindRequest{
			SiteId:       siteId,
			PermissionId: permissionId,
		})
		require.NoError(t, err)
		require.Len(t, found, 1)

		// Note: AccountKey.SecretKeyは作成時のみ参照可能なのでそれ以外の項目を比較する
		key := *permissionKey
		key.Secret = ""
		require.Equal(t, &key, found[0])
	})

	t.Run("delete return NotFoundError when account is not found", func(t *testing.T) {
		id := "not-exists-account-id"
		err := svc.Delete(&DeleteRequest{
			SiteId:       siteId,
			PermissionId: permissionId,
			Id:           id,
		})
		require.Error(t, err)
		require.True(t, v1.IsError404(err))
	})

	t.Run("delete", func(t *testing.T) {
		err := svc.Delete(&DeleteRequest{
			SiteId:       siteId,
			PermissionId: permissionId,
			Id:           permissionKey.Id.String(),
		})
		require.NoError(t, err)

		found, err := svc.Find(&FindRequest{
			SiteId:       siteId,
			PermissionId: permissionId,
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
			Permissions: []*v1.Permission{
				{
					DisplayName: "foobar",
					Id:          v1.PermissionID(permissionId),
				},
			},
		},
	}
	return httptest.NewServer(fakeServer.Handler())
}
