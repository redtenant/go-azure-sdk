package mobileapptroubleshootingevent

import (
	"context"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateMobileAppTroubleshootingEventOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.MobileAppTroubleshootingEvent
}

type CreateMobileAppTroubleshootingEventOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateMobileAppTroubleshootingEventOperationOptions() CreateMobileAppTroubleshootingEventOperationOptions {
	return CreateMobileAppTroubleshootingEventOperationOptions{}
}

func (o CreateMobileAppTroubleshootingEventOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateMobileAppTroubleshootingEventOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateMobileAppTroubleshootingEventOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateMobileAppTroubleshootingEvent - Create mobileAppTroubleshootingEvent. Create a new
// mobileAppTroubleshootingEvent object.
func (c MobileAppTroubleshootingEventClient) CreateMobileAppTroubleshootingEvent(ctx context.Context, input stable.MobileAppTroubleshootingEvent, options CreateMobileAppTroubleshootingEventOperationOptions) (result CreateMobileAppTroubleshootingEventOperationResponse, err error) {
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
		Path:          "/deviceManagement/mobileAppTroubleshootingEvents",
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

	var model stable.MobileAppTroubleshootingEvent
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
