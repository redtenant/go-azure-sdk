package collection

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

type ServicesListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ServicesDescription
}

type ServicesListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ServicesDescription
}

type ServicesListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ServicesListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ServicesList ...
func (c CollectionClient) ServicesList(ctx context.Context, id commonids.SubscriptionId) (result ServicesListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ServicesListCustomPager{},
		Path:       fmt.Sprintf("%s/providers/Microsoft.HealthcareApis/services", id.ID()),
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
		Values *[]ServicesDescription `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ServicesListComplete retrieves all the results into a single object
func (c CollectionClient) ServicesListComplete(ctx context.Context, id commonids.SubscriptionId) (ServicesListCompleteResult, error) {
	return c.ServicesListCompleteMatchingPredicate(ctx, id, ServicesDescriptionOperationPredicate{})
}

// ServicesListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CollectionClient) ServicesListCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, predicate ServicesDescriptionOperationPredicate) (result ServicesListCompleteResult, err error) {
	items := make([]ServicesDescription, 0)

	resp, err := c.ServicesList(ctx, id)
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

	result = ServicesListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
