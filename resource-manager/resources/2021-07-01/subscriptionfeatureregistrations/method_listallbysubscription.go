package subscriptionfeatureregistrations

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

type ListAllBySubscriptionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]SubscriptionFeatureRegistration
}

type ListAllBySubscriptionCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []SubscriptionFeatureRegistration
}

type ListAllBySubscriptionCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListAllBySubscriptionCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAllBySubscription ...
func (c SubscriptionFeatureRegistrationsClient) ListAllBySubscription(ctx context.Context, id commonids.SubscriptionId) (result ListAllBySubscriptionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListAllBySubscriptionCustomPager{},
		Path:       fmt.Sprintf("%s/providers/Microsoft.Features/subscriptionFeatureRegistrations", id.ID()),
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
		Values *[]SubscriptionFeatureRegistration `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAllBySubscriptionComplete retrieves all the results into a single object
func (c SubscriptionFeatureRegistrationsClient) ListAllBySubscriptionComplete(ctx context.Context, id commonids.SubscriptionId) (ListAllBySubscriptionCompleteResult, error) {
	return c.ListAllBySubscriptionCompleteMatchingPredicate(ctx, id, SubscriptionFeatureRegistrationOperationPredicate{})
}

// ListAllBySubscriptionCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SubscriptionFeatureRegistrationsClient) ListAllBySubscriptionCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, predicate SubscriptionFeatureRegistrationOperationPredicate) (result ListAllBySubscriptionCompleteResult, err error) {
	items := make([]SubscriptionFeatureRegistration, 0)

	resp, err := c.ListAllBySubscription(ctx, id)
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

	result = ListAllBySubscriptionCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
