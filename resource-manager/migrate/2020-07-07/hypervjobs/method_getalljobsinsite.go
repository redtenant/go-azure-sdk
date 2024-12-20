package hypervjobs

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetAllJobsInSiteOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]HyperVJob
}

type GetAllJobsInSiteCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []HyperVJob
}

type GetAllJobsInSiteCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *GetAllJobsInSiteCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// GetAllJobsInSite ...
func (c HyperVJobsClient) GetAllJobsInSite(ctx context.Context, id HyperVSiteId) (result GetAllJobsInSiteOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &GetAllJobsInSiteCustomPager{},
		Path:       fmt.Sprintf("%s/jobs", id.ID()),
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
		Values *[]HyperVJob `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// GetAllJobsInSiteComplete retrieves all the results into a single object
func (c HyperVJobsClient) GetAllJobsInSiteComplete(ctx context.Context, id HyperVSiteId) (GetAllJobsInSiteCompleteResult, error) {
	return c.GetAllJobsInSiteCompleteMatchingPredicate(ctx, id, HyperVJobOperationPredicate{})
}

// GetAllJobsInSiteCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c HyperVJobsClient) GetAllJobsInSiteCompleteMatchingPredicate(ctx context.Context, id HyperVSiteId, predicate HyperVJobOperationPredicate) (result GetAllJobsInSiteCompleteResult, err error) {
	items := make([]HyperVJob, 0)

	resp, err := c.GetAllJobsInSite(ctx, id)
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

	result = GetAllJobsInSiteCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
