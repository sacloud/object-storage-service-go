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

package bucket

import (
	"context"
	"fmt"

	service "github.com/sacloud/object-storage-service-go"
)

// Read バケットの参照
//
// 詳細はReadWithContextのコメントを参照してください
func (s *Service) Read(req *ReadRequest) (*Bucket, error) {
	return s.ReadWithContext(context.Background(), req)
}

// ReadWithContext バケットの参照
//
// 指定のId(バケット名)を持つバケットが見つからなかった場合はNotFoundErrorを返す
func (s *Service) ReadWithContext(ctx context.Context, req *ReadRequest) (*Bucket, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	findReq := &FindRequest{
		AccessKey: req.AccessKey,
		SecretKey: req.SecretKey,
		SiteId:    req.SiteId,
	}
	buckets, err := s.FindWithContext(ctx, findReq)
	if err != nil {
		return nil, err
	}

	for _, bucket := range buckets {
		if bucket.Name == req.Id {
			return bucket, nil
		}
	}
	var e error = service.NotFoundError(fmt.Errorf("bucket %q not found", req.Id))
	if _, ok := e.(service.NotFoundError); ok {
		fmt.Println("foo")
	}
	return nil, service.NotFoundError(fmt.Errorf("bucket %q not found", req.Id))
}
