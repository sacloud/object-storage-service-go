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

package bucketcontrol

import (
	"context"
	"fmt"
	"net/http"

	objectstorage "github.com/sacloud/object-storage-api-go"
	v1 "github.com/sacloud/object-storage-api-go/apis/v1"
)

func (s *Service) Create(req *CreateRequest) (*v1.BucketControl, error) {
	return s.CreateWithContext(context.Background(), req)
}

func (s *Service) CreateWithContext(ctx context.Context, req *CreateRequest) (*v1.BucketControl, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	client := objectstorage.NewPermissionOp(s.client)
	permission, err := client.Read(ctx, req.SiteId, req.PermissionId)
	if err != nil {
		return nil, err
	}

	for _, bc := range permission.BucketControls {
		if bc.BucketName.String() == req.BucketName {
			return nil, &v1.Error409{
				Detail: v1.ErrorDetail{
					Code:    http.StatusConflict,
					Message: v1.ErrorMessage(fmt.Sprintf("bucket control for %q already exist", req.BucketName)),
				},
			}
		}
	}

	permission.BucketControls = append(permission.BucketControls, v1.BucketControl{
		BucketName: v1.BucketName(req.BucketName),
		CanRead:    v1.CanRead(req.CanRead),
		CanWrite:   v1.CanWrite(req.CanWrite),
	})
	permission, err = client.Update(ctx, req.SiteId, req.PermissionId, &v1.UpdatePermissionParams{
		BucketControls: permission.BucketControls,
		DisplayName:    permission.DisplayName,
	})
	if err != nil {
		return nil, err
	}

	for _, bc := range permission.BucketControls {
		if bc.BucketName.String() == req.BucketName {
			return &bc, nil
		}
	}

	// 到達しないはず(さくらのクラウド側の障害というケースがあり得るためpanicはさせない)
	return nil, fmt.Errorf("created bucket-control not found")
}
