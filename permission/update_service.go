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

package permission

import (
	"context"

	objectstorage "github.com/sacloud/object-storage-api-go"
	v1 "github.com/sacloud/object-storage-api-go/apis/v1"
)

func (s *Service) Update(req *UpdateRequest) (*v1.Permission, error) {
	return s.UpdateWithContext(context.Background(), req)
}

func (s *Service) UpdateWithContext(ctx context.Context, req *UpdateRequest) (*v1.Permission, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	client := objectstorage.NewPermissionOp(s.client)

	current, err := client.Read(ctx, req.SiteId, req.Id)
	if err != nil {
		return nil, err
	}

	return client.Update(ctx, req.SiteId, req.Id, req.ToRequestParameter(current))
}
