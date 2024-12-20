package entitlementmanagementroleassignmentschedulerequest

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

type CancelEntitlementManagementRoleAssignmentScheduleRequestOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CancelEntitlementManagementRoleAssignmentScheduleRequestOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCancelEntitlementManagementRoleAssignmentScheduleRequestOperationOptions() CancelEntitlementManagementRoleAssignmentScheduleRequestOperationOptions {
	return CancelEntitlementManagementRoleAssignmentScheduleRequestOperationOptions{}
}

func (o CancelEntitlementManagementRoleAssignmentScheduleRequestOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CancelEntitlementManagementRoleAssignmentScheduleRequestOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CancelEntitlementManagementRoleAssignmentScheduleRequestOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CancelEntitlementManagementRoleAssignmentScheduleRequest - Invoke action cancel. Immediately cancel a
// unifiedRoleAssignmentScheduleRequest object that is in a Granted status, and have the system automatically delete the
// canceled request after 30 days. After calling this action, the status of the canceled
// unifiedRoleAssignmentScheduleRequest changes to Canceled.
func (c EntitlementManagementRoleAssignmentScheduleRequestClient) CancelEntitlementManagementRoleAssignmentScheduleRequest(ctx context.Context, id stable.RoleManagementEntitlementManagementRoleAssignmentScheduleRequestId, options CancelEntitlementManagementRoleAssignmentScheduleRequestOperationOptions) (result CancelEntitlementManagementRoleAssignmentScheduleRequestOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/cancel", id.ID()),
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
