package billingsubscription

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByBillingAccountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]BillingSubscription
}

type ListByBillingAccountCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []BillingSubscription
}

type ListByBillingAccountOperationOptions struct {
	Count                      *bool
	Expand                     *string
	Filter                     *string
	IncludeDeleted             *bool
	IncludeFailed              *bool
	IncludeTenantSubscriptions *bool
	OrderBy                    *string
	Search                     *string
	Skip                       *int64
	Top                        *int64
}

func DefaultListByBillingAccountOperationOptions() ListByBillingAccountOperationOptions {
	return ListByBillingAccountOperationOptions{}
}

func (o ListByBillingAccountOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListByBillingAccountOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ListByBillingAccountOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Count != nil {
		out.Append("count", fmt.Sprintf("%v", *o.Count))
	}
	if o.Expand != nil {
		out.Append("expand", fmt.Sprintf("%v", *o.Expand))
	}
	if o.Filter != nil {
		out.Append("filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.IncludeDeleted != nil {
		out.Append("includeDeleted", fmt.Sprintf("%v", *o.IncludeDeleted))
	}
	if o.IncludeFailed != nil {
		out.Append("includeFailed", fmt.Sprintf("%v", *o.IncludeFailed))
	}
	if o.IncludeTenantSubscriptions != nil {
		out.Append("includeTenantSubscriptions", fmt.Sprintf("%v", *o.IncludeTenantSubscriptions))
	}
	if o.OrderBy != nil {
		out.Append("orderBy", fmt.Sprintf("%v", *o.OrderBy))
	}
	if o.Search != nil {
		out.Append("search", fmt.Sprintf("%v", *o.Search))
	}
	if o.Skip != nil {
		out.Append("skip", fmt.Sprintf("%v", *o.Skip))
	}
	if o.Top != nil {
		out.Append("top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

type ListByBillingAccountCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByBillingAccountCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByBillingAccount ...
func (c BillingSubscriptionClient) ListByBillingAccount(ctx context.Context, id BillingAccountId, options ListByBillingAccountOperationOptions) (result ListByBillingAccountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListByBillingAccountCustomPager{},
		Path:          fmt.Sprintf("%s/billingSubscriptions", id.ID()),
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
		Values *[]BillingSubscription `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByBillingAccountComplete retrieves all the results into a single object
func (c BillingSubscriptionClient) ListByBillingAccountComplete(ctx context.Context, id BillingAccountId, options ListByBillingAccountOperationOptions) (ListByBillingAccountCompleteResult, error) {
	return c.ListByBillingAccountCompleteMatchingPredicate(ctx, id, options, BillingSubscriptionOperationPredicate{})
}

// ListByBillingAccountCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c BillingSubscriptionClient) ListByBillingAccountCompleteMatchingPredicate(ctx context.Context, id BillingAccountId, options ListByBillingAccountOperationOptions, predicate BillingSubscriptionOperationPredicate) (result ListByBillingAccountCompleteResult, err error) {
	items := make([]BillingSubscription, 0)

	resp, err := c.ListByBillingAccount(ctx, id, options)
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

	result = ListByBillingAccountCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
