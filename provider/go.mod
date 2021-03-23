module github.com/pulumi/pulumi-equinix-metal/provider

go 1.16

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/hashicorp/go-getter v1.5.0 => github.com/hashicorp/go-getter v1.4.0
)

require (
	contrib.go.opencensus.io/resource v0.1.1 // indirect
	github.com/equinix/terraform-provider-metal v1.0.1-0.20201209021325-714995571fa3
	github.com/googleapis/gax-go v2.0.2+incompatible // indirect
	github.com/hashicorp/terraform-plugin-sdk v1.9.1
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.22.1
	github.com/pulumi/pulumi/pkg/v2 v2.22.1-0.20210310211618-1f16423ede4c
	github.com/pulumi/pulumi/sdk/v2 v2.22.1-0.20210310211618-1f16423ede4c
)
