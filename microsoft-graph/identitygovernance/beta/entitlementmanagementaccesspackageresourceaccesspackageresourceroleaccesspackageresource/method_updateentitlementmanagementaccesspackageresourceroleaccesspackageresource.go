package entitlementmanagementaccesspackageresourceaccesspackageresourceroleaccesspackageresource

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

type UpdateEntitlementManagementAccessPackageResourceRoleAccessPackageResourceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateEntitlementManagementAccessPackageResourceRoleAccessPackageResourceOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateEntitlementManagementAccessPackageResourceRoleAccessPackageResourceOperationOptions() UpdateEntitlementManagementAccessPackageResourceRoleAccessPackageResourceOperationOptions {
	return UpdateEntitlementManagementAccessPackageResourceRoleAccessPackageResourceOperationOptions{}
}

func (o UpdateEntitlementManagementAccessPackageResourceRoleAccessPackageResourceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateEntitlementManagementAccessPackageResourceRoleAccessPackageResourceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateEntitlementManagementAccessPackageResourceRoleAccessPackageResourceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateEntitlementManagementAccessPackageResourceRoleAccessPackageResource - Update the navigation property
// accessPackageResource in identityGovernance
func (c EntitlementManagementAccessPackageResourceAccessPackageResourceRoleAccessPackageResourceClient) UpdateEntitlementManagementAccessPackageResourceRoleAccessPackageResource(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageResourceIdAccessPackageResourceRoleId, input beta.AccessPackageResource, options UpdateEntitlementManagementAccessPackageResourceRoleAccessPackageResourceOperationOptions) (result UpdateEntitlementManagementAccessPackageResourceRoleAccessPackageResourceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/accessPackageResource", id.ID()),
		RetryFunc:     options.RetryFunc,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
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
