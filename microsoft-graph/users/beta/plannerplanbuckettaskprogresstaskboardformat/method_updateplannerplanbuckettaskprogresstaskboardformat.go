package plannerplanbuckettaskprogresstaskboardformat

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

type UpdatePlannerPlanBucketTaskProgressTaskBoardFormatOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdatePlannerPlanBucketTaskProgressTaskBoardFormatOperationOptions struct {
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdatePlannerPlanBucketTaskProgressTaskBoardFormatOperationOptions() UpdatePlannerPlanBucketTaskProgressTaskBoardFormatOperationOptions {
	return UpdatePlannerPlanBucketTaskProgressTaskBoardFormatOperationOptions{}
}

func (o UpdatePlannerPlanBucketTaskProgressTaskBoardFormatOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o UpdatePlannerPlanBucketTaskProgressTaskBoardFormatOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdatePlannerPlanBucketTaskProgressTaskBoardFormatOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdatePlannerPlanBucketTaskProgressTaskBoardFormat - Update the navigation property progressTaskBoardFormat in users
func (c PlannerPlanBucketTaskProgressTaskBoardFormatClient) UpdatePlannerPlanBucketTaskProgressTaskBoardFormat(ctx context.Context, id beta.UserIdPlannerPlanIdBucketIdTaskId, input beta.PlannerProgressTaskBoardTaskFormat, options UpdatePlannerPlanBucketTaskProgressTaskBoardFormatOperationOptions) (result UpdatePlannerPlanBucketTaskProgressTaskBoardFormatOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/progressTaskBoardFormat", id.ID()),
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
