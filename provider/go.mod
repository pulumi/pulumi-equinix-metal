module github.com/pulumi/pulumi-equinix-metal/provider

go 1.15

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible

require (
	github.com/equinix/terraform-provider-metal v1.0.1-0.20201209021325-714995571fa3
	github.com/hashicorp/terraform-plugin-sdk v1.9.1
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.16.0
	github.com/pulumi/pulumi/pkg/v2 v2.15.1-0.20201202214525-260620430c4c
	github.com/pulumi/pulumi/sdk/v2 v2.16.0
)
