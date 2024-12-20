package entitlementmanagementassignmentrequest

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

type CancelEntitlementManagementAssignmentRequestOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CancelEntitlementManagementAssignmentRequestOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCancelEntitlementManagementAssignmentRequestOperationOptions() CancelEntitlementManagementAssignmentRequestOperationOptions {
	return CancelEntitlementManagementAssignmentRequestOperationOptions{}
}

func (o CancelEntitlementManagementAssignmentRequestOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CancelEntitlementManagementAssignmentRequestOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CancelEntitlementManagementAssignmentRequestOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CancelEntitlementManagementAssignmentRequest - Invoke action cancel. In Microsoft Entra Entitlement Management,
// cancel accessPackageAssignmentRequest objects that are in a cancellable state: accepted, pendingApproval,
// pendingNotBefore, pendingApprovalEscalated.
func (c EntitlementManagementAssignmentRequestClient) CancelEntitlementManagementAssignmentRequest(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementAssignmentRequestId, options CancelEntitlementManagementAssignmentRequestOperationOptions) (result CancelEntitlementManagementAssignmentRequestOperationResponse, err error) {
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
