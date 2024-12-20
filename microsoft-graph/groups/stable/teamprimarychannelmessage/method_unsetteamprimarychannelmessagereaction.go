package teamprimarychannelmessage

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

type UnsetTeamPrimaryChannelMessageReactionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UnsetTeamPrimaryChannelMessageReactionOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUnsetTeamPrimaryChannelMessageReactionOperationOptions() UnsetTeamPrimaryChannelMessageReactionOperationOptions {
	return UnsetTeamPrimaryChannelMessageReactionOperationOptions{}
}

func (o UnsetTeamPrimaryChannelMessageReactionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UnsetTeamPrimaryChannelMessageReactionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UnsetTeamPrimaryChannelMessageReactionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UnsetTeamPrimaryChannelMessageReaction - Invoke action unsetReaction
func (c TeamPrimaryChannelMessageClient) UnsetTeamPrimaryChannelMessageReaction(ctx context.Context, id stable.GroupIdTeamPrimaryChannelMessageId, input UnsetTeamPrimaryChannelMessageReactionRequest, options UnsetTeamPrimaryChannelMessageReactionOperationOptions) (result UnsetTeamPrimaryChannelMessageReactionOperationResponse, err error) {
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
