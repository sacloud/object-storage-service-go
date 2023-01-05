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

package permission

import (
	v1 "github.com/sacloud/object-storage-api-go/apis/v1"
	"github.com/sacloud/packages-go/validate"
)

type CreateRequest struct {
	SiteId         string           `service:"-" validate:"required"`
	DisplayName    string           `validate:"required"`
	BucketControls []*BucketControl `validate:"omitempty,dive,required"`
}

func (req *CreateRequest) Validate() error {
	return validate.New().Struct(req)
}

func (req *CreateRequest) ToRequestParameter() *v1.CreatePermissionParams {
	p := &v1.CreatePermissionParams{
		BucketControls: v1.BucketControls{},
		DisplayName:    v1.DisplayName(req.DisplayName),
	}
	for _, bc := range req.BucketControls {
		p.BucketControls = append(p.BucketControls, v1.BucketControl{
			BucketName: v1.BucketName(bc.BucketName),
			CanRead:    v1.CanRead(bc.CanRead),
			CanWrite:   v1.CanWrite(bc.CanWrite),
		})
	}
	return p
}
