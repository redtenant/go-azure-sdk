package azuretrafficcollectors

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

type BySubscriptionListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]AzureTrafficCollector
}

type BySubscriptionListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []AzureTrafficCollector
}

type BySubscriptionListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *BySubscriptionListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// BySubscriptionList ...
func (c AzureTrafficCollectorsClient) BySubscriptionList(ctx context.Context, id commonids.SubscriptionId) (result BySubscriptionListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &BySubscriptionListCustomPager{},
		Path:       fmt.Sprintf("%s/providers/Microsoft.NetworkFunction/azureTrafficCollectors", id.ID()),
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
		Values *[]AzureTrafficCollector `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// BySubscriptionListComplete retrieves all the results into a single object
func (c AzureTrafficCollectorsClient) BySubscriptionListComplete(ctx context.Context, id commonids.SubscriptionId) (BySubscriptionListCompleteResult, error) {
	return c.BySubscriptionListCompleteMatchingPredicate(ctx, id, AzureTrafficCollectorOperationPredicate{})
}

// BySubscriptionListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AzureTrafficCollectorsClient) BySubscriptionListCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, predicate AzureTrafficCollectorOperationPredicate) (result BySubscriptionListCompleteResult, err error) {
	items := make([]AzureTrafficCollector, 0)

	resp, err := c.BySubscriptionList(ctx, id)
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

	result = BySubscriptionListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
