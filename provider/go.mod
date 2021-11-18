module github.com/pulumi/pulumi-equinix-metal/provider/v3

go 1.16

replace (
	github.com/hashicorp/go-getter v1.5.0 => github.com/hashicorp/go-getter v1.4.0
	github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20210629210550-59d24255d71f
)

require (
	github.com/equinix/terraform-provider-metal v1.1.1-0.20210929123308-2d7fb2c4b0ce
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.11.0
	github.com/pulumi/pulumi/pkg/v3 v3.17.0
	github.com/pulumi/pulumi/sdk/v3 v3.17.0
)
