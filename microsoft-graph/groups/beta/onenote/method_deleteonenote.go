package onenote

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

type DeleteOnenoteOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteOnenoteOperationOptions struct {
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultDeleteOnenoteOperationOptions() DeleteOnenoteOperationOptions {
	return DeleteOnenoteOperationOptions{}
}

func (o DeleteOnenoteOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteOnenoteOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeleteOnenoteOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteOnenote - Delete navigation property onenote for groups
func (c OnenoteClient) DeleteOnenote(ctx context.Context, id beta.GroupId, options DeleteOnenoteOperationOptions) (result DeleteOnenoteOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/onenote", id.ID()),
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
