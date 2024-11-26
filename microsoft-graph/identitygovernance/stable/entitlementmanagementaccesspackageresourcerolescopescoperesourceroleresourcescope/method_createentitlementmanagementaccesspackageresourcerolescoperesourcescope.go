package entitlementmanagementaccesspackageresourcerolescopescoperesourceroleresourcescope

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

type CreateEntitlementManagementAccessPackageResourceRoleScopeResourceScopeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.AccessPackageResourceScope
}

type CreateEntitlementManagementAccessPackageResourceRoleScopeResourceScopeOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateEntitlementManagementAccessPackageResourceRoleScopeResourceScopeOperationOptions() CreateEntitlementManagementAccessPackageResourceRoleScopeResourceScopeOperationOptions {
	return CreateEntitlementManagementAccessPackageResourceRoleScopeResourceScopeOperationOptions{}
}

func (o CreateEntitlementManagementAccessPackageResourceRoleScopeResourceScopeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateEntitlementManagementAccessPackageResourceRoleScopeResourceScopeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateEntitlementManagementAccessPackageResourceRoleScopeResourceScopeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateEntitlementManagementAccessPackageResourceRoleScopeResourceScope - Create new navigation property to scopes for
// identityGovernance
func (c EntitlementManagementAccessPackageResourceRoleScopeScopeResourceRoleResourceScopeClient) CreateEntitlementManagementAccessPackageResourceRoleScopeResourceScope(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdScopeResourceRoleId, input stable.AccessPackageResourceScope, options CreateEntitlementManagementAccessPackageResourceRoleScopeResourceScopeOperationOptions) (result CreateEntitlementManagementAccessPackageResourceRoleScopeResourceScopeOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/resource/scopes", id.ID()),
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

	var model stable.AccessPackageResourceScope
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
