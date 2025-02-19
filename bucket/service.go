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

package bucket

import objectstorage "github.com/sacloud/object-storage-api-go"

// Service provides a high-level API of for Site
type Service struct {
	client *objectstorage.Client
}

// New returns new service instance of Archive
func New(client *objectstorage.Client) *Service {
	return &Service{client: client}
}
