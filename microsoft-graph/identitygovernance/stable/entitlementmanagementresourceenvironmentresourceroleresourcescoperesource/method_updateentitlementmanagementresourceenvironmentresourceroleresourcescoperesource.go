package entitlementmanagementresourceenvironmentresourceroleresourcescoperesource

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateEntitlementManagementResourceEnvironmentResourceRoleResourceScopeResourceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateEntitlementManagementResourceEnvironmentResourceRoleResourceScopeResourceOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateEntitlementManagementResourceEnvironmentResourceRoleResourceScopeResourceOperationOptions() UpdateEntitlementManagementResourceEnvironmentResourceRoleResourceScopeResourceOperationOptions {
	return UpdateEntitlementManagementResourceEnvironmentResourceRoleResourceScopeResourceOperationOptions{}
}

func (o UpdateEntitlementManagementResourceEnvironmentResourceRoleResourceScopeResourceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateEntitlementManagementResourceEnvironmentResourceRoleResourceScopeResourceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateEntitlementManagementResourceEnvironmentResourceRoleResourceScopeResourceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateEntitlementManagementResourceEnvironmentResourceRoleResourceScopeResource - Update the navigation property
// resource in identityGovernance
func (c EntitlementManagementResourceEnvironmentResourceRoleResourceScopeResourceClient) UpdateEntitlementManagementResourceEnvironmentResourceRoleResourceScopeResource(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdRoleIdResourceScopeId, input stable.AccessPackageResource, options UpdateEntitlementManagementResourceEnvironmentResourceRoleResourceScopeResourceOperationOptions) (result UpdateEntitlementManagementResourceEnvironmentResourceRoleResourceScopeResourceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/resource", id.ID()),
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
