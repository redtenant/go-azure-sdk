package joinedteamscheduleschedulinggroup

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

type CreateJoinedTeamScheduleSchedulingGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.SchedulingGroup
}

type CreateJoinedTeamScheduleSchedulingGroupOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateJoinedTeamScheduleSchedulingGroupOperationOptions() CreateJoinedTeamScheduleSchedulingGroupOperationOptions {
	return CreateJoinedTeamScheduleSchedulingGroupOperationOptions{}
}

func (o CreateJoinedTeamScheduleSchedulingGroupOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateJoinedTeamScheduleSchedulingGroupOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateJoinedTeamScheduleSchedulingGroupOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateJoinedTeamScheduleSchedulingGroup - Create new navigation property to schedulingGroups for me
func (c JoinedTeamScheduleSchedulingGroupClient) CreateJoinedTeamScheduleSchedulingGroup(ctx context.Context, id stable.MeJoinedTeamId, input stable.SchedulingGroup, options CreateJoinedTeamScheduleSchedulingGroupOperationOptions) (result CreateJoinedTeamScheduleSchedulingGroupOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/schedule/schedulingGroups", id.ID()),
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

	var model stable.SchedulingGroup
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
