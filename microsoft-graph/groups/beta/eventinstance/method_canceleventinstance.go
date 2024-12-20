package eventinstance

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

type CancelEventInstanceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CancelEventInstanceOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCancelEventInstanceOperationOptions() CancelEventInstanceOperationOptions {
	return CancelEventInstanceOperationOptions{}
}

func (o CancelEventInstanceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CancelEventInstanceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CancelEventInstanceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CancelEventInstance - Invoke action cancel. This action allows the organizer of a meeting to send a cancellation
// message and cancel the event. The action moves the event to the Deleted Items folder. The organizer can also cancel
// an occurrence of a recurring meeting by providing the occurrence event ID. An attendee calling this action gets an
// error (HTTP 400 Bad Request), with the following error message: 'Your request can't be completed. You need to be an
// organizer to cancel a meeting.' This action differs from Delete in that Cancel is available to only the organizer,
// and lets the organizer send a custom message to the attendees about the cancellation.
func (c EventInstanceClient) CancelEventInstance(ctx context.Context, id beta.GroupIdEventIdInstanceId, input CancelEventInstanceRequest, options CancelEventInstanceOperationOptions) (result CancelEventInstanceOperationResponse, err error) {
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
