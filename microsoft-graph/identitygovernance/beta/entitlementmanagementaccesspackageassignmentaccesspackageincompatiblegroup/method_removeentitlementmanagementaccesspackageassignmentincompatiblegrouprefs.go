package entitlementmanagementaccesspackageassignmentaccesspackageincompatiblegroup

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemoveEntitlementManagementAccessPackageAssignmentIncompatibleGroupRefsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type RemoveEntitlementManagementAccessPackageAssignmentIncompatibleGroupRefsOperationOptions struct {
	Id        *string
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultRemoveEntitlementManagementAccessPackageAssignmentIncompatibleGroupRefsOperationOptions() RemoveEntitlementManagementAccessPackageAssignmentIncompatibleGroupRefsOperationOptions {
	return RemoveEntitlementManagementAccessPackageAssignmentIncompatibleGroupRefsOperationOptions{}
}

func (o RemoveEntitlementManagementAccessPackageAssignmentIncompatibleGroupRefsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o RemoveEntitlementManagementAccessPackageAssignmentIncompatibleGroupRefsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o RemoveEntitlementManagementAccessPackageAssignmentIncompatibleGroupRefsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Id != nil {
		out.Append("@id", fmt.Sprintf("%v", *o.Id))
	}
	return &out
}

// RemoveEntitlementManagementAccessPackageAssignmentIncompatibleGroupRefs - Delete ref of navigation property
// incompatibleGroups for identityGovernance
func (c EntitlementManagementAccessPackageAssignmentAccessPackageIncompatibleGroupClient) RemoveEntitlementManagementAccessPackageAssignmentIncompatibleGroupRefs(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageAssignmentId, options RemoveEntitlementManagementAccessPackageAssignmentIncompatibleGroupRefsOperationOptions) (result RemoveEntitlementManagementAccessPackageAssignmentIncompatibleGroupRefsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/accessPackage/incompatibleGroups/$ref", id.ID()),
		RetryFunc:     options.RetryFunc,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	return
}
