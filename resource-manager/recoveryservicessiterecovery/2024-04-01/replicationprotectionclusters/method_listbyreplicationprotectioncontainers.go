package replicationprotectionclusters

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByReplicationProtectionContainersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ReplicationProtectionCluster
}

type ListByReplicationProtectionContainersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ReplicationProtectionCluster
}

type ListByReplicationProtectionContainersCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByReplicationProtectionContainersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByReplicationProtectionContainers ...
func (c ReplicationProtectionClustersClient) ListByReplicationProtectionContainers(ctx context.Context, id ReplicationProtectionContainerId) (result ListByReplicationProtectionContainersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListByReplicationProtectionContainersCustomPager{},
		Path:       fmt.Sprintf("%s/replicationProtectionClusters", id.ID()),
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
		Values *[]ReplicationProtectionCluster `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByReplicationProtectionContainersComplete retrieves all the results into a single object
func (c ReplicationProtectionClustersClient) ListByReplicationProtectionContainersComplete(ctx context.Context, id ReplicationProtectionContainerId) (ListByReplicationProtectionContainersCompleteResult, error) {
	return c.ListByReplicationProtectionContainersCompleteMatchingPredicate(ctx, id, ReplicationProtectionClusterOperationPredicate{})
}

// ListByReplicationProtectionContainersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ReplicationProtectionClustersClient) ListByReplicationProtectionContainersCompleteMatchingPredicate(ctx context.Context, id ReplicationProtectionContainerId, predicate ReplicationProtectionClusterOperationPredicate) (result ListByReplicationProtectionContainersCompleteResult, err error) {
	items := make([]ReplicationProtectionCluster, 0)

	resp, err := c.ListByReplicationProtectionContainers(ctx, id)
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

	result = ListByReplicationProtectionContainersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
