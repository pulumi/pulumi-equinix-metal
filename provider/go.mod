module github.com/pulumi/pulumi-equinix-metal/provider/v3

go 1.16

replace (
	github.com/hashicorp/go-getter v1.5.0 => github.com/hashicorp/go-getter v1.4.0
	github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20220927231159-9ccf93a70751
)

require (
	github.com/agext/levenshtein v1.2.3 // indirect
	github.com/equinix/terraform-provider-metal v1.1.1-0.20211208211059-f3d7fe9ae128
	github.com/hashicorp/go-hclog v1.3.1 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.1 // indirect
	github.com/hashicorp/hcl/v2 v2.15.0 // indirect
	github.com/hashicorp/terraform-plugin-go v0.14.2 // indirect
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.24.1 // indirect
	github.com/hashicorp/yamux v0.1.1 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/oklog/run v1.1.0 // indirect
	github.com/packethost/packngo v0.29.0 // indirect
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.21.0
	github.com/pulumi/pulumi/pkg/v3 v3.30.0
	github.com/pulumi/pulumi/sdk/v3 v3.30.0
	github.com/vmihailenco/tagparser v0.1.2 // indirect
	golang.org/x/crypto v0.3.0 // indirect
	google.golang.org/genproto v0.0.0-20221118155620-16455021b5e6 // indirect
)
