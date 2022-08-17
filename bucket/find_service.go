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

	objectstorage "github.com/sacloud/object-storage-api-go"
)

func (s *Service) Find(req *FindRequest) ([]*Bucket, error) {
	return s.FindWithContext(context.Background(), req)
}

func (s *Service) FindWithContext(ctx context.Context, req *FindRequest) ([]*Bucket, error) {
	if req == nil {
		req = &FindRequest{}
	}
	if err := req.Validate(); err != nil {
		return nil, err
	}

	siteOp := objectstorage.NewSiteOp(s.client)
	site, err := siteOp.Read(ctx, req.SiteId)
	if err != nil {
		return nil, err
	}

	s3Client, err := s3Client(ctx, site.S3Endpoint, req.AccessKey, req.SecretKey)
	if err != nil {
		return nil, err
	}

	outputs, err := s3Client.ListBuckets(ctx)
	if err != nil {
		return nil, err
	}

	var results []*Bucket
	for i := range outputs {
		results = append(results, &Bucket{Name: outputs[i].Name, CreationDate: &outputs[i].CreationDate})
	}
	return results, nil
}
