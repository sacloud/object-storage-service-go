# object-storage-service-go

[![Go Reference](https://pkg.go.dev/badge/github.com/sacloud/object-storage-service-go.svg)](https://pkg.go.dev/github.com/sacloud/object-storage-service-go)
[![Tests](https://github.com/sacloud/object-storage-service-go/workflows/Tests/badge.svg)](https://github.com/sacloud/object-storage-service-go/actions/workflows/tests.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/sacloud/object-storage-service-go)](https://goreportcard.com/report/github.com/sacloud/object-storage-service-go)


さくらのオブジェクトストレージ高レベルAPIライブラリ  

## 概要

さくらのオブジェクトストレージAPIをラップし、CRUD+L操作を統一的な手順で行えるインターフェースを提供します。  

インターフェースの例:
```go
// パーミッション操作の例
func (s *Service) Create(req *CreateRequest) (*v1.Permission, error)
func (s *Service) CreateWithContext(ctx context.Context, req *CreateRequest) (*v1.Permission, error)

func (s *Service) Read(req *ReadRequest) (*v1.Permission, error)
func (s *Service) ReadWithContext(ctx context.Context, req *ReadRequest) (*v1.Permission, error)

func (s *Service) Update(req *UpdateRequest) (*v1.Permission, error)
func (s *Service) UpdateWithContext(ctx context.Context, req *UpdateRequest) (*v1.Permission, error)

func (s *Service) Delete(req *DeleteRequest) error
func (s *Service) DeleteWithContext(ctx context.Context, req *DeleteRequest) error

func (s *Service) Find(req *FindRequest) ([]*v1.Permission, error)
func (s *Service) FindWithContext(ctx context.Context, req *FindRequest) ([]*v1.Permission, error)
```

以下のリソースに対応しています。

```console
.
├── account
│   └── accesskey
├── bucket
├── permission
│   ├── accesskey
│   └── bucketcontrol
└── site
    └── status
```

## License

`sacloud/object-storage-service-go` Copyright (C) 2022 [The sacloud/object-storage-service-go Authors](AUTHORS).

This project is published under [Apache 2.0 License](LICENSE.txt).
