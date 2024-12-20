package operationsinacluster

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

type OperationStatusListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]OperationStatusResult
}

type OperationStatusListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []OperationStatusResult
}

type OperationStatusListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *OperationStatusListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// OperationStatusList ...
func (c OperationsInAClusterClient) OperationStatusList(ctx context.Context, id commonids.ScopeId) (result OperationStatusListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &OperationStatusListCustomPager{},
		Path:       fmt.Sprintf("%s/providers/Microsoft.KubernetesConfiguration/operations", id.ID()),
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
		Values *[]OperationStatusResult `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// OperationStatusListComplete retrieves all the results into a single object
func (c OperationsInAClusterClient) OperationStatusListComplete(ctx context.Context, id commonids.ScopeId) (OperationStatusListCompleteResult, error) {
	return c.OperationStatusListCompleteMatchingPredicate(ctx, id, OperationStatusResultOperationPredicate{})
}

// OperationStatusListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OperationsInAClusterClient) OperationStatusListCompleteMatchingPredicate(ctx context.Context, id commonids.ScopeId, predicate OperationStatusResultOperationPredicate) (result OperationStatusListCompleteResult, err error) {
	items := make([]OperationStatusResult, 0)

	resp, err := c.OperationStatusList(ctx, id)
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

	result = OperationStatusListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
