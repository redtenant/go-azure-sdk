package conversationthreadpostinreplytoattachment

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

type GetConversationThreadPostInReplyToAttachmentsCountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetConversationThreadPostInReplyToAttachmentsCountOperationOptions struct {
	Filter    *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Search    *string
}

func DefaultGetConversationThreadPostInReplyToAttachmentsCountOperationOptions() GetConversationThreadPostInReplyToAttachmentsCountOperationOptions {
	return GetConversationThreadPostInReplyToAttachmentsCountOperationOptions{}
}

func (o GetConversationThreadPostInReplyToAttachmentsCountOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetConversationThreadPostInReplyToAttachmentsCountOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	return &out
}

func (o GetConversationThreadPostInReplyToAttachmentsCountOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetConversationThreadPostInReplyToAttachmentsCount - Get the number of the resource
func (c ConversationThreadPostInReplyToAttachmentClient) GetConversationThreadPostInReplyToAttachmentsCount(ctx context.Context, id beta.GroupIdConversationIdThreadIdPostId, options GetConversationThreadPostInReplyToAttachmentsCountOperationOptions) (result GetConversationThreadPostInReplyToAttachmentsCountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "text/plain",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/inReplyTo/attachments/$count", id.ID()),
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

	var model []byte
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
