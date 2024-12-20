package sharesubscription

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListSynchronizationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ShareSubscriptionSynchronization
}

type ListSynchronizationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ShareSubscriptionSynchronization
}

type ListSynchronizationsOperationOptions struct {
	Filter  *string
	Orderby *string
}

func DefaultListSynchronizationsOperationOptions() ListSynchronizationsOperationOptions {
	return ListSynchronizationsOperationOptions{}
}

func (o ListSynchronizationsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListSynchronizationsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ListSynchronizationsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.Orderby != nil {
		out.Append("$orderby", fmt.Sprintf("%v", *o.Orderby))
	}
	return &out
}

type ListSynchronizationsCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListSynchronizationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSynchronizations ...
func (c ShareSubscriptionClient) ListSynchronizations(ctx context.Context, id ShareSubscriptionId, options ListSynchronizationsOperationOptions) (result ListSynchronizationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &ListSynchronizationsCustomPager{},
		Path:          fmt.Sprintf("%s/listSynchronizations", id.ID()),
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
		Values *[]ShareSubscriptionSynchronization `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSynchronizationsComplete retrieves all the results into a single object
func (c ShareSubscriptionClient) ListSynchronizationsComplete(ctx context.Context, id ShareSubscriptionId, options ListSynchronizationsOperationOptions) (ListSynchronizationsCompleteResult, error) {
	return c.ListSynchronizationsCompleteMatchingPredicate(ctx, id, options, ShareSubscriptionSynchronizationOperationPredicate{})
}

// ListSynchronizationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ShareSubscriptionClient) ListSynchronizationsCompleteMatchingPredicate(ctx context.Context, id ShareSubscriptionId, options ListSynchronizationsOperationOptions, predicate ShareSubscriptionSynchronizationOperationPredicate) (result ListSynchronizationsCompleteResult, err error) {
	items := make([]ShareSubscriptionSynchronization, 0)

	resp, err := c.ListSynchronizations(ctx, id, options)
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

	result = ListSynchronizationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
