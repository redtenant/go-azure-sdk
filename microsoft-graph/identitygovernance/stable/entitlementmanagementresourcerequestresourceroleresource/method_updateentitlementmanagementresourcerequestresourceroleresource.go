package entitlementmanagementresourcerequestresourceroleresource

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

type UpdateEntitlementManagementResourceRequestResourceRoleResourceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateEntitlementManagementResourceRequestResourceRoleResourceOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateEntitlementManagementResourceRequestResourceRoleResourceOperationOptions() UpdateEntitlementManagementResourceRequestResourceRoleResourceOperationOptions {
	return UpdateEntitlementManagementResourceRequestResourceRoleResourceOperationOptions{}
}

func (o UpdateEntitlementManagementResourceRequestResourceRoleResourceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateEntitlementManagementResourceRequestResourceRoleResourceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateEntitlementManagementResourceRequestResourceRoleResourceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateEntitlementManagementResourceRequestResourceRoleResource - Update the navigation property resource in
// identityGovernance
func (c EntitlementManagementResourceRequestResourceRoleResourceClient) UpdateEntitlementManagementResourceRequestResourceRoleResource(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementResourceRequestIdResourceRoleId, input stable.AccessPackageResource, options UpdateEntitlementManagementResourceRequestResourceRoleResourceOperationOptions) (result UpdateEntitlementManagementResourceRequestResourceRoleResourceOperationResponse, err error) {
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
