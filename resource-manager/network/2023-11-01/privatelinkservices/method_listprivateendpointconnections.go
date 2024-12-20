package privatelinkservices

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListPrivateEndpointConnectionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]PrivateEndpointConnection
}

type ListPrivateEndpointConnectionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []PrivateEndpointConnection
}

type ListPrivateEndpointConnectionsCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListPrivateEndpointConnectionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListPrivateEndpointConnections ...
func (c PrivateLinkServicesClient) ListPrivateEndpointConnections(ctx context.Context, id PrivateLinkServiceId) (result ListPrivateEndpointConnectionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListPrivateEndpointConnectionsCustomPager{},
		Path:       fmt.Sprintf("%s/privateEndpointConnections", id.ID()),
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
		Values *[]PrivateEndpointConnection `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPrivateEndpointConnectionsComplete retrieves all the results into a single object
func (c PrivateLinkServicesClient) ListPrivateEndpointConnectionsComplete(ctx context.Context, id PrivateLinkServiceId) (ListPrivateEndpointConnectionsCompleteResult, error) {
	return c.ListPrivateEndpointConnectionsCompleteMatchingPredicate(ctx, id, PrivateEndpointConnectionOperationPredicate{})
}

// ListPrivateEndpointConnectionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PrivateLinkServicesClient) ListPrivateEndpointConnectionsCompleteMatchingPredicate(ctx context.Context, id PrivateLinkServiceId, predicate PrivateEndpointConnectionOperationPredicate) (result ListPrivateEndpointConnectionsCompleteResult, err error) {
	items := make([]PrivateEndpointConnection, 0)

	resp, err := c.ListPrivateEndpointConnections(ctx, id)
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

	result = ListPrivateEndpointConnectionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
