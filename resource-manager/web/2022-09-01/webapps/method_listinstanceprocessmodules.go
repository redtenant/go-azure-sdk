package webapps

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListInstanceProcessModulesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ProcessModuleInfo
}

type ListInstanceProcessModulesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ProcessModuleInfo
}

type ListInstanceProcessModulesCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListInstanceProcessModulesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListInstanceProcessModules ...
func (c WebAppsClient) ListInstanceProcessModules(ctx context.Context, id InstanceProcessId) (result ListInstanceProcessModulesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListInstanceProcessModulesCustomPager{},
		Path:       fmt.Sprintf("%s/modules", id.ID()),
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
		Values *[]ProcessModuleInfo `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListInstanceProcessModulesComplete retrieves all the results into a single object
func (c WebAppsClient) ListInstanceProcessModulesComplete(ctx context.Context, id InstanceProcessId) (ListInstanceProcessModulesCompleteResult, error) {
	return c.ListInstanceProcessModulesCompleteMatchingPredicate(ctx, id, ProcessModuleInfoOperationPredicate{})
}

// ListInstanceProcessModulesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c WebAppsClient) ListInstanceProcessModulesCompleteMatchingPredicate(ctx context.Context, id InstanceProcessId, predicate ProcessModuleInfoOperationPredicate) (result ListInstanceProcessModulesCompleteResult, err error) {
	items := make([]ProcessModuleInfo, 0)

	resp, err := c.ListInstanceProcessModules(ctx, id)
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

	result = ListInstanceProcessModulesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
