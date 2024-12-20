package plannertask

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

type ListPlannerTasksOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.PlannerTask
}

type ListPlannerTasksCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.PlannerTask
}

type ListPlannerTasksOperationOptions struct {
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

func DefaultListPlannerTasksOperationOptions() ListPlannerTasksOperationOptions {
	return ListPlannerTasksOperationOptions{}
}

func (o ListPlannerTasksOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListPlannerTasksOperationOptions) ToOData() *odata.Query {
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

func (o ListPlannerTasksOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListPlannerTasksCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListPlannerTasksCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListPlannerTasks - Get tasks from users. Read-only. Nullable. Returns the plannerPlans shared with the user.
func (c PlannerTaskClient) ListPlannerTasks(ctx context.Context, id stable.UserId, options ListPlannerTasksOperationOptions) (result ListPlannerTasksOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListPlannerTasksCustomPager{},
		Path:          fmt.Sprintf("%s/planner/tasks", id.ID()),
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
		Values *[]stable.PlannerTask `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPlannerTasksComplete retrieves all the results into a single object
func (c PlannerTaskClient) ListPlannerTasksComplete(ctx context.Context, id stable.UserId, options ListPlannerTasksOperationOptions) (ListPlannerTasksCompleteResult, error) {
	return c.ListPlannerTasksCompleteMatchingPredicate(ctx, id, options, PlannerTaskOperationPredicate{})
}

// ListPlannerTasksCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PlannerTaskClient) ListPlannerTasksCompleteMatchingPredicate(ctx context.Context, id stable.UserId, options ListPlannerTasksOperationOptions, predicate PlannerTaskOperationPredicate) (result ListPlannerTasksCompleteResult, err error) {
	items := make([]stable.PlannerTask, 0)

	resp, err := c.ListPlannerTasks(ctx, id, options)
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

	result = ListPlannerTasksCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
