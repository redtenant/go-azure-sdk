package calendarcalendarview

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

type DeclineCalendarViewOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeclineCalendarViewOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultDeclineCalendarViewOperationOptions() DeclineCalendarViewOperationOptions {
	return DeclineCalendarViewOperationOptions{}
}

func (o DeclineCalendarViewOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o DeclineCalendarViewOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeclineCalendarViewOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeclineCalendarView - Invoke action decline. Decline invitation to the specified event in a user calendar. If the
// event allows proposals for new times, on declining the event, an invitee can choose to suggest an alternative time by
// including the proposedNewTime parameter. For more information on how to propose a time, and how to receive and accept
// a new time proposal, see Propose new meeting times.
func (c CalendarCalendarViewClient) DeclineCalendarView(ctx context.Context, id stable.GroupIdCalendarCalendarViewId, input DeclineCalendarViewRequest, options DeclineCalendarViewOperationOptions) (result DeclineCalendarViewOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/decline", id.ID()),
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
