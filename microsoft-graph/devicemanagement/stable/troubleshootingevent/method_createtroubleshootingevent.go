package troubleshootingevent

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateTroubleshootingEventOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        stable.DeviceManagementTroubleshootingEvent
}

type CreateTroubleshootingEventOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateTroubleshootingEventOperationOptions() CreateTroubleshootingEventOperationOptions {
	return CreateTroubleshootingEventOperationOptions{}
}

func (o CreateTroubleshootingEventOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateTroubleshootingEventOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateTroubleshootingEventOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateTroubleshootingEvent - Create new navigation property to troubleshootingEvents for deviceManagement
func (c TroubleshootingEventClient) CreateTroubleshootingEvent(ctx context.Context, input stable.DeviceManagementTroubleshootingEvent, options CreateTroubleshootingEventOperationOptions) (result CreateTroubleshootingEventOperationResponse, err error) {
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
		Path:          "/deviceManagement/troubleshootingEvents",
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

	var respObj json.RawMessage
	if err = resp.Unmarshal(&respObj); err != nil {
		return
	}
	model, err := stable.UnmarshalDeviceManagementTroubleshootingEventImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
