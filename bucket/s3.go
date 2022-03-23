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

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func s3Client(ctx context.Context, s3Endpoint, key, secret string) (*s3.Client, error) {
	cred := credentials.NewStaticCredentialsProvider(key, secret, "")
	endpoint := aws.EndpointResolverWithOptionsFunc(
		func(service, region string, options ...interface{}) (aws.Endpoint, error) {
			return aws.Endpoint{
				URL:               "https://" + s3Endpoint,
				HostnameImmutable: true,
				SigningRegion:     region,
			}, nil
		},
	)
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion("jp-north-1"),
		config.WithCredentialsProvider(cred),
		config.WithEndpointResolverWithOptions(endpoint),
	)
	if err != nil {
		return nil, err
	}
	return s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true
	}), nil
}
