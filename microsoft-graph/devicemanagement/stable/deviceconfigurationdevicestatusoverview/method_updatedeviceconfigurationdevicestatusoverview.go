package deviceconfigurationdevicestatusoverview

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

type UpdateDeviceConfigurationDeviceStatusOverviewOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateDeviceConfigurationDeviceStatusOverviewOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateDeviceConfigurationDeviceStatusOverviewOperationOptions() UpdateDeviceConfigurationDeviceStatusOverviewOperationOptions {
	return UpdateDeviceConfigurationDeviceStatusOverviewOperationOptions{}
}

func (o UpdateDeviceConfigurationDeviceStatusOverviewOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateDeviceConfigurationDeviceStatusOverviewOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateDeviceConfigurationDeviceStatusOverviewOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateDeviceConfigurationDeviceStatusOverview - Update deviceConfigurationDeviceOverview. Update the properties of a
// deviceConfigurationDeviceOverview object.
func (c DeviceConfigurationDeviceStatusOverviewClient) UpdateDeviceConfigurationDeviceStatusOverview(ctx context.Context, id stable.DeviceManagementDeviceConfigurationId, input stable.DeviceConfigurationDeviceOverview, options UpdateDeviceConfigurationDeviceStatusOverviewOperationOptions) (result UpdateDeviceConfigurationDeviceStatusOverviewOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/deviceStatusOverview", id.ID()),
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
