package calendarcalendarviewinstance

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

type GetCalendarViewInstanceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.Event
}

type GetCalendarViewInstanceOperationOptions struct {
	EndDateTime   *string
	Expand        *odata.Expand
	Metadata      *odata.Metadata
	RetryFunc     client.RequestRetryFunc
	Select        *[]string
	StartDateTime *string
}

func DefaultGetCalendarViewInstanceOperationOptions() GetCalendarViewInstanceOperationOptions {
	return GetCalendarViewInstanceOperationOptions{}
}

func (o GetCalendarViewInstanceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetCalendarViewInstanceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetCalendarViewInstanceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.EndDateTime != nil {
		out.Append("endDateTime", fmt.Sprintf("%v", *o.EndDateTime))
	}
	if o.StartDateTime != nil {
		out.Append("startDateTime", fmt.Sprintf("%v", *o.StartDateTime))
	}
	return &out
}

// GetCalendarViewInstance - Get instances from groups. The occurrences of a recurring series, if the event is a series
// master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been
// modified, but doesn't include occurrences that have been canceled from the series. Navigation property. Read-only.
// Nullable.
func (c CalendarCalendarViewInstanceClient) GetCalendarViewInstance(ctx context.Context, id beta.GroupIdCalendarCalendarViewIdInstanceId, options GetCalendarViewInstanceOperationOptions) (result GetCalendarViewInstanceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
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

	var model beta.Event
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
