package teamtag

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

type CreateTeamTagOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.TeamworkTag
}

type CreateTeamTagOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateTeamTagOperationOptions() CreateTeamTagOperationOptions {
	return CreateTeamTagOperationOptions{}
}

func (o CreateTeamTagOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateTeamTagOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateTeamTagOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateTeamTag - Create new navigation property to tags for groups
func (c TeamTagClient) CreateTeamTag(ctx context.Context, id beta.GroupId, input beta.TeamworkTag, options CreateTeamTagOperationOptions) (result CreateTeamTagOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/team/tags", id.ID()),
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

	var model beta.TeamworkTag
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
