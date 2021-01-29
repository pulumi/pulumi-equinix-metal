// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package equinix

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an Equinix Metal organization datasource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-equinix-metal/sdk/go/equinix-metal"
// 	"github.com/pulumi/pulumi-equinix-metal/sdk/go/equinix-metal/"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := local.Org_id
// 		test, err := equinix - metal.LookupOrganization(ctx, &equinix-metal.LookupOrganizationArgs{
// 			OrganizationId: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("projectsInTheOrg", test.ProjectIds)
// 		return nil
// 	})
// }
// ```
func LookupOrganization(ctx *pulumi.Context, args *LookupOrganizationArgs, opts ...pulumi.InvokeOption) (*LookupOrganizationResult, error) {
	var rv LookupOrganizationResult
	err := ctx.Invoke("equinix-metal:index/getOrganization:getOrganization", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOrganization.
type LookupOrganizationArgs struct {
	// The organization name
	Name *string `pulumi:"name"`
	// The UUID of the organization resource
	OrganizationId *string `pulumi:"organizationId"`
}

// A collection of values returned by getOrganization.
type LookupOrganizationResult struct {
	// Description string
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Logo URL
	Logo           string `pulumi:"logo"`
	Name           string `pulumi:"name"`
	OrganizationId string `pulumi:"organizationId"`
	// UUIDs of project resources which belong to this organization
	ProjectIds []string `pulumi:"projectIds"`
	// Twitter handle
	Twitter string `pulumi:"twitter"`
	// Website link
	Website string `pulumi:"website"`
}
