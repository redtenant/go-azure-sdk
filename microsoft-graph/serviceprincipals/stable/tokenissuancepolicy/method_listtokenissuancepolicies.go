package tokenissuancepolicy

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

type ListTokenIssuancePoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.TokenIssuancePolicy
}

type ListTokenIssuancePoliciesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.TokenIssuancePolicy
}

type ListTokenIssuancePoliciesOperationOptions struct {
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

func DefaultListTokenIssuancePoliciesOperationOptions() ListTokenIssuancePoliciesOperationOptions {
	return ListTokenIssuancePoliciesOperationOptions{}
}

func (o ListTokenIssuancePoliciesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListTokenIssuancePoliciesOperationOptions) ToOData() *odata.Query {
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

func (o ListTokenIssuancePoliciesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListTokenIssuancePoliciesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTokenIssuancePoliciesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTokenIssuancePolicies - Get tokenIssuancePolicies from servicePrincipals. The tokenIssuancePolicies assigned to
// this service principal.
func (c TokenIssuancePolicyClient) ListTokenIssuancePolicies(ctx context.Context, id stable.ServicePrincipalId, options ListTokenIssuancePoliciesOperationOptions) (result ListTokenIssuancePoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListTokenIssuancePoliciesCustomPager{},
		Path:          fmt.Sprintf("%s/tokenIssuancePolicies", id.ID()),
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
		Values *[]stable.TokenIssuancePolicy `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTokenIssuancePoliciesComplete retrieves all the results into a single object
func (c TokenIssuancePolicyClient) ListTokenIssuancePoliciesComplete(ctx context.Context, id stable.ServicePrincipalId, options ListTokenIssuancePoliciesOperationOptions) (ListTokenIssuancePoliciesCompleteResult, error) {
	return c.ListTokenIssuancePoliciesCompleteMatchingPredicate(ctx, id, options, TokenIssuancePolicyOperationPredicate{})
}

// ListTokenIssuancePoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TokenIssuancePolicyClient) ListTokenIssuancePoliciesCompleteMatchingPredicate(ctx context.Context, id stable.ServicePrincipalId, options ListTokenIssuancePoliciesOperationOptions, predicate TokenIssuancePolicyOperationPredicate) (result ListTokenIssuancePoliciesCompleteResult, err error) {
	items := make([]stable.TokenIssuancePolicy, 0)

	resp, err := c.ListTokenIssuancePolicies(ctx, id, options)
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

	result = ListTokenIssuancePoliciesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
