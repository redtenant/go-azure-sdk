package eventinstanceexceptionoccurrenceextension

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

type CreateEventInstanceExceptionOccurrenceExtensionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.Extension
}

type CreateEventInstanceExceptionOccurrenceExtensionOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateEventInstanceExceptionOccurrenceExtensionOperationOptions() CreateEventInstanceExceptionOccurrenceExtensionOperationOptions {
	return CreateEventInstanceExceptionOccurrenceExtensionOperationOptions{}
}

func (o CreateEventInstanceExceptionOccurrenceExtensionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateEventInstanceExceptionOccurrenceExtensionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateEventInstanceExceptionOccurrenceExtensionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateEventInstanceExceptionOccurrenceExtension - Create new navigation property to extensions for groups
func (c EventInstanceExceptionOccurrenceExtensionClient) CreateEventInstanceExceptionOccurrenceExtension(ctx context.Context, id beta.GroupIdEventIdInstanceIdExceptionOccurrenceId, input beta.Extension, options CreateEventInstanceExceptionOccurrenceExtensionOperationOptions) (result CreateEventInstanceExceptionOccurrenceExtensionOperationResponse, err error) {
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
