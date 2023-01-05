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

package bucket

import (
	"context"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func s3Client(ctx context.Context, s3Endpoint, key, secret string) (*minio.Client, error) {
	return minio.New(s3Endpoint, &minio.Options{
		Creds:        credentials.NewStaticV4(key, secret, ""),
		Region:       "jp-north-1",
		Secure:       true,
		BucketLookup: minio.BucketLookupPath,
	})
}
