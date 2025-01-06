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

func (s *Service) Delete(req *DeleteRequest) error {
	return s.DeleteWithContext(context.Background(), req)
}

func (s *Service) DeleteWithContext(ctx context.Context, req *DeleteRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}
	client := objectstorage.NewPermissionOp(s.client)
	permission, err := client.Read(ctx, req.SiteId, req.PermissionId)
	if err != nil {
		return err
	}

	var bucketControls v1.BucketControls
	for _, bc := range permission.BucketControls {
		if bc.BucketName.String() != req.BucketName {
			bucketControls = append(bucketControls, bc)
		}
	}

	if len(bucketControls) == len(permission.BucketControls) {
		return &v1.Error404{
			Detail: v1.ErrorDetail{
				Code:    http.StatusNotFound,
				Message: v1.ErrorMessage(fmt.Sprintf("bucket control for %q not found", req.BucketName)),
			},
		}
	}
	permission.BucketControls = bucketControls

	_, err = client.Update(ctx, req.SiteId, req.PermissionId, &v1.UpdatePermissionParams{
		BucketControls: permission.BucketControls,
		DisplayName:    permission.DisplayName,
	})
	return err
}
