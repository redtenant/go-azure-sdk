package joinedteamscheduleopenshift

import (
	"context"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateJoinedTeamScheduleOpenShiftOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateJoinedTeamScheduleOpenShiftOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateJoinedTeamScheduleOpenShiftOperationOptions() UpdateJoinedTeamScheduleOpenShiftOperationOptions {
	return UpdateJoinedTeamScheduleOpenShiftOperationOptions{}
}

func (o UpdateJoinedTeamScheduleOpenShiftOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateJoinedTeamScheduleOpenShiftOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateJoinedTeamScheduleOpenShiftOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateJoinedTeamScheduleOpenShift - Update the navigation property openShifts in users
func (c JoinedTeamScheduleOpenShiftClient) UpdateJoinedTeamScheduleOpenShift(ctx context.Context, id stable.UserIdJoinedTeamIdScheduleOpenShiftId, input stable.OpenShift, options UpdateJoinedTeamScheduleOpenShiftOperationOptions) (result UpdateJoinedTeamScheduleOpenShiftOperationResponse, err error) {
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
