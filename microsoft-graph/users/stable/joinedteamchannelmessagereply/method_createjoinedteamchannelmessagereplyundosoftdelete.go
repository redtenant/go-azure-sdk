package joinedteamchannelmessagereply

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

type CreateJoinedTeamChannelMessageReplyUndoSoftDeleteOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateJoinedTeamChannelMessageReplyUndoSoftDeleteOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateJoinedTeamChannelMessageReplyUndoSoftDeleteOperationOptions() CreateJoinedTeamChannelMessageReplyUndoSoftDeleteOperationOptions {
	return CreateJoinedTeamChannelMessageReplyUndoSoftDeleteOperationOptions{}
}

func (o CreateJoinedTeamChannelMessageReplyUndoSoftDeleteOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateJoinedTeamChannelMessageReplyUndoSoftDeleteOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateJoinedTeamChannelMessageReplyUndoSoftDeleteOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateJoinedTeamChannelMessageReplyUndoSoftDelete - Invoke action undoSoftDelete. Undo soft deletion of a single
// chatMessage or a chat message reply in a channel or a chat.
func (c JoinedTeamChannelMessageReplyClient) CreateJoinedTeamChannelMessageReplyUndoSoftDelete(ctx context.Context, id stable.UserIdJoinedTeamIdChannelIdMessageIdReplyId, options CreateJoinedTeamChannelMessageReplyUndoSoftDeleteOperationOptions) (result CreateJoinedTeamChannelMessageReplyUndoSoftDeleteOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/undoSoftDelete", id.ID()),
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

	return
}
