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

package site

import (
	"net/http/httptest"
	"testing"

	client "github.com/sacloud/api-client-go"
	objectstorage "github.com/sacloud/object-storage-api-go"
	v1 "github.com/sacloud/object-storage-api-go/apis/v1"
	"github.com/sacloud/object-storage-api-go/fake"
	"github.com/sacloud/object-storage-api-go/fake/server"
	service "github.com/sacloud/object-storage-service-go"
	"github.com/stretchr/testify/require"
)

var siteId = "isk01"

func TestSite_CRUD_plus_L(t *testing.T) {
	server := initFakeServer()
	client := &objectstorage.Client{
		APIRootURL: server.URL,
		Options: &client.Options{
			UserAgent: service.UserAgent,
		},
	}
	svc := New(client)
	var site *v1.Cluster

	t.Run("read", func(t *testing.T) {
		read, err := svc.Read(&ReadRequest{
			Id: siteId,
		})
		require.NoError(t, err)
		require.NotNil(t, read)

		site = read
	})

	t.Run("read return NotFoundError when site is not found", func(t *testing.T) {
		id := "not-exists-site-id"
		read, err := svc.Read(&ReadRequest{
			Id: id,
		})
		require.Nil(t, read)
		require.Error(t, err)
		require.True(t, v1.IsError404(err))
	})

	t.Run("list", func(t *testing.T) {
		found, err := svc.Find(&FindRequest{})
		require.NoError(t, err)
		require.Len(t, found, 1)

		require.Equal(t, site, found[0])
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
