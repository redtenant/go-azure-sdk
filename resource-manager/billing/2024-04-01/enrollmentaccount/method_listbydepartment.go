package enrollmentaccount

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByDepartmentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]EnrollmentAccount
}

type ListByDepartmentCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []EnrollmentAccount
}

type ListByDepartmentOperationOptions struct {
	Count   *bool
	Filter  *string
	OrderBy *string
	Search  *string
	Skip    *int64
	Top     *int64
}

func DefaultListByDepartmentOperationOptions() ListByDepartmentOperationOptions {
	return ListByDepartmentOperationOptions{}
}

func (o ListByDepartmentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListByDepartmentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ListByDepartmentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Count != nil {
		out.Append("count", fmt.Sprintf("%v", *o.Count))
	}
	if o.Filter != nil {
		out.Append("filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.OrderBy != nil {
		out.Append("orderBy", fmt.Sprintf("%v", *o.OrderBy))
	}
	if o.Search != nil {
		out.Append("search", fmt.Sprintf("%v", *o.Search))
	}
	if o.Skip != nil {
		out.Append("skip", fmt.Sprintf("%v", *o.Skip))
	}
	if o.Top != nil {
		out.Append("top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

type ListByDepartmentCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByDepartmentCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByDepartment ...
func (c EnrollmentAccountClient) ListByDepartment(ctx context.Context, id DepartmentId, options ListByDepartmentOperationOptions) (result ListByDepartmentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListByDepartmentCustomPager{},
		Path:          fmt.Sprintf("%s/enrollmentAccounts", id.ID()),
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
		Values *[]EnrollmentAccount `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByDepartmentComplete retrieves all the results into a single object
func (c EnrollmentAccountClient) ListByDepartmentComplete(ctx context.Context, id DepartmentId, options ListByDepartmentOperationOptions) (ListByDepartmentCompleteResult, error) {
	return c.ListByDepartmentCompleteMatchingPredicate(ctx, id, options, EnrollmentAccountOperationPredicate{})
}

// ListByDepartmentCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EnrollmentAccountClient) ListByDepartmentCompleteMatchingPredicate(ctx context.Context, id DepartmentId, options ListByDepartmentOperationOptions, predicate EnrollmentAccountOperationPredicate) (result ListByDepartmentCompleteResult, err error) {
	items := make([]EnrollmentAccount, 0)

	resp, err := c.ListByDepartment(ctx, id, options)
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

	result = ListByDepartmentCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
