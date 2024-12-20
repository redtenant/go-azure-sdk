package endpointservices

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AvailableEndpointServicesListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]EndpointServiceResult
}

type AvailableEndpointServicesListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []EndpointServiceResult
}

type AvailableEndpointServicesListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *AvailableEndpointServicesListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// AvailableEndpointServicesList ...
func (c EndpointServicesClient) AvailableEndpointServicesList(ctx context.Context, id LocationId) (result AvailableEndpointServicesListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &AvailableEndpointServicesListCustomPager{},
		Path:       fmt.Sprintf("%s/virtualNetworkAvailableEndpointServices", id.ID()),
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
		Values *[]EndpointServiceResult `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// AvailableEndpointServicesListComplete retrieves all the results into a single object
func (c EndpointServicesClient) AvailableEndpointServicesListComplete(ctx context.Context, id LocationId) (AvailableEndpointServicesListCompleteResult, error) {
	return c.AvailableEndpointServicesListCompleteMatchingPredicate(ctx, id, EndpointServiceResultOperationPredicate{})
}

// AvailableEndpointServicesListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EndpointServicesClient) AvailableEndpointServicesListCompleteMatchingPredicate(ctx context.Context, id LocationId, predicate EndpointServiceResultOperationPredicate) (result AvailableEndpointServicesListCompleteResult, err error) {
	items := make([]EndpointServiceResult, 0)

	resp, err := c.AvailableEndpointServicesList(ctx, id)
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

	result = AvailableEndpointServicesListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
