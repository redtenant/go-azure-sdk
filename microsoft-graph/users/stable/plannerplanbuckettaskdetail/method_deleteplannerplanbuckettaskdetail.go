package plannerplanbuckettaskdetail

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

type DeletePlannerPlanBucketTaskDetailOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeletePlannerPlanBucketTaskDetailOperationOptions struct {
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultDeletePlannerPlanBucketTaskDetailOperationOptions() DeletePlannerPlanBucketTaskDetailOperationOptions {
	return DeletePlannerPlanBucketTaskDetailOperationOptions{}
}

func (o DeletePlannerPlanBucketTaskDetailOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeletePlannerPlanBucketTaskDetailOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeletePlannerPlanBucketTaskDetailOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeletePlannerPlanBucketTaskDetail - Delete navigation property details for users
func (c PlannerPlanBucketTaskDetailClient) DeletePlannerPlanBucketTaskDetail(ctx context.Context, id stable.UserIdPlannerPlanIdBucketIdTaskId, options DeletePlannerPlanBucketTaskDetailOperationOptions) (result DeletePlannerPlanBucketTaskDetailOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/details", id.ID()),
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
