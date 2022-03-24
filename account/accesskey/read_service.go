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
	"context"
	"fmt"

	objectstorage "github.com/sacloud/object-storage-api-go"
	v1 "github.com/sacloud/object-storage-api-go/apis/v1"
	service "github.com/sacloud/object-storage-service-go"
)

func (s *Service) Read(req *ReadRequest) (*v1.AccountKey, error) {
	return s.ReadWithContext(context.Background(), req)
}

func (s *Service) ReadWithContext(ctx context.Context, req *ReadRequest) (*v1.AccountKey, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	client := objectstorage.NewAccountOp(s.client)
	accountKey, err := client.ReadAccessKey(ctx, req.SiteId, req.Id)
	if err != nil {
		if v1.IsError404(err) {
			return nil, service.NotFoundError(fmt.Errorf("account-key %q not found", req.Id))
		}
		return nil, err
	}
	return accountKey, nil
}
