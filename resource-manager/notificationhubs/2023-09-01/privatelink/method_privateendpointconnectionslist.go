package privatelink

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateEndpointConnectionsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]PrivateEndpointConnectionResource
}

type PrivateEndpointConnectionsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []PrivateEndpointConnectionResource
}

type PrivateEndpointConnectionsListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *PrivateEndpointConnectionsListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// PrivateEndpointConnectionsList ...
func (c PrivateLinkClient) PrivateEndpointConnectionsList(ctx context.Context, id NamespaceId) (result PrivateEndpointConnectionsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &PrivateEndpointConnectionsListCustomPager{},
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
		Values *[]PrivateEndpointConnectionResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// PrivateEndpointConnectionsListComplete retrieves all the results into a single object
func (c PrivateLinkClient) PrivateEndpointConnectionsListComplete(ctx context.Context, id NamespaceId) (PrivateEndpointConnectionsListCompleteResult, error) {
	return c.PrivateEndpointConnectionsListCompleteMatchingPredicate(ctx, id, PrivateEndpointConnectionResourceOperationPredicate{})
}

// PrivateEndpointConnectionsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PrivateLinkClient) PrivateEndpointConnectionsListCompleteMatchingPredicate(ctx context.Context, id NamespaceId, predicate PrivateEndpointConnectionResourceOperationPredicate) (result PrivateEndpointConnectionsListCompleteResult, err error) {
	items := make([]PrivateEndpointConnectionResource, 0)

	resp, err := c.PrivateEndpointConnectionsList(ctx, id)
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

	result = PrivateEndpointConnectionsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
