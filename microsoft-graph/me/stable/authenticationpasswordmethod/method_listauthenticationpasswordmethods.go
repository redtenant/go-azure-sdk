package authenticationpasswordmethod

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

type ListAuthenticationPasswordMethodsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.PasswordAuthenticationMethod
}

type ListAuthenticationPasswordMethodsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.PasswordAuthenticationMethod
}

type ListAuthenticationPasswordMethodsOperationOptions struct {
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

func DefaultListAuthenticationPasswordMethodsOperationOptions() ListAuthenticationPasswordMethodsOperationOptions {
	return ListAuthenticationPasswordMethodsOperationOptions{}
}

func (o ListAuthenticationPasswordMethodsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListAuthenticationPasswordMethodsOperationOptions) ToOData() *odata.Query {
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

func (o ListAuthenticationPasswordMethodsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListAuthenticationPasswordMethodsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAuthenticationPasswordMethodsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAuthenticationPasswordMethods - List passwordMethods. Retrieve a list of the passwords registered to a user,
// represented by a passwordAuthenticationMethod object. This API returns exactly one object referenced by ID
// 28c10230-6103-485e-b985-444c60001490, as a user can have exactly one password. For security, the password itself is
// never returned in the object and the password property is always null.
func (c AuthenticationPasswordMethodClient) ListAuthenticationPasswordMethods(ctx context.Context, options ListAuthenticationPasswordMethodsOperationOptions) (result ListAuthenticationPasswordMethodsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListAuthenticationPasswordMethodsCustomPager{},
		Path:          "/me/authentication/passwordMethods",
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
		Values *[]stable.PasswordAuthenticationMethod `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAuthenticationPasswordMethodsComplete retrieves all the results into a single object
func (c AuthenticationPasswordMethodClient) ListAuthenticationPasswordMethodsComplete(ctx context.Context, options ListAuthenticationPasswordMethodsOperationOptions) (ListAuthenticationPasswordMethodsCompleteResult, error) {
	return c.ListAuthenticationPasswordMethodsCompleteMatchingPredicate(ctx, options, PasswordAuthenticationMethodOperationPredicate{})
}

// ListAuthenticationPasswordMethodsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AuthenticationPasswordMethodClient) ListAuthenticationPasswordMethodsCompleteMatchingPredicate(ctx context.Context, options ListAuthenticationPasswordMethodsOperationOptions, predicate PasswordAuthenticationMethodOperationPredicate) (result ListAuthenticationPasswordMethodsCompleteResult, err error) {
	items := make([]stable.PasswordAuthenticationMethod, 0)

	resp, err := c.ListAuthenticationPasswordMethods(ctx, options)
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

	result = ListAuthenticationPasswordMethodsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
