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

//go:build acctest
// +build acctest

package bucket

import (
	"os"
	"testing"

	v1 "github.com/sacloud/object-storage-api-go/apis/v1"
	service "github.com/sacloud/object-storage-service-go"
	"github.com/sacloud/packages-go/testutil"
	"github.com/stretchr/testify/require"
)

func TestAccBucket_CRUD_plus_L(t *testing.T) {
	testutil.PreCheckEnvsFunc(
		"SAKURACLOUD_ACCESS_TOKEN",
		"SAKURACLOUD_ACCESS_TOKEN_SECRET",
		"SACLOUD_OJS_ACCESS_KEY_ID",
		"SACLOUD_OJS_SECRET_ACCESS_KEY",
	)(t)
	key := os.Getenv("SACLOUD_OJS_ACCESS_KEY_ID")
	secret := os.Getenv("SACLOUD_OJS_SECRET_ACCESS_KEY")

	svc := New(service.NewClient())
	siteId := "isk01"
	bucketName := testutil.RandomName("object-storage-service-go-", 16, testutil.CharSetAlpha)
	notExistName := testutil.RandomName("object-storage-service-go-", 16, testutil.CharSetAlpha)

	t.Run("create", func(t *testing.T) {
		bucket, err := svc.Create(&CreateRequest{
			AccessKey: key,
			SecretKey: secret,
			SiteId:    siteId,
			Id:        bucketName,
		})
		require.NoError(t, err)
		require.NotNil(t, bucket)
		t.Logf("created: name: %s, creation-date: %s", bucket.Name, bucket.CreationDate)
	})

	t.Run("list and read", func(t *testing.T) {
		// Note: このサービスのReadは内部でListを読んでいるため、ここではReadのみ実施している
		bucket, err := svc.Read(&ReadRequest{
			AccessKey: key,
			SecretKey: secret,
			SiteId:    siteId,
			Id:        bucketName,
		})
		require.NoError(t, err)
		require.NotNil(t, bucket)
		t.Logf("read: name: %s, creation-date: %s", bucket.Name, bucket.CreationDate)
	})

	t.Run("read return NotFoundError when bucket is not found", func(t *testing.T) {
		_, err := svc.Read(&ReadRequest{
			AccessKey: key,
			SecretKey: secret,
			SiteId:    siteId,
			Id:        notExistName,
		})
		require.Error(t, err)
		require.True(t, v1.IsError404(err))
	})

	t.Run("delete", func(t *testing.T) {
		err := svc.Delete(&DeleteRequest{
			AccessKey: key,
			SecretKey: secret,
			SiteId:    siteId,
			Id:        bucketName,
		})
		require.NoError(t, err)
	})
}
