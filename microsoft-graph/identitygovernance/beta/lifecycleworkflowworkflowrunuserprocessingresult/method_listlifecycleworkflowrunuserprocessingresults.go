package lifecycleworkflowworkflowrunuserprocessingresult

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

type ListLifecycleWorkflowRunUserProcessingResultsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.IdentityGovernanceUserProcessingResult
}

type ListLifecycleWorkflowRunUserProcessingResultsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.IdentityGovernanceUserProcessingResult
}

type ListLifecycleWorkflowRunUserProcessingResultsOperationOptions struct {
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

func DefaultListLifecycleWorkflowRunUserProcessingResultsOperationOptions() ListLifecycleWorkflowRunUserProcessingResultsOperationOptions {
	return ListLifecycleWorkflowRunUserProcessingResultsOperationOptions{}
}

func (o ListLifecycleWorkflowRunUserProcessingResultsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListLifecycleWorkflowRunUserProcessingResultsOperationOptions) ToOData() *odata.Query {
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

func (o ListLifecycleWorkflowRunUserProcessingResultsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListLifecycleWorkflowRunUserProcessingResultsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListLifecycleWorkflowRunUserProcessingResultsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListLifecycleWorkflowRunUserProcessingResults - List userProcessingResults. Get user processing results of a workflow
// run object.
func (c LifecycleWorkflowWorkflowRunUserProcessingResultClient) ListLifecycleWorkflowRunUserProcessingResults(ctx context.Context, id beta.IdentityGovernanceLifecycleWorkflowWorkflowIdRunId, options ListLifecycleWorkflowRunUserProcessingResultsOperationOptions) (result ListLifecycleWorkflowRunUserProcessingResultsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListLifecycleWorkflowRunUserProcessingResultsCustomPager{},
		Path:          fmt.Sprintf("%s/userProcessingResults", id.ID()),
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
		Values *[]beta.IdentityGovernanceUserProcessingResult `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListLifecycleWorkflowRunUserProcessingResultsComplete retrieves all the results into a single object
func (c LifecycleWorkflowWorkflowRunUserProcessingResultClient) ListLifecycleWorkflowRunUserProcessingResultsComplete(ctx context.Context, id beta.IdentityGovernanceLifecycleWorkflowWorkflowIdRunId, options ListLifecycleWorkflowRunUserProcessingResultsOperationOptions) (ListLifecycleWorkflowRunUserProcessingResultsCompleteResult, error) {
	return c.ListLifecycleWorkflowRunUserProcessingResultsCompleteMatchingPredicate(ctx, id, options, IdentityGovernanceUserProcessingResultOperationPredicate{})
}

// ListLifecycleWorkflowRunUserProcessingResultsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c LifecycleWorkflowWorkflowRunUserProcessingResultClient) ListLifecycleWorkflowRunUserProcessingResultsCompleteMatchingPredicate(ctx context.Context, id beta.IdentityGovernanceLifecycleWorkflowWorkflowIdRunId, options ListLifecycleWorkflowRunUserProcessingResultsOperationOptions, predicate IdentityGovernanceUserProcessingResultOperationPredicate) (result ListLifecycleWorkflowRunUserProcessingResultsCompleteResult, err error) {
	items := make([]beta.IdentityGovernanceUserProcessingResult, 0)

	resp, err := c.ListLifecycleWorkflowRunUserProcessingResults(ctx, id, options)
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

	result = ListLifecycleWorkflowRunUserProcessingResultsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
