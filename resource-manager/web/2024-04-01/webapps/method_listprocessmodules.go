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

type ListProcessModulesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ProcessModuleInfo
}

type ListProcessModulesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ProcessModuleInfo
}

type ListProcessModulesCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListProcessModulesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListProcessModules ...
func (c WebAppsClient) ListProcessModules(ctx context.Context, id ProcessId) (result ListProcessModulesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListProcessModulesCustomPager{},
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

// ListProcessModulesComplete retrieves all the results into a single object
func (c WebAppsClient) ListProcessModulesComplete(ctx context.Context, id ProcessId) (ListProcessModulesCompleteResult, error) {
	return c.ListProcessModulesCompleteMatchingPredicate(ctx, id, ProcessModuleInfoOperationPredicate{})
}

// ListProcessModulesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c WebAppsClient) ListProcessModulesCompleteMatchingPredicate(ctx context.Context, id ProcessId, predicate ProcessModuleInfoOperationPredicate) (result ListProcessModulesCompleteResult, err error) {
	items := make([]ProcessModuleInfo, 0)

	resp, err := c.ListProcessModules(ctx, id)
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

	result = ListProcessModulesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
