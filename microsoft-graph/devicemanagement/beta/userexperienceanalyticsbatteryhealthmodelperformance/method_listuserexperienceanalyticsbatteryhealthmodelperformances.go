package userexperienceanalyticsbatteryhealthmodelperformance

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

type ListUserExperienceAnalyticsBatteryHealthModelPerformancesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UserExperienceAnalyticsBatteryHealthModelPerformance
}

type ListUserExperienceAnalyticsBatteryHealthModelPerformancesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UserExperienceAnalyticsBatteryHealthModelPerformance
}

type ListUserExperienceAnalyticsBatteryHealthModelPerformancesOperationOptions struct {
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

func DefaultListUserExperienceAnalyticsBatteryHealthModelPerformancesOperationOptions() ListUserExperienceAnalyticsBatteryHealthModelPerformancesOperationOptions {
	return ListUserExperienceAnalyticsBatteryHealthModelPerformancesOperationOptions{}
}

func (o ListUserExperienceAnalyticsBatteryHealthModelPerformancesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListUserExperienceAnalyticsBatteryHealthModelPerformancesOperationOptions) ToOData() *odata.Query {
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

func (o ListUserExperienceAnalyticsBatteryHealthModelPerformancesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListUserExperienceAnalyticsBatteryHealthModelPerformancesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserExperienceAnalyticsBatteryHealthModelPerformancesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserExperienceAnalyticsBatteryHealthModelPerformances - Get userExperienceAnalyticsBatteryHealthModelPerformance
// from deviceManagement. User Experience Analytics Battery Health Model Performance
func (c UserExperienceAnalyticsBatteryHealthModelPerformanceClient) ListUserExperienceAnalyticsBatteryHealthModelPerformances(ctx context.Context, options ListUserExperienceAnalyticsBatteryHealthModelPerformancesOperationOptions) (result ListUserExperienceAnalyticsBatteryHealthModelPerformancesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListUserExperienceAnalyticsBatteryHealthModelPerformancesCustomPager{},
		Path:          "/deviceManagement/userExperienceAnalyticsBatteryHealthModelPerformance",
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
		Values *[]beta.UserExperienceAnalyticsBatteryHealthModelPerformance `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserExperienceAnalyticsBatteryHealthModelPerformancesComplete retrieves all the results into a single object
func (c UserExperienceAnalyticsBatteryHealthModelPerformanceClient) ListUserExperienceAnalyticsBatteryHealthModelPerformancesComplete(ctx context.Context, options ListUserExperienceAnalyticsBatteryHealthModelPerformancesOperationOptions) (ListUserExperienceAnalyticsBatteryHealthModelPerformancesCompleteResult, error) {
	return c.ListUserExperienceAnalyticsBatteryHealthModelPerformancesCompleteMatchingPredicate(ctx, options, UserExperienceAnalyticsBatteryHealthModelPerformanceOperationPredicate{})
}

// ListUserExperienceAnalyticsBatteryHealthModelPerformancesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserExperienceAnalyticsBatteryHealthModelPerformanceClient) ListUserExperienceAnalyticsBatteryHealthModelPerformancesCompleteMatchingPredicate(ctx context.Context, options ListUserExperienceAnalyticsBatteryHealthModelPerformancesOperationOptions, predicate UserExperienceAnalyticsBatteryHealthModelPerformanceOperationPredicate) (result ListUserExperienceAnalyticsBatteryHealthModelPerformancesCompleteResult, err error) {
	items := make([]beta.UserExperienceAnalyticsBatteryHealthModelPerformance, 0)

	resp, err := c.ListUserExperienceAnalyticsBatteryHealthModelPerformances(ctx, options)
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

	result = ListUserExperienceAnalyticsBatteryHealthModelPerformancesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
