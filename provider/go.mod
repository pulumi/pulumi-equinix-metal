module github.com/pulumi/pulumi-equinix-metal/provider/v2

go 1.16

replace (
	github.com/hashicorp/go-getter v1.5.0 => github.com/hashicorp/go-getter v1.4.0
)

require (
	github.com/equinix/terraform-provider-metal v1.1.0
	github.com/hashicorp/terraform-plugin-sdk v1.9.1
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.0.0-beta.1
	github.com/pulumi/pulumi/pkg/v3 v3.0.0-beta.2
	github.com/pulumi/pulumi/sdk/v3 v3.0.0-beta.2
)
