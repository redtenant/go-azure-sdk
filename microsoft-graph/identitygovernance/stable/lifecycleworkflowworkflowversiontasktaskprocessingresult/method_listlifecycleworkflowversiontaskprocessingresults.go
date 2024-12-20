package lifecycleworkflowworkflowversiontasktaskprocessingresult

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

type ListLifecycleWorkflowVersionTaskProcessingResultsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.IdentityGovernanceTaskProcessingResult
}

type ListLifecycleWorkflowVersionTaskProcessingResultsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.IdentityGovernanceTaskProcessingResult
}

type ListLifecycleWorkflowVersionTaskProcessingResultsOperationOptions struct {
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

func DefaultListLifecycleWorkflowVersionTaskProcessingResultsOperationOptions() ListLifecycleWorkflowVersionTaskProcessingResultsOperationOptions {
	return ListLifecycleWorkflowVersionTaskProcessingResultsOperationOptions{}
}

func (o ListLifecycleWorkflowVersionTaskProcessingResultsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListLifecycleWorkflowVersionTaskProcessingResultsOperationOptions) ToOData() *odata.Query {
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

func (o ListLifecycleWorkflowVersionTaskProcessingResultsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListLifecycleWorkflowVersionTaskProcessingResultsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListLifecycleWorkflowVersionTaskProcessingResultsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListLifecycleWorkflowVersionTaskProcessingResults - Get taskProcessingResults from identityGovernance. The result of
// processing the task.
func (c LifecycleWorkflowWorkflowVersionTaskTaskProcessingResultClient) ListLifecycleWorkflowVersionTaskProcessingResults(ctx context.Context, id stable.IdentityGovernanceLifecycleWorkflowWorkflowIdVersionIdTaskId, options ListLifecycleWorkflowVersionTaskProcessingResultsOperationOptions) (result ListLifecycleWorkflowVersionTaskProcessingResultsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListLifecycleWorkflowVersionTaskProcessingResultsCustomPager{},
		Path:          fmt.Sprintf("%s/taskProcessingResults", id.ID()),
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
		Values *[]stable.IdentityGovernanceTaskProcessingResult `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListLifecycleWorkflowVersionTaskProcessingResultsComplete retrieves all the results into a single object
func (c LifecycleWorkflowWorkflowVersionTaskTaskProcessingResultClient) ListLifecycleWorkflowVersionTaskProcessingResultsComplete(ctx context.Context, id stable.IdentityGovernanceLifecycleWorkflowWorkflowIdVersionIdTaskId, options ListLifecycleWorkflowVersionTaskProcessingResultsOperationOptions) (ListLifecycleWorkflowVersionTaskProcessingResultsCompleteResult, error) {
	return c.ListLifecycleWorkflowVersionTaskProcessingResultsCompleteMatchingPredicate(ctx, id, options, IdentityGovernanceTaskProcessingResultOperationPredicate{})
}

// ListLifecycleWorkflowVersionTaskProcessingResultsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c LifecycleWorkflowWorkflowVersionTaskTaskProcessingResultClient) ListLifecycleWorkflowVersionTaskProcessingResultsCompleteMatchingPredicate(ctx context.Context, id stable.IdentityGovernanceLifecycleWorkflowWorkflowIdVersionIdTaskId, options ListLifecycleWorkflowVersionTaskProcessingResultsOperationOptions, predicate IdentityGovernanceTaskProcessingResultOperationPredicate) (result ListLifecycleWorkflowVersionTaskProcessingResultsCompleteResult, err error) {
	items := make([]stable.IdentityGovernanceTaskProcessingResult, 0)

	resp, err := c.ListLifecycleWorkflowVersionTaskProcessingResults(ctx, id, options)
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

	result = ListLifecycleWorkflowVersionTaskProcessingResultsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
