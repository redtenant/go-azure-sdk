package windowsautopilotsetting

import (
	"context"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateWindowsAutopilotSettingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateWindowsAutopilotSettingOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateWindowsAutopilotSettingOperationOptions() UpdateWindowsAutopilotSettingOperationOptions {
	return UpdateWindowsAutopilotSettingOperationOptions{}
}

func (o UpdateWindowsAutopilotSettingOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateWindowsAutopilotSettingOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateWindowsAutopilotSettingOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateWindowsAutopilotSetting - Update the navigation property windowsAutopilotSettings in deviceManagement
func (c WindowsAutopilotSettingClient) UpdateWindowsAutopilotSetting(ctx context.Context, input beta.WindowsAutopilotSettings, options UpdateWindowsAutopilotSettingOperationOptions) (result UpdateWindowsAutopilotSettingOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          "/deviceManagement/windowsAutopilotSettings",
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
