package driveroot

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

type AssignDriveRootSensitivityLabelOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type AssignDriveRootSensitivityLabelOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultAssignDriveRootSensitivityLabelOperationOptions() AssignDriveRootSensitivityLabelOperationOptions {
	return AssignDriveRootSensitivityLabelOperationOptions{}
}

func (o AssignDriveRootSensitivityLabelOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AssignDriveRootSensitivityLabelOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o AssignDriveRootSensitivityLabelOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// AssignDriveRootSensitivityLabel - Invoke action assignSensitivityLabel
func (c DriveRootClient) AssignDriveRootSensitivityLabel(ctx context.Context, id beta.MeDriveId, input AssignDriveRootSensitivityLabelRequest, options AssignDriveRootSensitivityLabelOperationOptions) (result AssignDriveRootSensitivityLabelOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/root/assignSensitivityLabel", id.ID()),
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
