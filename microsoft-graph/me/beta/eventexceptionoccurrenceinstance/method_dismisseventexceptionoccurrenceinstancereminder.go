package eventexceptionoccurrenceinstance

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

type DismissEventExceptionOccurrenceInstanceReminderOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DismissEventExceptionOccurrenceInstanceReminderOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultDismissEventExceptionOccurrenceInstanceReminderOperationOptions() DismissEventExceptionOccurrenceInstanceReminderOperationOptions {
	return DismissEventExceptionOccurrenceInstanceReminderOperationOptions{}
}

func (o DismissEventExceptionOccurrenceInstanceReminderOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o DismissEventExceptionOccurrenceInstanceReminderOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DismissEventExceptionOccurrenceInstanceReminderOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DismissEventExceptionOccurrenceInstanceReminder - Invoke action dismissReminder. Dismiss a reminder that has been
// triggered for an event in a user calendar.
func (c EventExceptionOccurrenceInstanceClient) DismissEventExceptionOccurrenceInstanceReminder(ctx context.Context, id beta.MeEventIdExceptionOccurrenceIdInstanceId, options DismissEventExceptionOccurrenceInstanceReminderOperationOptions) (result DismissEventExceptionOccurrenceInstanceReminderOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/dismissReminder", id.ID()),
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
