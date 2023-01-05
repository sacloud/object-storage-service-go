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

	objectstorage "github.com/sacloud/object-storage-api-go"
)

func (s *Service) Create(req *CreateRequest) (*Bucket, error) {
	return s.CreateWithContext(context.Background(), req)
}

func (s *Service) CreateWithContext(ctx context.Context, req *CreateRequest) (*Bucket, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	bucketOp := objectstorage.NewBucketOp(s.client)
	_, err := bucketOp.Create(ctx, req.SiteId, req.Id)
	if err != nil {
		return nil, err
	}
	return s.ReadWithContext(ctx, &ReadRequest{
		AccessKey: req.AccessKey,
		SecretKey: req.SecretKey,
		SiteId:    req.SiteId,
		Id:        req.Id,
	})
}
