package eventsubscriptions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PartnerTopicEventSubscriptionsListByPartnerTopicOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]EventSubscription
}

type PartnerTopicEventSubscriptionsListByPartnerTopicCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []EventSubscription
}

type PartnerTopicEventSubscriptionsListByPartnerTopicOperationOptions struct {
	Filter *string
	Top    *int64
}

func DefaultPartnerTopicEventSubscriptionsListByPartnerTopicOperationOptions() PartnerTopicEventSubscriptionsListByPartnerTopicOperationOptions {
	return PartnerTopicEventSubscriptionsListByPartnerTopicOperationOptions{}
}

func (o PartnerTopicEventSubscriptionsListByPartnerTopicOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o PartnerTopicEventSubscriptionsListByPartnerTopicOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o PartnerTopicEventSubscriptionsListByPartnerTopicOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

type PartnerTopicEventSubscriptionsListByPartnerTopicCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *PartnerTopicEventSubscriptionsListByPartnerTopicCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// PartnerTopicEventSubscriptionsListByPartnerTopic ...
func (c EventSubscriptionsClient) PartnerTopicEventSubscriptionsListByPartnerTopic(ctx context.Context, id PartnerTopicId, options PartnerTopicEventSubscriptionsListByPartnerTopicOperationOptions) (result PartnerTopicEventSubscriptionsListByPartnerTopicOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &PartnerTopicEventSubscriptionsListByPartnerTopicCustomPager{},
		Path:          fmt.Sprintf("%s/eventSubscriptions", id.ID()),
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

// PartnerTopicEventSubscriptionsListByPartnerTopicComplete retrieves all the results into a single object
func (c EventSubscriptionsClient) PartnerTopicEventSubscriptionsListByPartnerTopicComplete(ctx context.Context, id PartnerTopicId, options PartnerTopicEventSubscriptionsListByPartnerTopicOperationOptions) (PartnerTopicEventSubscriptionsListByPartnerTopicCompleteResult, error) {
	return c.PartnerTopicEventSubscriptionsListByPartnerTopicCompleteMatchingPredicate(ctx, id, options, EventSubscriptionOperationPredicate{})
}

// PartnerTopicEventSubscriptionsListByPartnerTopicCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EventSubscriptionsClient) PartnerTopicEventSubscriptionsListByPartnerTopicCompleteMatchingPredicate(ctx context.Context, id PartnerTopicId, options PartnerTopicEventSubscriptionsListByPartnerTopicOperationOptions, predicate EventSubscriptionOperationPredicate) (result PartnerTopicEventSubscriptionsListByPartnerTopicCompleteResult, err error) {
	items := make([]EventSubscription, 0)

	resp, err := c.PartnerTopicEventSubscriptionsListByPartnerTopic(ctx, id, options)
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

	result = PartnerTopicEventSubscriptionsListByPartnerTopicCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
