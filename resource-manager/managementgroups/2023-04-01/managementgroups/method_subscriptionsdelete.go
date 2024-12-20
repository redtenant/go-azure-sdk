package managementgroups

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubscriptionsDeleteOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SubscriptionsDeleteOperationOptions struct {
	CacheControl *string
}

func DefaultSubscriptionsDeleteOperationOptions() SubscriptionsDeleteOperationOptions {
	return SubscriptionsDeleteOperationOptions{}
}

func (o SubscriptionsDeleteOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.CacheControl != nil {
		out.Append("Cache-Control", fmt.Sprintf("%v", *o.CacheControl))
	}
	return &out
}

func (o SubscriptionsDeleteOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o SubscriptionsDeleteOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SubscriptionsDelete ...
func (c ManagementGroupsClient) SubscriptionsDelete(ctx context.Context, id SubscriptionId, options SubscriptionsDeleteOperationOptions) (result SubscriptionsDeleteOperationResponse, err error) {
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
