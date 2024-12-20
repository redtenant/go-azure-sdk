package userinsightdailymfacompletion

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

type ListUserInsightDailyMfaCompletionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.MfaCompletionMetric
}

type ListUserInsightDailyMfaCompletionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.MfaCompletionMetric
}

type ListUserInsightDailyMfaCompletionsOperationOptions struct {
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

func DefaultListUserInsightDailyMfaCompletionsOperationOptions() ListUserInsightDailyMfaCompletionsOperationOptions {
	return ListUserInsightDailyMfaCompletionsOperationOptions{}
}

func (o ListUserInsightDailyMfaCompletionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListUserInsightDailyMfaCompletionsOperationOptions) ToOData() *odata.Query {
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

func (o ListUserInsightDailyMfaCompletionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListUserInsightDailyMfaCompletionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserInsightDailyMfaCompletionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserInsightDailyMfaCompletions - List daily mfaCompletions. Get a list of daily MFA completions on apps
// registered in your tenant configured for Microsoft Entra External ID for customers.
func (c UserInsightDailyMfaCompletionClient) ListUserInsightDailyMfaCompletions(ctx context.Context, options ListUserInsightDailyMfaCompletionsOperationOptions) (result ListUserInsightDailyMfaCompletionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListUserInsightDailyMfaCompletionsCustomPager{},
		Path:          "/reports/userInsights/daily/mfaCompletions",
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
		Values *[]beta.MfaCompletionMetric `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserInsightDailyMfaCompletionsComplete retrieves all the results into a single object
func (c UserInsightDailyMfaCompletionClient) ListUserInsightDailyMfaCompletionsComplete(ctx context.Context, options ListUserInsightDailyMfaCompletionsOperationOptions) (ListUserInsightDailyMfaCompletionsCompleteResult, error) {
	return c.ListUserInsightDailyMfaCompletionsCompleteMatchingPredicate(ctx, options, MfaCompletionMetricOperationPredicate{})
}

// ListUserInsightDailyMfaCompletionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserInsightDailyMfaCompletionClient) ListUserInsightDailyMfaCompletionsCompleteMatchingPredicate(ctx context.Context, options ListUserInsightDailyMfaCompletionsOperationOptions, predicate MfaCompletionMetricOperationPredicate) (result ListUserInsightDailyMfaCompletionsCompleteResult, err error) {
	items := make([]beta.MfaCompletionMetric, 0)

	resp, err := c.ListUserInsightDailyMfaCompletions(ctx, options)
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

	result = ListUserInsightDailyMfaCompletionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
