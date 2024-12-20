package message

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateMessageReplyAllOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        stable.Message
}

type CreateMessageReplyAllOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateMessageReplyAllOperationOptions() CreateMessageReplyAllOperationOptions {
	return CreateMessageReplyAllOperationOptions{}
}

func (o CreateMessageReplyAllOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateMessageReplyAllOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateMessageReplyAllOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateMessageReplyAll - Invoke action createReplyAll. Create a draft to reply to the sender and all recipients of a
// message in either JSON or MIME format. When using JSON format: - Specify either a comment or the body property of the
// message parameter. Specifying both will return an HTTP 400 Bad Request error. - If the original message specifies a
// recipient in the replyTo property, per Internet Message Format (RFC 2822), you should send the reply to the
// recipients in the replyTo and toRecipients properties, and not the recipients in the from and toRecipients
// properties. - You can update the draft later to add reply content to the body or change other message properties.
// When using MIME format: - Provide the applicable Internet message headers and the MIME content, all encoded in base64
// format in the request body. - Add any attachments and S/MIME properties to the MIME content. Send the draft message
// in a subsequent operation. Alternatively, reply-all to a message in a single action.
func (c MessageClient) CreateMessageReplyAll(ctx context.Context, id stable.MeMessageId, input CreateMessageReplyAllRequest, options CreateMessageReplyAllOperationOptions) (result CreateMessageReplyAllOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/createReplyAll", id.ID()),
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
	model, err := stable.UnmarshalMessageImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
