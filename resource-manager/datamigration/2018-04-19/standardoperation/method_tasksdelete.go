package standardoperation

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TasksDeleteOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type TasksDeleteOperationOptions struct {
	DeleteRunningTasks *bool
}

func DefaultTasksDeleteOperationOptions() TasksDeleteOperationOptions {
	return TasksDeleteOperationOptions{}
}

func (o TasksDeleteOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o TasksDeleteOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o TasksDeleteOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.DeleteRunningTasks != nil {
		out.Append("deleteRunningTasks", fmt.Sprintf("%v", *o.DeleteRunningTasks))
	}
	return &out
}

// TasksDelete ...
func (c StandardOperationClient) TasksDelete(ctx context.Context, id TaskId, options TasksDeleteOperationOptions) (result TasksDeleteOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          id.ID(),
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
