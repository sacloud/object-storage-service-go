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

import "testing"

func TestCreateRequest_Validate(t *testing.T) {
	type fields struct {
		SiteId         string
		DisplayName    string
		BucketControls []*BucketControl
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "minimum",
			fields: fields{
				SiteId:         "isk01",
				DisplayName:    "minimum",
				BucketControls: nil,
			},
			wantErr: false,
		},
		{
			name: "with bucket controls",
			fields: fields{
				SiteId:      "isk01",
				DisplayName: "with bucket controls",
				BucketControls: []*BucketControl{
					{BucketName: "bucket"},
				},
			},
			wantErr: false,
		},
		{
			name: "invalid bucket controls",
			fields: fields{
				SiteId:      "isk01",
				DisplayName: "invalid bucket controls",
				BucketControls: []*BucketControl{
					nil,
					{BucketName: "bucket"},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &CreateRequest{
				SiteId:         tt.fields.SiteId,
				DisplayName:    tt.fields.DisplayName,
				BucketControls: tt.fields.BucketControls,
			}
			if err := req.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
