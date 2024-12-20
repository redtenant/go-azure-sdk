package teamprimarychannel

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

type CreateTeamPrimaryChannelArchiveOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateTeamPrimaryChannelArchiveOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateTeamPrimaryChannelArchiveOperationOptions() CreateTeamPrimaryChannelArchiveOperationOptions {
	return CreateTeamPrimaryChannelArchiveOperationOptions{}
}

func (o CreateTeamPrimaryChannelArchiveOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateTeamPrimaryChannelArchiveOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateTeamPrimaryChannelArchiveOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateTeamPrimaryChannelArchive - Invoke action archive. Archive a channel in a team. When a channel is archived,
// users can't send new messages or react to existing messages in the channel, edit the channel settings, or make other
// changes to the channel. You can delete an archived channel or add and remove members from it. If you archive a team,
// its channels are also archived. Archiving is an asynchronous operation; a channel is archived after the asynchronous
// archiving operation completes successfully, which might occur after the response returns. A channel without an owner
// or that belongs to a group that has no owner, can't be archived. To restore a channel from its archived state, use
// the channel: unarchive method. A channel can’t be archived or unarchived if its team is archived.
func (c TeamPrimaryChannelClient) CreateTeamPrimaryChannelArchive(ctx context.Context, id beta.GroupId, input CreateTeamPrimaryChannelArchiveRequest, options CreateTeamPrimaryChannelArchiveOperationOptions) (result CreateTeamPrimaryChannelArchiveOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/team/primaryChannel/archive", id.ID()),
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
