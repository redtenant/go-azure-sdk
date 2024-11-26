package permissionsmanagementscheduledpermissionsapprovalstep

import (
	"context"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdatePermissionsManagementScheduledPermissionsApprovalStepOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdatePermissionsManagementScheduledPermissionsApprovalStepOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdatePermissionsManagementScheduledPermissionsApprovalStepOperationOptions() UpdatePermissionsManagementScheduledPermissionsApprovalStepOperationOptions {
	return UpdatePermissionsManagementScheduledPermissionsApprovalStepOperationOptions{}
}

func (o UpdatePermissionsManagementScheduledPermissionsApprovalStepOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdatePermissionsManagementScheduledPermissionsApprovalStepOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdatePermissionsManagementScheduledPermissionsApprovalStepOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdatePermissionsManagementScheduledPermissionsApprovalStep - Update the navigation property steps in
// identityGovernance
func (c PermissionsManagementScheduledPermissionsApprovalStepClient) UpdatePermissionsManagementScheduledPermissionsApprovalStep(ctx context.Context, id beta.IdentityGovernancePermissionsManagementScheduledPermissionsApprovalIdStepId, input beta.ApprovalStep, options UpdatePermissionsManagementScheduledPermissionsApprovalStepOperationOptions) (result UpdatePermissionsManagementScheduledPermissionsApprovalStepOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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
