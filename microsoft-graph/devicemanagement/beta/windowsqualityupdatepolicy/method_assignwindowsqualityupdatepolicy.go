package windowsqualityupdatepolicy

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

type AssignWindowsQualityUpdatePolicyOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type AssignWindowsQualityUpdatePolicyOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultAssignWindowsQualityUpdatePolicyOperationOptions() AssignWindowsQualityUpdatePolicyOperationOptions {
	return AssignWindowsQualityUpdatePolicyOperationOptions{}
}

func (o AssignWindowsQualityUpdatePolicyOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AssignWindowsQualityUpdatePolicyOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o AssignWindowsQualityUpdatePolicyOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// AssignWindowsQualityUpdatePolicy - Invoke action assign
func (c WindowsQualityUpdatePolicyClient) AssignWindowsQualityUpdatePolicy(ctx context.Context, id beta.DeviceManagementWindowsQualityUpdatePolicyId, input AssignWindowsQualityUpdatePolicyRequest, options AssignWindowsQualityUpdatePolicyOperationOptions) (result AssignWindowsQualityUpdatePolicyOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/assign", id.ID()),
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
