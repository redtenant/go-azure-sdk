package jobversions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByJobOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Resource
}

type ListByJobCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Resource
}

type ListByJobCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByJobCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByJob ...
func (c JobVersionsClient) ListByJob(ctx context.Context, id JobId) (result ListByJobOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListByJobCustomPager{},
		Path:       fmt.Sprintf("%s/versions", id.ID()),
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
		Values *[]Resource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByJobComplete retrieves all the results into a single object
func (c JobVersionsClient) ListByJobComplete(ctx context.Context, id JobId) (ListByJobCompleteResult, error) {
	return c.ListByJobCompleteMatchingPredicate(ctx, id, ResourceOperationPredicate{})
}

// ListByJobCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c JobVersionsClient) ListByJobCompleteMatchingPredicate(ctx context.Context, id JobId, predicate ResourceOperationPredicate) (result ListByJobCompleteResult, err error) {
	items := make([]Resource, 0)

	resp, err := c.ListByJob(ctx, id)
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

	result = ListByJobCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
