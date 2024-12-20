package userexperienceanalyticsanomalycorrelationgroupoverview

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

type ListUserExperienceAnalyticsAnomalyCorrelationGroupOverviewsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UserExperienceAnalyticsAnomalyCorrelationGroupOverview
}

type ListUserExperienceAnalyticsAnomalyCorrelationGroupOverviewsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UserExperienceAnalyticsAnomalyCorrelationGroupOverview
}

type ListUserExperienceAnalyticsAnomalyCorrelationGroupOverviewsOperationOptions struct {
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

func DefaultListUserExperienceAnalyticsAnomalyCorrelationGroupOverviewsOperationOptions() ListUserExperienceAnalyticsAnomalyCorrelationGroupOverviewsOperationOptions {
	return ListUserExperienceAnalyticsAnomalyCorrelationGroupOverviewsOperationOptions{}
}

func (o ListUserExperienceAnalyticsAnomalyCorrelationGroupOverviewsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListUserExperienceAnalyticsAnomalyCorrelationGroupOverviewsOperationOptions) ToOData() *odata.Query {
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

func (o ListUserExperienceAnalyticsAnomalyCorrelationGroupOverviewsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListUserExperienceAnalyticsAnomalyCorrelationGroupOverviewsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserExperienceAnalyticsAnomalyCorrelationGroupOverviewsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserExperienceAnalyticsAnomalyCorrelationGroupOverviews - Get
// userExperienceAnalyticsAnomalyCorrelationGroupOverview from deviceManagement. The user experience analytics anomaly
// correlation group overview entity contains the information for each correlation group of an anomaly.
func (c UserExperienceAnalyticsAnomalyCorrelationGroupOverviewClient) ListUserExperienceAnalyticsAnomalyCorrelationGroupOverviews(ctx context.Context, options ListUserExperienceAnalyticsAnomalyCorrelationGroupOverviewsOperationOptions) (result ListUserExperienceAnalyticsAnomalyCorrelationGroupOverviewsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListUserExperienceAnalyticsAnomalyCorrelationGroupOverviewsCustomPager{},
		Path:          "/deviceManagement/userExperienceAnalyticsAnomalyCorrelationGroupOverview",
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
		Values *[]beta.UserExperienceAnalyticsAnomalyCorrelationGroupOverview `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserExperienceAnalyticsAnomalyCorrelationGroupOverviewsComplete retrieves all the results into a single object
func (c UserExperienceAnalyticsAnomalyCorrelationGroupOverviewClient) ListUserExperienceAnalyticsAnomalyCorrelationGroupOverviewsComplete(ctx context.Context, options ListUserExperienceAnalyticsAnomalyCorrelationGroupOverviewsOperationOptions) (ListUserExperienceAnalyticsAnomalyCorrelationGroupOverviewsCompleteResult, error) {
	return c.ListUserExperienceAnalyticsAnomalyCorrelationGroupOverviewsCompleteMatchingPredicate(ctx, options, UserExperienceAnalyticsAnomalyCorrelationGroupOverviewOperationPredicate{})
}

// ListUserExperienceAnalyticsAnomalyCorrelationGroupOverviewsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserExperienceAnalyticsAnomalyCorrelationGroupOverviewClient) ListUserExperienceAnalyticsAnomalyCorrelationGroupOverviewsCompleteMatchingPredicate(ctx context.Context, options ListUserExperienceAnalyticsAnomalyCorrelationGroupOverviewsOperationOptions, predicate UserExperienceAnalyticsAnomalyCorrelationGroupOverviewOperationPredicate) (result ListUserExperienceAnalyticsAnomalyCorrelationGroupOverviewsCompleteResult, err error) {
	items := make([]beta.UserExperienceAnalyticsAnomalyCorrelationGroupOverview, 0)

	resp, err := c.ListUserExperienceAnalyticsAnomalyCorrelationGroupOverviews(ctx, options)
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

	result = ListUserExperienceAnalyticsAnomalyCorrelationGroupOverviewsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
