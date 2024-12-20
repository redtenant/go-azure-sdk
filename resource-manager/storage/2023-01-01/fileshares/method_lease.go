package fileshares

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LeaseOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *LeaseShareResponse
}

type LeaseOperationOptions struct {
	XMsSnapshot *string
}

func DefaultLeaseOperationOptions() LeaseOperationOptions {
	return LeaseOperationOptions{}
}

func (o LeaseOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.XMsSnapshot != nil {
		out.Append("x-ms-snapshot", fmt.Sprintf("%v", *o.XMsSnapshot))
	}
	return &out
}

func (o LeaseOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o LeaseOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// Lease ...
func (c FileSharesClient) Lease(ctx context.Context, id ShareId, input LeaseShareRequest, options LeaseOperationOptions) (result LeaseOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/lease", id.ID()),
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

	var model LeaseShareResponse
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
