package trafficfilter

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DetachAndDeleteTrafficFilterDeleteOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DetachAndDeleteTrafficFilterDeleteOperationOptions struct {
	RulesetId *string
}

func DefaultDetachAndDeleteTrafficFilterDeleteOperationOptions() DetachAndDeleteTrafficFilterDeleteOperationOptions {
	return DetachAndDeleteTrafficFilterDeleteOperationOptions{}
}

func (o DetachAndDeleteTrafficFilterDeleteOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o DetachAndDeleteTrafficFilterDeleteOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DetachAndDeleteTrafficFilterDeleteOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.RulesetId != nil {
		out.Append("rulesetId", fmt.Sprintf("%v", *o.RulesetId))
	}
	return &out
}

// DetachAndDeleteTrafficFilterDelete ...
func (c TrafficFilterClient) DetachAndDeleteTrafficFilterDelete(ctx context.Context, id MonitorId, options DetachAndDeleteTrafficFilterDeleteOperationOptions) (result DetachAndDeleteTrafficFilterDeleteOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/detachAndDeleteTrafficFilter", id.ID()),
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
