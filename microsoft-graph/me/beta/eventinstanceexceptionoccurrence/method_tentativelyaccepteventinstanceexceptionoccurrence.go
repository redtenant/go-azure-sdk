package eventinstanceexceptionoccurrence

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

type TentativelyAcceptEventInstanceExceptionOccurrenceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type TentativelyAcceptEventInstanceExceptionOccurrenceOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultTentativelyAcceptEventInstanceExceptionOccurrenceOperationOptions() TentativelyAcceptEventInstanceExceptionOccurrenceOperationOptions {
	return TentativelyAcceptEventInstanceExceptionOccurrenceOperationOptions{}
}

func (o TentativelyAcceptEventInstanceExceptionOccurrenceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o TentativelyAcceptEventInstanceExceptionOccurrenceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o TentativelyAcceptEventInstanceExceptionOccurrenceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// TentativelyAcceptEventInstanceExceptionOccurrence - Invoke action tentativelyAccept. Tentatively accept the specified
// event in a user calendar. If the event allows proposals for new times, on responding tentative to the event, an
// invitee can choose to suggest an alternative time by including the proposedNewTime parameter. For more information on
// how to propose a time, and how to receive and accept a new time proposal, see Propose new meeting times.
func (c EventInstanceExceptionOccurrenceClient) TentativelyAcceptEventInstanceExceptionOccurrence(ctx context.Context, id beta.MeEventIdInstanceIdExceptionOccurrenceId, input TentativelyAcceptEventInstanceExceptionOccurrenceRequest, options TentativelyAcceptEventInstanceExceptionOccurrenceOperationOptions) (result TentativelyAcceptEventInstanceExceptionOccurrenceOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/tentativelyAccept", id.ID()),
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
