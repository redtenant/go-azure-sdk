package entitlementmanagementresourcerequestcatalogresourcescoperesource

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

type DeleteEntitlementManagementResourceRequestCatalogResourceScopeResourceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteEntitlementManagementResourceRequestCatalogResourceScopeResourceOperationOptions struct {
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultDeleteEntitlementManagementResourceRequestCatalogResourceScopeResourceOperationOptions() DeleteEntitlementManagementResourceRequestCatalogResourceScopeResourceOperationOptions {
	return DeleteEntitlementManagementResourceRequestCatalogResourceScopeResourceOperationOptions{}
}

func (o DeleteEntitlementManagementResourceRequestCatalogResourceScopeResourceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteEntitlementManagementResourceRequestCatalogResourceScopeResourceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeleteEntitlementManagementResourceRequestCatalogResourceScopeResourceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteEntitlementManagementResourceRequestCatalogResourceScopeResource - Delete navigation property resource for
// identityGovernance
func (c EntitlementManagementResourceRequestCatalogResourceScopeResourceClient) DeleteEntitlementManagementResourceRequestCatalogResourceScopeResource(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdScopeId, options DeleteEntitlementManagementResourceRequestCatalogResourceScopeResourceOperationOptions) (result DeleteEntitlementManagementResourceRequestCatalogResourceScopeResourceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/resource", id.ID()),
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
