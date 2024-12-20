package productpolicy

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspaceProductPolicyListByProductOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]PolicyContract
}

type WorkspaceProductPolicyListByProductCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []PolicyContract
}

type WorkspaceProductPolicyListByProductCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *WorkspaceProductPolicyListByProductCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// WorkspaceProductPolicyListByProduct ...
func (c ProductPolicyClient) WorkspaceProductPolicyListByProduct(ctx context.Context, id WorkspaceProductId) (result WorkspaceProductPolicyListByProductOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &WorkspaceProductPolicyListByProductCustomPager{},
		Path:       fmt.Sprintf("%s/policies", id.ID()),
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
		Values *[]PolicyContract `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// WorkspaceProductPolicyListByProductComplete retrieves all the results into a single object
func (c ProductPolicyClient) WorkspaceProductPolicyListByProductComplete(ctx context.Context, id WorkspaceProductId) (WorkspaceProductPolicyListByProductCompleteResult, error) {
	return c.WorkspaceProductPolicyListByProductCompleteMatchingPredicate(ctx, id, PolicyContractOperationPredicate{})
}

// WorkspaceProductPolicyListByProductCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ProductPolicyClient) WorkspaceProductPolicyListByProductCompleteMatchingPredicate(ctx context.Context, id WorkspaceProductId, predicate PolicyContractOperationPredicate) (result WorkspaceProductPolicyListByProductCompleteResult, err error) {
	items := make([]PolicyContract, 0)

	resp, err := c.WorkspaceProductPolicyListByProduct(ctx, id)
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

	result = WorkspaceProductPolicyListByProductCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
