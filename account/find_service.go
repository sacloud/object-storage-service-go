// Copyright 2022-2025 The sacloud/object-storage-service-go Authors
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

package account

import (
	"context"

	objectstorage "github.com/sacloud/object-storage-api-go"
	v1 "github.com/sacloud/object-storage-api-go/apis/v1"
)

func (s *Service) Find(req *FindRequest) ([]*v1.Account, error) {
	return s.FindWithContext(context.Background(), req)
}

func (s *Service) FindWithContext(ctx context.Context, req *FindRequest) ([]*v1.Account, error) {
	if req == nil {
		req = &FindRequest{}
	}
	if err := req.Validate(); err != nil {
		return nil, err
	}

	client := objectstorage.NewAccountOp(s.client)
	account, err := client.Read(ctx, req.SiteId)
	if err != nil {
		// Note: serviceインターフェースを満たすためにclient.Read()が404エラーを返した時はエラーではなく[]*v1.Account{}を返す
		if _, ok := err.(*v1.Error404); ok {
			return []*v1.Account{}, nil
		}
		return nil, err
	}
	return []*v1.Account{account}, nil
}
