package eventsubscriptions

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

type ListGlobalByResourceGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]EventSubscription
}

type ListGlobalByResourceGroupCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []EventSubscription
}

type ListGlobalByResourceGroupOperationOptions struct {
	Filter *string
	Top    *int64
}

func DefaultListGlobalByResourceGroupOperationOptions() ListGlobalByResourceGroupOperationOptions {
	return ListGlobalByResourceGroupOperationOptions{}
}

func (o ListGlobalByResourceGroupOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListGlobalByResourceGroupOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ListGlobalByResourceGroupOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

type ListGlobalByResourceGroupCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListGlobalByResourceGroupCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListGlobalByResourceGroup ...
func (c EventSubscriptionsClient) ListGlobalByResourceGroup(ctx context.Context, id commonids.ResourceGroupId, options ListGlobalByResourceGroupOperationOptions) (result ListGlobalByResourceGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListGlobalByResourceGroupCustomPager{},
		Path:          fmt.Sprintf("%s/providers/Microsoft.EventGrid/eventSubscriptions", id.ID()),
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
		Values *[]EventSubscription `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGlobalByResourceGroupComplete retrieves all the results into a single object
func (c EventSubscriptionsClient) ListGlobalByResourceGroupComplete(ctx context.Context, id commonids.ResourceGroupId, options ListGlobalByResourceGroupOperationOptions) (ListGlobalByResourceGroupCompleteResult, error) {
	return c.ListGlobalByResourceGroupCompleteMatchingPredicate(ctx, id, options, EventSubscriptionOperationPredicate{})
}

// ListGlobalByResourceGroupCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EventSubscriptionsClient) ListGlobalByResourceGroupCompleteMatchingPredicate(ctx context.Context, id commonids.ResourceGroupId, options ListGlobalByResourceGroupOperationOptions, predicate EventSubscriptionOperationPredicate) (result ListGlobalByResourceGroupCompleteResult, err error) {
	items := make([]EventSubscription, 0)

	resp, err := c.ListGlobalByResourceGroup(ctx, id, options)
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

	result = ListGlobalByResourceGroupCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
