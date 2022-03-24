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

package sitestatus

import (
	"net/http/httptest"
	"testing"

	objectstorage "github.com/sacloud/object-storage-api-go"
	v1 "github.com/sacloud/object-storage-api-go/apis/v1"
	"github.com/sacloud/object-storage-api-go/fake"
	"github.com/sacloud/object-storage-api-go/fake/server"
	service "github.com/sacloud/object-storage-service-go"
	"github.com/stretchr/testify/require"
)

var siteId = "isk01"

func TestService_CRUD_plus_L(t *testing.T) {
	server := initFakeServer()
	client := &objectstorage.Client{
		APIRootURL: server.URL,
	}
	svc := New(client)

	t.Run("read", func(t *testing.T) {
		read, err := svc.Read(&ReadRequest{
			Id: siteId,
		})
		require.NoError(t, err)
		require.NotNil(t, read)
	})

	t.Run("read return NotFoundError when site is not found", func(t *testing.T) {
		id := "not-exists-site-id"
		read, err := svc.Read(&ReadRequest{
			Id: id,
		})
		require.Nil(t, read)
		require.EqualError(t, err, `site "`+id+`" not found`)

		_, errIsNotFoundError := err.(service.NotFoundError)
		require.True(t, errIsNotFoundError)
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
