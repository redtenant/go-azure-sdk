package entitlementmanagementaccesspackageresourceenvironment

import (
	"context"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateEntitlementManagementAccessPackageResourceEnvironmentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.AccessPackageResourceEnvironment
}

type CreateEntitlementManagementAccessPackageResourceEnvironmentOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateEntitlementManagementAccessPackageResourceEnvironmentOperationOptions() CreateEntitlementManagementAccessPackageResourceEnvironmentOperationOptions {
	return CreateEntitlementManagementAccessPackageResourceEnvironmentOperationOptions{}
}

func (o CreateEntitlementManagementAccessPackageResourceEnvironmentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateEntitlementManagementAccessPackageResourceEnvironmentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateEntitlementManagementAccessPackageResourceEnvironmentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateEntitlementManagementAccessPackageResourceEnvironment - Create new navigation property to
// accessPackageResourceEnvironments for identityGovernance
func (c EntitlementManagementAccessPackageResourceEnvironmentClient) CreateEntitlementManagementAccessPackageResourceEnvironment(ctx context.Context, input beta.AccessPackageResourceEnvironment, options CreateEntitlementManagementAccessPackageResourceEnvironmentOperationOptions) (result CreateEntitlementManagementAccessPackageResourceEnvironmentOperationResponse, err error) {
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
		Path:          "/identityGovernance/entitlementManagement/accessPackageResourceEnvironments",
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

	var model beta.AccessPackageResourceEnvironment
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
