package migrates

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

type HyperVSitesListBySubscriptionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]HyperVSite
}

type HyperVSitesListBySubscriptionCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []HyperVSite
}

type HyperVSitesListBySubscriptionCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *HyperVSitesListBySubscriptionCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// HyperVSitesListBySubscription ...
func (c MigratesClient) HyperVSitesListBySubscription(ctx context.Context, id commonids.SubscriptionId) (result HyperVSitesListBySubscriptionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &HyperVSitesListBySubscriptionCustomPager{},
		Path:       fmt.Sprintf("%s/providers/Microsoft.OffAzure/hyperVSites", id.ID()),
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
		Values *[]HyperVSite `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// HyperVSitesListBySubscriptionComplete retrieves all the results into a single object
func (c MigratesClient) HyperVSitesListBySubscriptionComplete(ctx context.Context, id commonids.SubscriptionId) (HyperVSitesListBySubscriptionCompleteResult, error) {
	return c.HyperVSitesListBySubscriptionCompleteMatchingPredicate(ctx, id, HyperVSiteOperationPredicate{})
}

// HyperVSitesListBySubscriptionCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MigratesClient) HyperVSitesListBySubscriptionCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, predicate HyperVSiteOperationPredicate) (result HyperVSitesListBySubscriptionCompleteResult, err error) {
	items := make([]HyperVSite, 0)

	resp, err := c.HyperVSitesListBySubscription(ctx, id)
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

	result = HyperVSitesListBySubscriptionCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
