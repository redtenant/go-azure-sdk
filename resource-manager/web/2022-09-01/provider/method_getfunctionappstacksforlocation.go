package provider

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetFunctionAppStacksForLocationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]FunctionAppStack
}

type GetFunctionAppStacksForLocationCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []FunctionAppStack
}

type GetFunctionAppStacksForLocationOperationOptions struct {
	StackOsType *ProviderStackOsType
}

func DefaultGetFunctionAppStacksForLocationOperationOptions() GetFunctionAppStacksForLocationOperationOptions {
	return GetFunctionAppStacksForLocationOperationOptions{}
}

func (o GetFunctionAppStacksForLocationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetFunctionAppStacksForLocationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o GetFunctionAppStacksForLocationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.StackOsType != nil {
		out.Append("stackOsType", fmt.Sprintf("%v", *o.StackOsType))
	}
	return &out
}

type GetFunctionAppStacksForLocationCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *GetFunctionAppStacksForLocationCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// GetFunctionAppStacksForLocation ...
func (c ProviderClient) GetFunctionAppStacksForLocation(ctx context.Context, id LocationId, options GetFunctionAppStacksForLocationOperationOptions) (result GetFunctionAppStacksForLocationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &GetFunctionAppStacksForLocationCustomPager{},
		Path:          fmt.Sprintf("%s/functionAppStacks", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]FunctionAppStack `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// GetFunctionAppStacksForLocationComplete retrieves all the results into a single object
func (c ProviderClient) GetFunctionAppStacksForLocationComplete(ctx context.Context, id LocationId, options GetFunctionAppStacksForLocationOperationOptions) (GetFunctionAppStacksForLocationCompleteResult, error) {
	return c.GetFunctionAppStacksForLocationCompleteMatchingPredicate(ctx, id, options, FunctionAppStackOperationPredicate{})
}

// GetFunctionAppStacksForLocationCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ProviderClient) GetFunctionAppStacksForLocationCompleteMatchingPredicate(ctx context.Context, id LocationId, options GetFunctionAppStacksForLocationOperationOptions, predicate FunctionAppStackOperationPredicate) (result GetFunctionAppStacksForLocationCompleteResult, err error) {
	items := make([]FunctionAppStack, 0)

	resp, err := c.GetFunctionAppStacksForLocation(ctx, id, options)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = GetFunctionAppStacksForLocationCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
