package entitlementmanagementconnectedorganization

import (
	"context"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateEntitlementManagementConnectedOrganizationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.ConnectedOrganization
}

type CreateEntitlementManagementConnectedOrganizationOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateEntitlementManagementConnectedOrganizationOperationOptions() CreateEntitlementManagementConnectedOrganizationOperationOptions {
	return CreateEntitlementManagementConnectedOrganizationOperationOptions{}
}

func (o CreateEntitlementManagementConnectedOrganizationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateEntitlementManagementConnectedOrganizationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateEntitlementManagementConnectedOrganizationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateEntitlementManagementConnectedOrganization - Create connectedOrganization. Create a new connectedOrganization
// object.
func (c EntitlementManagementConnectedOrganizationClient) CreateEntitlementManagementConnectedOrganization(ctx context.Context, input stable.ConnectedOrganization, options CreateEntitlementManagementConnectedOrganizationOperationOptions) (result CreateEntitlementManagementConnectedOrganizationOperationResponse, err error) {
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
		Path:          "/identityGovernance/entitlementManagement/connectedOrganizations",
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

	var model stable.ConnectedOrganization
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
