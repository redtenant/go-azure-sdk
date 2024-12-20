package lifecycleworkflowworkflowexecutionscope

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

type ListLifecycleWorkflowExecutionScopesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.IdentityGovernanceUserProcessingResult
}

type ListLifecycleWorkflowExecutionScopesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.IdentityGovernanceUserProcessingResult
}

type ListLifecycleWorkflowExecutionScopesOperationOptions struct {
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

func DefaultListLifecycleWorkflowExecutionScopesOperationOptions() ListLifecycleWorkflowExecutionScopesOperationOptions {
	return ListLifecycleWorkflowExecutionScopesOperationOptions{}
}

func (o ListLifecycleWorkflowExecutionScopesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListLifecycleWorkflowExecutionScopesOperationOptions) ToOData() *odata.Query {
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

func (o ListLifecycleWorkflowExecutionScopesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListLifecycleWorkflowExecutionScopesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListLifecycleWorkflowExecutionScopesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListLifecycleWorkflowExecutionScopes - Get executionScope from identityGovernance. The unique identifier of the
// Microsoft Entra identity that last modified the workflow object.
func (c LifecycleWorkflowWorkflowExecutionScopeClient) ListLifecycleWorkflowExecutionScopes(ctx context.Context, id beta.IdentityGovernanceLifecycleWorkflowWorkflowId, options ListLifecycleWorkflowExecutionScopesOperationOptions) (result ListLifecycleWorkflowExecutionScopesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListLifecycleWorkflowExecutionScopesCustomPager{},
		Path:          fmt.Sprintf("%s/executionScope", id.ID()),
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

// ListLifecycleWorkflowExecutionScopesComplete retrieves all the results into a single object
func (c LifecycleWorkflowWorkflowExecutionScopeClient) ListLifecycleWorkflowExecutionScopesComplete(ctx context.Context, id beta.IdentityGovernanceLifecycleWorkflowWorkflowId, options ListLifecycleWorkflowExecutionScopesOperationOptions) (ListLifecycleWorkflowExecutionScopesCompleteResult, error) {
	return c.ListLifecycleWorkflowExecutionScopesCompleteMatchingPredicate(ctx, id, options, IdentityGovernanceUserProcessingResultOperationPredicate{})
}

// ListLifecycleWorkflowExecutionScopesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c LifecycleWorkflowWorkflowExecutionScopeClient) ListLifecycleWorkflowExecutionScopesCompleteMatchingPredicate(ctx context.Context, id beta.IdentityGovernanceLifecycleWorkflowWorkflowId, options ListLifecycleWorkflowExecutionScopesOperationOptions, predicate IdentityGovernanceUserProcessingResultOperationPredicate) (result ListLifecycleWorkflowExecutionScopesCompleteResult, err error) {
	items := make([]beta.IdentityGovernanceUserProcessingResult, 0)

	resp, err := c.ListLifecycleWorkflowExecutionScopes(ctx, id, options)
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

	result = ListLifecycleWorkflowExecutionScopesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
