package image

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByLabPlanOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Image
}

type ListByLabPlanCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Image
}

type ListByLabPlanCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByLabPlanCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByLabPlan ...
func (c ImageClient) ListByLabPlan(ctx context.Context, id LabPlanId) (result ListByLabPlanOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListByLabPlanCustomPager{},
		Path:       fmt.Sprintf("%s/images", id.ID()),
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
		Values *[]Image `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByLabPlanComplete retrieves all the results into a single object
func (c ImageClient) ListByLabPlanComplete(ctx context.Context, id LabPlanId) (ListByLabPlanCompleteResult, error) {
	return c.ListByLabPlanCompleteMatchingPredicate(ctx, id, ImageOperationPredicate{})
}

// ListByLabPlanCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ImageClient) ListByLabPlanCompleteMatchingPredicate(ctx context.Context, id LabPlanId, predicate ImageOperationPredicate) (result ListByLabPlanCompleteResult, err error) {
	items := make([]Image, 0)

	resp, err := c.ListByLabPlan(ctx, id)
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

	result = ListByLabPlanCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
