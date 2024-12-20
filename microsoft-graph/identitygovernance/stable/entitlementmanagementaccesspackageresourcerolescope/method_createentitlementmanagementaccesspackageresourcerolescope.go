package entitlementmanagementaccesspackageresourcerolescope

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

type CreateEntitlementManagementAccessPackageResourceRoleScopeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.AccessPackageResourceRoleScope
}

type CreateEntitlementManagementAccessPackageResourceRoleScopeOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateEntitlementManagementAccessPackageResourceRoleScopeOperationOptions() CreateEntitlementManagementAccessPackageResourceRoleScopeOperationOptions {
	return CreateEntitlementManagementAccessPackageResourceRoleScopeOperationOptions{}
}

func (o CreateEntitlementManagementAccessPackageResourceRoleScopeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateEntitlementManagementAccessPackageResourceRoleScopeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateEntitlementManagementAccessPackageResourceRoleScopeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateEntitlementManagementAccessPackageResourceRoleScope - Create resourceRoleScope. Create a new
// accessPackageResourceRoleScope for adding a resource role to an access package. The access package resource, for a
// group, an app, or a SharePoint Online site, must already exist in the access package catalog, and the originId for
// the resource role retrieved from the list of the resource roles. Once you add the resource role scope to the access
// package, the user will receive this resource role through any current and future access package assignments.
func (c EntitlementManagementAccessPackageResourceRoleScopeClient) CreateEntitlementManagementAccessPackageResourceRoleScope(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementAccessPackageId, input stable.AccessPackageResourceRoleScope, options CreateEntitlementManagementAccessPackageResourceRoleScopeOperationOptions) (result CreateEntitlementManagementAccessPackageResourceRoleScopeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/resourceRoleScopes", id.ID()),
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

	var model stable.AccessPackageResourceRoleScope
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
