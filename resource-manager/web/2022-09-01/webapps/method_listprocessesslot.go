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

type ListProcessesSlotOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ProcessInfo
}

type ListProcessesSlotCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ProcessInfo
}

type ListProcessesSlotCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListProcessesSlotCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListProcessesSlot ...
func (c WebAppsClient) ListProcessesSlot(ctx context.Context, id SlotId) (result ListProcessesSlotOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListProcessesSlotCustomPager{},
		Path:       fmt.Sprintf("%s/processes", id.ID()),
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
		Values *[]ProcessInfo `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListProcessesSlotComplete retrieves all the results into a single object
func (c WebAppsClient) ListProcessesSlotComplete(ctx context.Context, id SlotId) (ListProcessesSlotCompleteResult, error) {
	return c.ListProcessesSlotCompleteMatchingPredicate(ctx, id, ProcessInfoOperationPredicate{})
}

// ListProcessesSlotCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c WebAppsClient) ListProcessesSlotCompleteMatchingPredicate(ctx context.Context, id SlotId, predicate ProcessInfoOperationPredicate) (result ListProcessesSlotCompleteResult, err error) {
	items := make([]ProcessInfo, 0)

	resp, err := c.ListProcessesSlot(ctx, id)
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

	result = ListProcessesSlotCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
