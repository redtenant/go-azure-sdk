package userinsightmonthlyinactiveuser

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListUserInsightMonthlyInactiveUsersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.MonthlyInactiveUsersMetric
}

type ListUserInsightMonthlyInactiveUsersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.MonthlyInactiveUsersMetric
}

type ListUserInsightMonthlyInactiveUsersOperationOptions struct {
	Count     *bool
	Expand    *odata.Expand
	Filter    *string
	Metadata  *odata.Metadata
	OrderBy   *odata.OrderBy
	RetryFunc client.RequestRetryFunc
	Search    *string
	Select    *[]string
	Skip      *int64
	Top       *int64
}

func DefaultListUserInsightMonthlyInactiveUsersOperationOptions() ListUserInsightMonthlyInactiveUsersOperationOptions {
	return ListUserInsightMonthlyInactiveUsersOperationOptions{}
}

func (o ListUserInsightMonthlyInactiveUsersOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListUserInsightMonthlyInactiveUsersOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Count != nil {
		out.Count = *o.Count
	}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.OrderBy != nil {
		out.OrderBy = *o.OrderBy
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListUserInsightMonthlyInactiveUsersOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListUserInsightMonthlyInactiveUsersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserInsightMonthlyInactiveUsersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserInsightMonthlyInactiveUsers - Get inactiveUsers from reports
func (c UserInsightMonthlyInactiveUserClient) ListUserInsightMonthlyInactiveUsers(ctx context.Context, options ListUserInsightMonthlyInactiveUsersOperationOptions) (result ListUserInsightMonthlyInactiveUsersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListUserInsightMonthlyInactiveUsersCustomPager{},
		Path:          "/reports/userInsights/monthly/inactiveUsers",
		RetryFunc:     options.RetryFunc,
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
		Values *[]beta.MonthlyInactiveUsersMetric `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserInsightMonthlyInactiveUsersComplete retrieves all the results into a single object
func (c UserInsightMonthlyInactiveUserClient) ListUserInsightMonthlyInactiveUsersComplete(ctx context.Context, options ListUserInsightMonthlyInactiveUsersOperationOptions) (ListUserInsightMonthlyInactiveUsersCompleteResult, error) {
	return c.ListUserInsightMonthlyInactiveUsersCompleteMatchingPredicate(ctx, options, MonthlyInactiveUsersMetricOperationPredicate{})
}

// ListUserInsightMonthlyInactiveUsersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserInsightMonthlyInactiveUserClient) ListUserInsightMonthlyInactiveUsersCompleteMatchingPredicate(ctx context.Context, options ListUserInsightMonthlyInactiveUsersOperationOptions, predicate MonthlyInactiveUsersMetricOperationPredicate) (result ListUserInsightMonthlyInactiveUsersCompleteResult, err error) {
	items := make([]beta.MonthlyInactiveUsersMetric, 0)

	resp, err := c.ListUserInsightMonthlyInactiveUsers(ctx, options)
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

	result = ListUserInsightMonthlyInactiveUsersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
