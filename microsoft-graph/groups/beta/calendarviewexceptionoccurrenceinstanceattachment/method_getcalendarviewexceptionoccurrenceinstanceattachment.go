package calendarviewexceptionoccurrenceinstanceattachment

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetCalendarViewExceptionOccurrenceInstanceAttachmentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.Attachment
}

type GetCalendarViewExceptionOccurrenceInstanceAttachmentOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetCalendarViewExceptionOccurrenceInstanceAttachmentOperationOptions() GetCalendarViewExceptionOccurrenceInstanceAttachmentOperationOptions {
	return GetCalendarViewExceptionOccurrenceInstanceAttachmentOperationOptions{}
}

func (o GetCalendarViewExceptionOccurrenceInstanceAttachmentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetCalendarViewExceptionOccurrenceInstanceAttachmentOperationOptions) ToOData() *odata.Query {
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

func (o GetCalendarViewExceptionOccurrenceInstanceAttachmentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetCalendarViewExceptionOccurrenceInstanceAttachment - Get attachments from groups. The collection of FileAttachment,
// ItemAttachment, and referenceAttachment attachments for the event. Navigation property. Read-only. Nullable.
func (c CalendarViewExceptionOccurrenceInstanceAttachmentClient) GetCalendarViewExceptionOccurrenceInstanceAttachment(ctx context.Context, id beta.GroupIdCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentId, options GetCalendarViewExceptionOccurrenceInstanceAttachmentOperationOptions) (result GetCalendarViewExceptionOccurrenceInstanceAttachmentOperationResponse, err error) {
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

	var respObj json.RawMessage
	if err = resp.Unmarshal(&respObj); err != nil {
		return
	}
	model, err := beta.UnmarshalAttachmentImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
