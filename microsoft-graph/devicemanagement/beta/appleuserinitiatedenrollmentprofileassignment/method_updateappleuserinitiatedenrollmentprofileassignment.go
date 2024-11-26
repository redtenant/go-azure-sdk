package appleuserinitiatedenrollmentprofileassignment

import (
	"context"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateAppleUserInitiatedEnrollmentProfileAssignmentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateAppleUserInitiatedEnrollmentProfileAssignmentOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateAppleUserInitiatedEnrollmentProfileAssignmentOperationOptions() UpdateAppleUserInitiatedEnrollmentProfileAssignmentOperationOptions {
	return UpdateAppleUserInitiatedEnrollmentProfileAssignmentOperationOptions{}
}

func (o UpdateAppleUserInitiatedEnrollmentProfileAssignmentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateAppleUserInitiatedEnrollmentProfileAssignmentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateAppleUserInitiatedEnrollmentProfileAssignmentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateAppleUserInitiatedEnrollmentProfileAssignment - Update the navigation property assignments in deviceManagement
func (c AppleUserInitiatedEnrollmentProfileAssignmentClient) UpdateAppleUserInitiatedEnrollmentProfileAssignment(ctx context.Context, id beta.DeviceManagementAppleUserInitiatedEnrollmentProfileIdAssignmentId, input beta.AppleEnrollmentProfileAssignment, options UpdateAppleUserInitiatedEnrollmentProfileAssignmentOperationOptions) (result UpdateAppleUserInitiatedEnrollmentProfileAssignmentOperationResponse, err error) {
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
