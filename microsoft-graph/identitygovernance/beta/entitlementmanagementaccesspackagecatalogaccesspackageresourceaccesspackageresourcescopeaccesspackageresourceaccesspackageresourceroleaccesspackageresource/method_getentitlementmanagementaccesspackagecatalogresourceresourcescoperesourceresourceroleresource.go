package entitlementmanagementaccesspackagecatalogaccesspackageresourceaccesspackageresourcescopeaccesspackageresourceaccesspackageresourceroleaccesspackageresource

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

type GetEntitlementManagementAccessPackageCatalogResourceResourceScopeResourceResourceRoleResourceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.AccessPackageResource
}

type GetEntitlementManagementAccessPackageCatalogResourceResourceScopeResourceResourceRoleResourceOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetEntitlementManagementAccessPackageCatalogResourceResourceScopeResourceResourceRoleResourceOperationOptions() GetEntitlementManagementAccessPackageCatalogResourceResourceScopeResourceResourceRoleResourceOperationOptions {
	return GetEntitlementManagementAccessPackageCatalogResourceResourceScopeResourceResourceRoleResourceOperationOptions{}
}

func (o GetEntitlementManagementAccessPackageCatalogResourceResourceScopeResourceResourceRoleResourceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetEntitlementManagementAccessPackageCatalogResourceResourceScopeResourceResourceRoleResourceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetEntitlementManagementAccessPackageCatalogResourceResourceScopeResourceResourceRoleResourceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetEntitlementManagementAccessPackageCatalogResourceResourceScopeResourceResourceRoleResource - Get
// accessPackageResource from identityGovernance
func (c EntitlementManagementAccessPackageCatalogAccessPackageResourceAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleAccessPackageResourceClient) GetEntitlementManagementAccessPackageCatalogResourceResourceScopeResourceResourceRoleResource(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageCatalogIdAccessPackageResourceIdAccessPackageResourceScopeIdAccessPackageResourceAccessPackageResourceRoleId, options GetEntitlementManagementAccessPackageCatalogResourceResourceScopeResourceResourceRoleResourceOperationOptions) (result GetEntitlementManagementAccessPackageCatalogResourceResourceScopeResourceResourceRoleResourceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/accessPackageResource", id.ID()),
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

	var model beta.AccessPackageResource
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
