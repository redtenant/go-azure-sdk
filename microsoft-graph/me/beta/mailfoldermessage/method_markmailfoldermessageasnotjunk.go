package mailfoldermessage

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

type MarkMailFolderMessageAsNotJunkOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.Message
}

type MarkMailFolderMessageAsNotJunkOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultMarkMailFolderMessageAsNotJunkOperationOptions() MarkMailFolderMessageAsNotJunkOperationOptions {
	return MarkMailFolderMessageAsNotJunkOperationOptions{}
}

func (o MarkMailFolderMessageAsNotJunkOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o MarkMailFolderMessageAsNotJunkOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o MarkMailFolderMessageAsNotJunkOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// MarkMailFolderMessageAsNotJunk - Invoke action markAsNotJunk. Mark a message as not junk. This API removes the sender
// from the list of blocked senders and moves the message to the Inbox folder, when moveToInbox is true.
func (c MailFolderMessageClient) MarkMailFolderMessageAsNotJunk(ctx context.Context, id beta.MeMailFolderIdMessageId, input MarkMailFolderMessageAsNotJunkRequest, options MarkMailFolderMessageAsNotJunkOperationOptions) (result MarkMailFolderMessageAsNotJunkOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/markAsNotJunk", id.ID()),
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
	model, err := beta.UnmarshalMessageImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
