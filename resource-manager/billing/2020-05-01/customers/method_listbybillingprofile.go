package customers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByBillingProfileOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Customer
}

type ListByBillingProfileCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Customer
}

type ListByBillingProfileOperationOptions struct {
	Filter *string
	Search *string
}

func DefaultListByBillingProfileOperationOptions() ListByBillingProfileOperationOptions {
	return ListByBillingProfileOperationOptions{}
}

func (o ListByBillingProfileOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListByBillingProfileOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ListByBillingProfileOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.Search != nil {
		out.Append("$search", fmt.Sprintf("%v", *o.Search))
	}
	return &out
}

type ListByBillingProfileCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByBillingProfileCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByBillingProfile ...
func (c CustomersClient) ListByBillingProfile(ctx context.Context, id BillingProfileId, options ListByBillingProfileOperationOptions) (result ListByBillingProfileOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListByBillingProfileCustomPager{},
		Path:          fmt.Sprintf("%s/customers", id.ID()),
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
		Values *[]Customer `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByBillingProfileComplete retrieves all the results into a single object
func (c CustomersClient) ListByBillingProfileComplete(ctx context.Context, id BillingProfileId, options ListByBillingProfileOperationOptions) (ListByBillingProfileCompleteResult, error) {
	return c.ListByBillingProfileCompleteMatchingPredicate(ctx, id, options, CustomerOperationPredicate{})
}

// ListByBillingProfileCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CustomersClient) ListByBillingProfileCompleteMatchingPredicate(ctx context.Context, id BillingProfileId, options ListByBillingProfileOperationOptions, predicate CustomerOperationPredicate) (result ListByBillingProfileCompleteResult, err error) {
	items := make([]Customer, 0)

	resp, err := c.ListByBillingProfile(ctx, id, options)
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

	result = ListByBillingProfileCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
