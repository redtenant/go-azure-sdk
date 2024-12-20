package appconsentrequestsforapprovaluserconsentrequest

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

type ListAppConsentRequestsForApprovalUserConsentRequestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UserConsentRequest
}

type ListAppConsentRequestsForApprovalUserConsentRequestsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UserConsentRequest
}

type ListAppConsentRequestsForApprovalUserConsentRequestsOperationOptions struct {
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

func DefaultListAppConsentRequestsForApprovalUserConsentRequestsOperationOptions() ListAppConsentRequestsForApprovalUserConsentRequestsOperationOptions {
	return ListAppConsentRequestsForApprovalUserConsentRequestsOperationOptions{}
}

func (o ListAppConsentRequestsForApprovalUserConsentRequestsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListAppConsentRequestsForApprovalUserConsentRequestsOperationOptions) ToOData() *odata.Query {
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

func (o ListAppConsentRequestsForApprovalUserConsentRequestsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListAppConsentRequestsForApprovalUserConsentRequestsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAppConsentRequestsForApprovalUserConsentRequestsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAppConsentRequestsForApprovalUserConsentRequests - Get userConsentRequests from me. A list of pending user
// consent requests. Supports $filter (eq).
func (c AppConsentRequestsForApprovalUserConsentRequestClient) ListAppConsentRequestsForApprovalUserConsentRequests(ctx context.Context, id beta.MeAppConsentRequestsForApprovalId, options ListAppConsentRequestsForApprovalUserConsentRequestsOperationOptions) (result ListAppConsentRequestsForApprovalUserConsentRequestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListAppConsentRequestsForApprovalUserConsentRequestsCustomPager{},
		Path:          fmt.Sprintf("%s/userConsentRequests", id.ID()),
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
		Values *[]beta.UserConsentRequest `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAppConsentRequestsForApprovalUserConsentRequestsComplete retrieves all the results into a single object
func (c AppConsentRequestsForApprovalUserConsentRequestClient) ListAppConsentRequestsForApprovalUserConsentRequestsComplete(ctx context.Context, id beta.MeAppConsentRequestsForApprovalId, options ListAppConsentRequestsForApprovalUserConsentRequestsOperationOptions) (ListAppConsentRequestsForApprovalUserConsentRequestsCompleteResult, error) {
	return c.ListAppConsentRequestsForApprovalUserConsentRequestsCompleteMatchingPredicate(ctx, id, options, UserConsentRequestOperationPredicate{})
}

// ListAppConsentRequestsForApprovalUserConsentRequestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AppConsentRequestsForApprovalUserConsentRequestClient) ListAppConsentRequestsForApprovalUserConsentRequestsCompleteMatchingPredicate(ctx context.Context, id beta.MeAppConsentRequestsForApprovalId, options ListAppConsentRequestsForApprovalUserConsentRequestsOperationOptions, predicate UserConsentRequestOperationPredicate) (result ListAppConsentRequestsForApprovalUserConsentRequestsCompleteResult, err error) {
	items := make([]beta.UserConsentRequest, 0)

	resp, err := c.ListAppConsentRequestsForApprovalUserConsentRequests(ctx, id, options)
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

	result = ListAppConsentRequestsForApprovalUserConsentRequestsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
