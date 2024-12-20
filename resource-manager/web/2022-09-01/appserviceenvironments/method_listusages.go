package appserviceenvironments

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

type ListUsagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]CsmUsageQuota
}

type ListUsagesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []CsmUsageQuota
}

type ListUsagesOperationOptions struct {
	Filter *string
}

func DefaultListUsagesOperationOptions() ListUsagesOperationOptions {
	return ListUsagesOperationOptions{}
}

func (o ListUsagesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListUsagesOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ListUsagesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	return &out
}

type ListUsagesCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListUsagesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUsages ...
func (c AppServiceEnvironmentsClient) ListUsages(ctx context.Context, id commonids.AppServiceEnvironmentId, options ListUsagesOperationOptions) (result ListUsagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListUsagesCustomPager{},
		Path:          fmt.Sprintf("%s/usages", id.ID()),
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
		Values *[]CsmUsageQuota `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUsagesComplete retrieves all the results into a single object
func (c AppServiceEnvironmentsClient) ListUsagesComplete(ctx context.Context, id commonids.AppServiceEnvironmentId, options ListUsagesOperationOptions) (ListUsagesCompleteResult, error) {
	return c.ListUsagesCompleteMatchingPredicate(ctx, id, options, CsmUsageQuotaOperationPredicate{})
}

// ListUsagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AppServiceEnvironmentsClient) ListUsagesCompleteMatchingPredicate(ctx context.Context, id commonids.AppServiceEnvironmentId, options ListUsagesOperationOptions, predicate CsmUsageQuotaOperationPredicate) (result ListUsagesCompleteResult, err error) {
	items := make([]CsmUsageQuota, 0)

	resp, err := c.ListUsages(ctx, id, options)
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

	result = ListUsagesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
