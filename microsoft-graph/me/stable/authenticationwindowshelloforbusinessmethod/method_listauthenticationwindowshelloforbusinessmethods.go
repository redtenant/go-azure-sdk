package authenticationwindowshelloforbusinessmethod

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

type ListAuthenticationWindowsHelloForBusinessMethodsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.WindowsHelloForBusinessAuthenticationMethod
}

type ListAuthenticationWindowsHelloForBusinessMethodsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.WindowsHelloForBusinessAuthenticationMethod
}

type ListAuthenticationWindowsHelloForBusinessMethodsOperationOptions struct {
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

func DefaultListAuthenticationWindowsHelloForBusinessMethodsOperationOptions() ListAuthenticationWindowsHelloForBusinessMethodsOperationOptions {
	return ListAuthenticationWindowsHelloForBusinessMethodsOperationOptions{}
}

func (o ListAuthenticationWindowsHelloForBusinessMethodsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListAuthenticationWindowsHelloForBusinessMethodsOperationOptions) ToOData() *odata.Query {
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

func (o ListAuthenticationWindowsHelloForBusinessMethodsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListAuthenticationWindowsHelloForBusinessMethodsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAuthenticationWindowsHelloForBusinessMethodsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAuthenticationWindowsHelloForBusinessMethods - Get windowsHelloForBusinessMethods from me. Represents the Windows
// Hello for Business authentication method registered to a user for authentication.
func (c AuthenticationWindowsHelloForBusinessMethodClient) ListAuthenticationWindowsHelloForBusinessMethods(ctx context.Context, options ListAuthenticationWindowsHelloForBusinessMethodsOperationOptions) (result ListAuthenticationWindowsHelloForBusinessMethodsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListAuthenticationWindowsHelloForBusinessMethodsCustomPager{},
		Path:          "/me/authentication/windowsHelloForBusinessMethods",
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
		Values *[]stable.WindowsHelloForBusinessAuthenticationMethod `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAuthenticationWindowsHelloForBusinessMethodsComplete retrieves all the results into a single object
func (c AuthenticationWindowsHelloForBusinessMethodClient) ListAuthenticationWindowsHelloForBusinessMethodsComplete(ctx context.Context, options ListAuthenticationWindowsHelloForBusinessMethodsOperationOptions) (ListAuthenticationWindowsHelloForBusinessMethodsCompleteResult, error) {
	return c.ListAuthenticationWindowsHelloForBusinessMethodsCompleteMatchingPredicate(ctx, options, WindowsHelloForBusinessAuthenticationMethodOperationPredicate{})
}

// ListAuthenticationWindowsHelloForBusinessMethodsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AuthenticationWindowsHelloForBusinessMethodClient) ListAuthenticationWindowsHelloForBusinessMethodsCompleteMatchingPredicate(ctx context.Context, options ListAuthenticationWindowsHelloForBusinessMethodsOperationOptions, predicate WindowsHelloForBusinessAuthenticationMethodOperationPredicate) (result ListAuthenticationWindowsHelloForBusinessMethodsCompleteResult, err error) {
	items := make([]stable.WindowsHelloForBusinessAuthenticationMethod, 0)

	resp, err := c.ListAuthenticationWindowsHelloForBusinessMethods(ctx, options)
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

	result = ListAuthenticationWindowsHelloForBusinessMethodsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
