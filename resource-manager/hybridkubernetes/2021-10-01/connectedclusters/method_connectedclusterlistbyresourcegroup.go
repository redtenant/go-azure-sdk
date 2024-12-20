package connectedclusters

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

type ConnectedClusterListByResourceGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ConnectedCluster
}

type ConnectedClusterListByResourceGroupCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ConnectedCluster
}

type ConnectedClusterListByResourceGroupCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ConnectedClusterListByResourceGroupCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ConnectedClusterListByResourceGroup ...
func (c ConnectedClustersClient) ConnectedClusterListByResourceGroup(ctx context.Context, id commonids.ResourceGroupId) (result ConnectedClusterListByResourceGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ConnectedClusterListByResourceGroupCustomPager{},
		Path:       fmt.Sprintf("%s/providers/Microsoft.Kubernetes/connectedClusters", id.ID()),
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
		Values *[]ConnectedCluster `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ConnectedClusterListByResourceGroupComplete retrieves all the results into a single object
func (c ConnectedClustersClient) ConnectedClusterListByResourceGroupComplete(ctx context.Context, id commonids.ResourceGroupId) (ConnectedClusterListByResourceGroupCompleteResult, error) {
	return c.ConnectedClusterListByResourceGroupCompleteMatchingPredicate(ctx, id, ConnectedClusterOperationPredicate{})
}

// ConnectedClusterListByResourceGroupCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ConnectedClustersClient) ConnectedClusterListByResourceGroupCompleteMatchingPredicate(ctx context.Context, id commonids.ResourceGroupId, predicate ConnectedClusterOperationPredicate) (result ConnectedClusterListByResourceGroupCompleteResult, err error) {
	items := make([]ConnectedCluster, 0)

	resp, err := c.ConnectedClusterListByResourceGroup(ctx, id)
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

	result = ConnectedClusterListByResourceGroupCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
