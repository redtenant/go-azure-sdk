package managedenvironments

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedEnvironmentPrivateEndpointConnectionsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]PrivateEndpointConnection
}

type ManagedEnvironmentPrivateEndpointConnectionsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []PrivateEndpointConnection
}

type ManagedEnvironmentPrivateEndpointConnectionsListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ManagedEnvironmentPrivateEndpointConnectionsListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ManagedEnvironmentPrivateEndpointConnectionsList ...
func (c ManagedEnvironmentsClient) ManagedEnvironmentPrivateEndpointConnectionsList(ctx context.Context, id ManagedEnvironmentId) (result ManagedEnvironmentPrivateEndpointConnectionsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ManagedEnvironmentPrivateEndpointConnectionsListCustomPager{},
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

// ManagedEnvironmentPrivateEndpointConnectionsListComplete retrieves all the results into a single object
func (c ManagedEnvironmentsClient) ManagedEnvironmentPrivateEndpointConnectionsListComplete(ctx context.Context, id ManagedEnvironmentId) (ManagedEnvironmentPrivateEndpointConnectionsListCompleteResult, error) {
	return c.ManagedEnvironmentPrivateEndpointConnectionsListCompleteMatchingPredicate(ctx, id, PrivateEndpointConnectionOperationPredicate{})
}

// ManagedEnvironmentPrivateEndpointConnectionsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ManagedEnvironmentsClient) ManagedEnvironmentPrivateEndpointConnectionsListCompleteMatchingPredicate(ctx context.Context, id ManagedEnvironmentId, predicate PrivateEndpointConnectionOperationPredicate) (result ManagedEnvironmentPrivateEndpointConnectionsListCompleteResult, err error) {
	items := make([]PrivateEndpointConnection, 0)

	resp, err := c.ManagedEnvironmentPrivateEndpointConnectionsList(ctx, id)
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

	result = ManagedEnvironmentPrivateEndpointConnectionsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
