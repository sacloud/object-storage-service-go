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

package accesskey

import (
	"context"

	objectstorage "github.com/sacloud/object-storage-api-go"
)

func (s *Service) Delete(req *DeleteRequest) error {
	return s.DeleteWithContext(context.Background(), req)
}

func (s *Service) DeleteWithContext(ctx context.Context, req *DeleteRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}
	_, err := s.ReadWithContext(ctx, &ReadRequest{
		SiteId: req.SiteId,
		Id:     req.Id,
	})
	if err != nil {
		return err
	}

	client := objectstorage.NewAccountOp(s.client)
	return client.DeleteAccessKey(ctx, req.SiteId, req.Id)
}
