package calendareventexceptionoccurrenceinstanceextension

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateCalendarEventExceptionOccurrenceInstanceExtensionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.Extension
}

type CreateCalendarEventExceptionOccurrenceInstanceExtensionOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateCalendarEventExceptionOccurrenceInstanceExtensionOperationOptions() CreateCalendarEventExceptionOccurrenceInstanceExtensionOperationOptions {
	return CreateCalendarEventExceptionOccurrenceInstanceExtensionOperationOptions{}
}

func (o CreateCalendarEventExceptionOccurrenceInstanceExtensionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateCalendarEventExceptionOccurrenceInstanceExtensionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateCalendarEventExceptionOccurrenceInstanceExtensionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateCalendarEventExceptionOccurrenceInstanceExtension - Create new navigation property to extensions for groups
func (c CalendarEventExceptionOccurrenceInstanceExtensionClient) CreateCalendarEventExceptionOccurrenceInstanceExtension(ctx context.Context, id beta.GroupIdCalendarEventIdExceptionOccurrenceIdInstanceId, input beta.Extension, options CreateCalendarEventExceptionOccurrenceInstanceExtensionOperationOptions) (result CreateCalendarEventExceptionOccurrenceInstanceExtensionOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/extensions", id.ID()),
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
	model, err := beta.UnmarshalExtensionImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
