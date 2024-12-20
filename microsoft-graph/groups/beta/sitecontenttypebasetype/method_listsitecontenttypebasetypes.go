package sitecontenttypebasetype

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

type ListSiteContentTypeBaseTypesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ContentType
}

type ListSiteContentTypeBaseTypesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ContentType
}

type ListSiteContentTypeBaseTypesOperationOptions struct {
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

func DefaultListSiteContentTypeBaseTypesOperationOptions() ListSiteContentTypeBaseTypesOperationOptions {
	return ListSiteContentTypeBaseTypesOperationOptions{}
}

func (o ListSiteContentTypeBaseTypesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListSiteContentTypeBaseTypesOperationOptions) ToOData() *odata.Query {
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

func (o ListSiteContentTypeBaseTypesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListSiteContentTypeBaseTypesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSiteContentTypeBaseTypesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSiteContentTypeBaseTypes - Get baseTypes from groups. The collection of content types that are ancestors of this
// content type.
func (c SiteContentTypeBaseTypeClient) ListSiteContentTypeBaseTypes(ctx context.Context, id beta.GroupIdSiteIdContentTypeId, options ListSiteContentTypeBaseTypesOperationOptions) (result ListSiteContentTypeBaseTypesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListSiteContentTypeBaseTypesCustomPager{},
		Path:          fmt.Sprintf("%s/baseTypes", id.ID()),
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
		Values *[]beta.ContentType `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSiteContentTypeBaseTypesComplete retrieves all the results into a single object
func (c SiteContentTypeBaseTypeClient) ListSiteContentTypeBaseTypesComplete(ctx context.Context, id beta.GroupIdSiteIdContentTypeId, options ListSiteContentTypeBaseTypesOperationOptions) (ListSiteContentTypeBaseTypesCompleteResult, error) {
	return c.ListSiteContentTypeBaseTypesCompleteMatchingPredicate(ctx, id, options, ContentTypeOperationPredicate{})
}

// ListSiteContentTypeBaseTypesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteContentTypeBaseTypeClient) ListSiteContentTypeBaseTypesCompleteMatchingPredicate(ctx context.Context, id beta.GroupIdSiteIdContentTypeId, options ListSiteContentTypeBaseTypesOperationOptions, predicate ContentTypeOperationPredicate) (result ListSiteContentTypeBaseTypesCompleteResult, err error) {
	items := make([]beta.ContentType, 0)

	resp, err := c.ListSiteContentTypeBaseTypes(ctx, id, options)
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

	result = ListSiteContentTypeBaseTypesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
