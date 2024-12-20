package redis

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PatchSchedulesListByRedisResourceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]RedisPatchSchedule
}

type PatchSchedulesListByRedisResourceCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []RedisPatchSchedule
}

type PatchSchedulesListByRedisResourceCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *PatchSchedulesListByRedisResourceCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// PatchSchedulesListByRedisResource ...
func (c RedisClient) PatchSchedulesListByRedisResource(ctx context.Context, id RediId) (result PatchSchedulesListByRedisResourceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &PatchSchedulesListByRedisResourceCustomPager{},
		Path:       fmt.Sprintf("%s/patchSchedules", id.ID()),
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
		Values *[]RedisPatchSchedule `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// PatchSchedulesListByRedisResourceComplete retrieves all the results into a single object
func (c RedisClient) PatchSchedulesListByRedisResourceComplete(ctx context.Context, id RediId) (PatchSchedulesListByRedisResourceCompleteResult, error) {
	return c.PatchSchedulesListByRedisResourceCompleteMatchingPredicate(ctx, id, RedisPatchScheduleOperationPredicate{})
}

// PatchSchedulesListByRedisResourceCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RedisClient) PatchSchedulesListByRedisResourceCompleteMatchingPredicate(ctx context.Context, id RediId, predicate RedisPatchScheduleOperationPredicate) (result PatchSchedulesListByRedisResourceCompleteResult, err error) {
	items := make([]RedisPatchSchedule, 0)

	resp, err := c.PatchSchedulesListByRedisResource(ctx, id)
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

	result = PatchSchedulesListByRedisResourceCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
