module github.com/aidtechnology/affinityctl

go 1.14

require (
	github.com/golang/protobuf v1.4.2
	github.com/grpc-ecosystem/grpc-gateway v1.14.8
	github.com/mitchellh/go-homedir v1.1.0
	github.com/spf13/cobra v1.0.0
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.6.1
	go.bryk.io/x v0.0.0-20200917170803-d4792ef6c59f
	google.golang.org/genproto v0.0.0-20200911024640-645f7a48b24f
	google.golang.org/grpc v1.31.1
	google.golang.org/protobuf v1.25.0
)

replace github.com/cloudflare/cfssl => github.com/bryk-io/cfssl v0.0.0-20191204191638-bb9c164a4cb1
