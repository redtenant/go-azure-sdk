package calendareventexceptionoccurrence

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

type SnoozeCalendarEventExceptionOccurrenceReminderOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SnoozeCalendarEventExceptionOccurrenceReminderOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultSnoozeCalendarEventExceptionOccurrenceReminderOperationOptions() SnoozeCalendarEventExceptionOccurrenceReminderOperationOptions {
	return SnoozeCalendarEventExceptionOccurrenceReminderOperationOptions{}
}

func (o SnoozeCalendarEventExceptionOccurrenceReminderOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SnoozeCalendarEventExceptionOccurrenceReminderOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SnoozeCalendarEventExceptionOccurrenceReminderOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SnoozeCalendarEventExceptionOccurrenceReminder - Invoke action snoozeReminder. Postpone a reminder for an event in a
// user calendar until a new time.
func (c CalendarEventExceptionOccurrenceClient) SnoozeCalendarEventExceptionOccurrenceReminder(ctx context.Context, id beta.GroupIdCalendarEventIdExceptionOccurrenceId, input SnoozeCalendarEventExceptionOccurrenceReminderRequest, options SnoozeCalendarEventExceptionOccurrenceReminderOperationOptions) (result SnoozeCalendarEventExceptionOccurrenceReminderOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/snoozeReminder", id.ID()),
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
