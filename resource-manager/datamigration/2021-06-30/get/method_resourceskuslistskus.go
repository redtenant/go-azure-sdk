package get

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

type ResourceSkusListSkusOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ResourceSku
}

type ResourceSkusListSkusCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ResourceSku
}

type ResourceSkusListSkusCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ResourceSkusListSkusCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ResourceSkusListSkus ...
func (c GETClient) ResourceSkusListSkus(ctx context.Context, id commonids.SubscriptionId) (result ResourceSkusListSkusOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ResourceSkusListSkusCustomPager{},
		Path:       fmt.Sprintf("%s/providers/Microsoft.DataMigration/skus", id.ID()),
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
		Values *[]ResourceSku `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ResourceSkusListSkusComplete retrieves all the results into a single object
func (c GETClient) ResourceSkusListSkusComplete(ctx context.Context, id commonids.SubscriptionId) (ResourceSkusListSkusCompleteResult, error) {
	return c.ResourceSkusListSkusCompleteMatchingPredicate(ctx, id, ResourceSkuOperationPredicate{})
}

// ResourceSkusListSkusCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GETClient) ResourceSkusListSkusCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, predicate ResourceSkuOperationPredicate) (result ResourceSkusListSkusCompleteResult, err error) {
	items := make([]ResourceSku, 0)

	resp, err := c.ResourceSkusListSkus(ctx, id)
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

	result = ResourceSkusListSkusCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
