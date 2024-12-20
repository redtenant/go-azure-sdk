package teamworkassociatedteam

import (
	"context"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateTeamworkAssociatedTeamOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateTeamworkAssociatedTeamOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateTeamworkAssociatedTeamOperationOptions() UpdateTeamworkAssociatedTeamOperationOptions {
	return UpdateTeamworkAssociatedTeamOperationOptions{}
}

func (o UpdateTeamworkAssociatedTeamOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateTeamworkAssociatedTeamOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateTeamworkAssociatedTeamOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateTeamworkAssociatedTeam - Update the navigation property associatedTeams in users
func (c TeamworkAssociatedTeamClient) UpdateTeamworkAssociatedTeam(ctx context.Context, id stable.UserIdTeamworkAssociatedTeamId, input stable.AssociatedTeamInfo, options UpdateTeamworkAssociatedTeamOperationOptions) (result UpdateTeamworkAssociatedTeamOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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
