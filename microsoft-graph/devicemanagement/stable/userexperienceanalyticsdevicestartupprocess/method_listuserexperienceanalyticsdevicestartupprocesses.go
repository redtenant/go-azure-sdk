package userexperienceanalyticsdevicestartupprocess

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListUserExperienceAnalyticsDeviceStartupProcessesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.UserExperienceAnalyticsDeviceStartupProcess
}

type ListUserExperienceAnalyticsDeviceStartupProcessesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.UserExperienceAnalyticsDeviceStartupProcess
}

type ListUserExperienceAnalyticsDeviceStartupProcessesOperationOptions struct {
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

func DefaultListUserExperienceAnalyticsDeviceStartupProcessesOperationOptions() ListUserExperienceAnalyticsDeviceStartupProcessesOperationOptions {
	return ListUserExperienceAnalyticsDeviceStartupProcessesOperationOptions{}
}

func (o ListUserExperienceAnalyticsDeviceStartupProcessesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListUserExperienceAnalyticsDeviceStartupProcessesOperationOptions) ToOData() *odata.Query {
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

func (o ListUserExperienceAnalyticsDeviceStartupProcessesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListUserExperienceAnalyticsDeviceStartupProcessesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserExperienceAnalyticsDeviceStartupProcessesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserExperienceAnalyticsDeviceStartupProcesses - Get userExperienceAnalyticsDeviceStartupProcesses from
// deviceManagement. User experience analytics device Startup Processes
func (c UserExperienceAnalyticsDeviceStartupProcessClient) ListUserExperienceAnalyticsDeviceStartupProcesses(ctx context.Context, options ListUserExperienceAnalyticsDeviceStartupProcessesOperationOptions) (result ListUserExperienceAnalyticsDeviceStartupProcessesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListUserExperienceAnalyticsDeviceStartupProcessesCustomPager{},
		Path:          "/deviceManagement/userExperienceAnalyticsDeviceStartupProcesses",
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
		Values *[]stable.UserExperienceAnalyticsDeviceStartupProcess `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserExperienceAnalyticsDeviceStartupProcessesComplete retrieves all the results into a single object
func (c UserExperienceAnalyticsDeviceStartupProcessClient) ListUserExperienceAnalyticsDeviceStartupProcessesComplete(ctx context.Context, options ListUserExperienceAnalyticsDeviceStartupProcessesOperationOptions) (ListUserExperienceAnalyticsDeviceStartupProcessesCompleteResult, error) {
	return c.ListUserExperienceAnalyticsDeviceStartupProcessesCompleteMatchingPredicate(ctx, options, UserExperienceAnalyticsDeviceStartupProcessOperationPredicate{})
}

// ListUserExperienceAnalyticsDeviceStartupProcessesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserExperienceAnalyticsDeviceStartupProcessClient) ListUserExperienceAnalyticsDeviceStartupProcessesCompleteMatchingPredicate(ctx context.Context, options ListUserExperienceAnalyticsDeviceStartupProcessesOperationOptions, predicate UserExperienceAnalyticsDeviceStartupProcessOperationPredicate) (result ListUserExperienceAnalyticsDeviceStartupProcessesCompleteResult, err error) {
	items := make([]stable.UserExperienceAnalyticsDeviceStartupProcess, 0)

	resp, err := c.ListUserExperienceAnalyticsDeviceStartupProcesses(ctx, options)
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

	result = ListUserExperienceAnalyticsDeviceStartupProcessesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
