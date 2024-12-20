package pendingaccessreviewinstancestagedecisioninstancedecision

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

type ListPendingAccessReviewInstanceStageDecisionInstanceDecisionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AccessReviewInstanceDecisionItem
}

type ListPendingAccessReviewInstanceStageDecisionInstanceDecisionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AccessReviewInstanceDecisionItem
}

type ListPendingAccessReviewInstanceStageDecisionInstanceDecisionsOperationOptions struct {
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

func DefaultListPendingAccessReviewInstanceStageDecisionInstanceDecisionsOperationOptions() ListPendingAccessReviewInstanceStageDecisionInstanceDecisionsOperationOptions {
	return ListPendingAccessReviewInstanceStageDecisionInstanceDecisionsOperationOptions{}
}

func (o ListPendingAccessReviewInstanceStageDecisionInstanceDecisionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListPendingAccessReviewInstanceStageDecisionInstanceDecisionsOperationOptions) ToOData() *odata.Query {
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

func (o ListPendingAccessReviewInstanceStageDecisionInstanceDecisionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListPendingAccessReviewInstanceStageDecisionInstanceDecisionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListPendingAccessReviewInstanceStageDecisionInstanceDecisionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListPendingAccessReviewInstanceStageDecisionInstanceDecisions - Get decisions from me. Each user reviewed in an
// accessReviewInstance has a decision item representing if they were approved, denied, or not yet reviewed.
func (c PendingAccessReviewInstanceStageDecisionInstanceDecisionClient) ListPendingAccessReviewInstanceStageDecisionInstanceDecisions(ctx context.Context, id beta.MePendingAccessReviewInstanceIdStageIdDecisionId, options ListPendingAccessReviewInstanceStageDecisionInstanceDecisionsOperationOptions) (result ListPendingAccessReviewInstanceStageDecisionInstanceDecisionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListPendingAccessReviewInstanceStageDecisionInstanceDecisionsCustomPager{},
		Path:          fmt.Sprintf("%s/instance/decisions", id.ID()),
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
		Values *[]beta.AccessReviewInstanceDecisionItem `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPendingAccessReviewInstanceStageDecisionInstanceDecisionsComplete retrieves all the results into a single object
func (c PendingAccessReviewInstanceStageDecisionInstanceDecisionClient) ListPendingAccessReviewInstanceStageDecisionInstanceDecisionsComplete(ctx context.Context, id beta.MePendingAccessReviewInstanceIdStageIdDecisionId, options ListPendingAccessReviewInstanceStageDecisionInstanceDecisionsOperationOptions) (ListPendingAccessReviewInstanceStageDecisionInstanceDecisionsCompleteResult, error) {
	return c.ListPendingAccessReviewInstanceStageDecisionInstanceDecisionsCompleteMatchingPredicate(ctx, id, options, AccessReviewInstanceDecisionItemOperationPredicate{})
}

// ListPendingAccessReviewInstanceStageDecisionInstanceDecisionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PendingAccessReviewInstanceStageDecisionInstanceDecisionClient) ListPendingAccessReviewInstanceStageDecisionInstanceDecisionsCompleteMatchingPredicate(ctx context.Context, id beta.MePendingAccessReviewInstanceIdStageIdDecisionId, options ListPendingAccessReviewInstanceStageDecisionInstanceDecisionsOperationOptions, predicate AccessReviewInstanceDecisionItemOperationPredicate) (result ListPendingAccessReviewInstanceStageDecisionInstanceDecisionsCompleteResult, err error) {
	items := make([]beta.AccessReviewInstanceDecisionItem, 0)

	resp, err := c.ListPendingAccessReviewInstanceStageDecisionInstanceDecisions(ctx, id, options)
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

	result = ListPendingAccessReviewInstanceStageDecisionInstanceDecisionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
