package sitelistcontenttypecolumn

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

type ListSiteListContentTypeColumnsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ColumnDefinition
}

type ListSiteListContentTypeColumnsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ColumnDefinition
}

type ListSiteListContentTypeColumnsOperationOptions struct {
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

func DefaultListSiteListContentTypeColumnsOperationOptions() ListSiteListContentTypeColumnsOperationOptions {
	return ListSiteListContentTypeColumnsOperationOptions{}
}

func (o ListSiteListContentTypeColumnsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListSiteListContentTypeColumnsOperationOptions) ToOData() *odata.Query {
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

func (o ListSiteListContentTypeColumnsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListSiteListContentTypeColumnsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSiteListContentTypeColumnsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSiteListContentTypeColumns - Get columns from groups. The collection of column definitions for this content type.
func (c SiteListContentTypeColumnClient) ListSiteListContentTypeColumns(ctx context.Context, id stable.GroupIdSiteIdListIdContentTypeId, options ListSiteListContentTypeColumnsOperationOptions) (result ListSiteListContentTypeColumnsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListSiteListContentTypeColumnsCustomPager{},
		Path:          fmt.Sprintf("%s/columns", id.ID()),
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
		Values *[]stable.ColumnDefinition `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSiteListContentTypeColumnsComplete retrieves all the results into a single object
func (c SiteListContentTypeColumnClient) ListSiteListContentTypeColumnsComplete(ctx context.Context, id stable.GroupIdSiteIdListIdContentTypeId, options ListSiteListContentTypeColumnsOperationOptions) (ListSiteListContentTypeColumnsCompleteResult, error) {
	return c.ListSiteListContentTypeColumnsCompleteMatchingPredicate(ctx, id, options, ColumnDefinitionOperationPredicate{})
}

// ListSiteListContentTypeColumnsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteListContentTypeColumnClient) ListSiteListContentTypeColumnsCompleteMatchingPredicate(ctx context.Context, id stable.GroupIdSiteIdListIdContentTypeId, options ListSiteListContentTypeColumnsOperationOptions, predicate ColumnDefinitionOperationPredicate) (result ListSiteListContentTypeColumnsCompleteResult, err error) {
	items := make([]stable.ColumnDefinition, 0)

	resp, err := c.ListSiteListContentTypeColumns(ctx, id, options)
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

	result = ListSiteListContentTypeColumnsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
