package recommendations

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResetAllFiltersForHostingEnvironmentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type ResetAllFiltersForHostingEnvironmentOperationOptions struct {
	EnvironmentName *string
}

func DefaultResetAllFiltersForHostingEnvironmentOperationOptions() ResetAllFiltersForHostingEnvironmentOperationOptions {
	return ResetAllFiltersForHostingEnvironmentOperationOptions{}
}

func (o ResetAllFiltersForHostingEnvironmentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ResetAllFiltersForHostingEnvironmentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ResetAllFiltersForHostingEnvironmentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.EnvironmentName != nil {
		out.Append("environmentName", fmt.Sprintf("%v", *o.EnvironmentName))
	}
	return &out
}

// ResetAllFiltersForHostingEnvironment ...
func (c RecommendationsClient) ResetAllFiltersForHostingEnvironment(ctx context.Context, id commonids.AppServiceEnvironmentId, options ResetAllFiltersForHostingEnvironmentOperationOptions) (result ResetAllFiltersForHostingEnvironmentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/recommendations/reset", id.ID()),
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
