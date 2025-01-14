// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mongodbatlas

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"regress-8403/mongodbatlas/internal"
)

func LookupCustomDbRoles(ctx *pulumi.Context, args *LookupCustomDbRolesArgs, opts ...pulumi.InvokeOption) (*LookupCustomDbRolesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupCustomDbRolesResult
	err := ctx.Invoke("mongodbatlas::getCustomDbRoles", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCustomDbRolesArgs struct {
}

type LookupCustomDbRolesResult struct {
	Result *GetCustomDbRolesResult `pulumi:"result"`
}
