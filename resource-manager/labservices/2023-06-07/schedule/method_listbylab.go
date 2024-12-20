package schedule

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByLabOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Schedule
}

type ListByLabCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Schedule
}

type ListByLabCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByLabCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByLab ...
func (c ScheduleClient) ListByLab(ctx context.Context, id LabId) (result ListByLabOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListByLabCustomPager{},
		Path:       fmt.Sprintf("%s/schedules", id.ID()),
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
		Values *[]Schedule `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByLabComplete retrieves all the results into a single object
func (c ScheduleClient) ListByLabComplete(ctx context.Context, id LabId) (ListByLabCompleteResult, error) {
	return c.ListByLabCompleteMatchingPredicate(ctx, id, ScheduleOperationPredicate{})
}

// ListByLabCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ScheduleClient) ListByLabCompleteMatchingPredicate(ctx context.Context, id LabId, predicate ScheduleOperationPredicate) (result ListByLabCompleteResult, err error) {
	items := make([]Schedule, 0)

	resp, err := c.ListByLab(ctx, id)
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

	result = ListByLabCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
