package entitlementmanagementtransitiveroleassignment

import (
	"context"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateEntitlementManagementTransitiveRoleAssignmentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.UnifiedRoleAssignment
}

type CreateEntitlementManagementTransitiveRoleAssignmentOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateEntitlementManagementTransitiveRoleAssignmentOperationOptions() CreateEntitlementManagementTransitiveRoleAssignmentOperationOptions {
	return CreateEntitlementManagementTransitiveRoleAssignmentOperationOptions{}
}

func (o CreateEntitlementManagementTransitiveRoleAssignmentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateEntitlementManagementTransitiveRoleAssignmentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateEntitlementManagementTransitiveRoleAssignmentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateEntitlementManagementTransitiveRoleAssignment - Create new navigation property to transitiveRoleAssignments for
// roleManagement
func (c EntitlementManagementTransitiveRoleAssignmentClient) CreateEntitlementManagementTransitiveRoleAssignment(ctx context.Context, input beta.UnifiedRoleAssignment, options CreateEntitlementManagementTransitiveRoleAssignmentOperationOptions) (result CreateEntitlementManagementTransitiveRoleAssignmentOperationResponse, err error) {
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
		Path:          "/roleManagement/entitlementManagement/transitiveRoleAssignments",
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

	var model beta.UnifiedRoleAssignment
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
