package teamchannelmessagereply

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

type UnsetTeamChannelMessageReplyReactionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UnsetTeamChannelMessageReplyReactionOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUnsetTeamChannelMessageReplyReactionOperationOptions() UnsetTeamChannelMessageReplyReactionOperationOptions {
	return UnsetTeamChannelMessageReplyReactionOperationOptions{}
}

func (o UnsetTeamChannelMessageReplyReactionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UnsetTeamChannelMessageReplyReactionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UnsetTeamChannelMessageReplyReactionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UnsetTeamChannelMessageReplyReaction - Invoke action unsetReaction
func (c TeamChannelMessageReplyClient) UnsetTeamChannelMessageReplyReaction(ctx context.Context, id stable.GroupIdTeamChannelIdMessageIdReplyId, input UnsetTeamChannelMessageReplyReactionRequest, options UnsetTeamChannelMessageReplyReactionOperationOptions) (result UnsetTeamChannelMessageReplyReactionOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/unsetReaction", id.ID()),
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
