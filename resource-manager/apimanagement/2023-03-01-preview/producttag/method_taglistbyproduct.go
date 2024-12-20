package producttag

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TagListByProductOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]TagContract
}

type TagListByProductCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []TagContract
}

type TagListByProductOperationOptions struct {
	Filter *string
	Skip   *int64
	Top    *int64
}

func DefaultTagListByProductOperationOptions() TagListByProductOperationOptions {
	return TagListByProductOperationOptions{}
}

func (o TagListByProductOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o TagListByProductOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o TagListByProductOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.Skip != nil {
		out.Append("$skip", fmt.Sprintf("%v", *o.Skip))
	}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

type TagListByProductCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *TagListByProductCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// TagListByProduct ...
func (c ProductTagClient) TagListByProduct(ctx context.Context, id ProductId, options TagListByProductOperationOptions) (result TagListByProductOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &TagListByProductCustomPager{},
		Path:          fmt.Sprintf("%s/tags", id.ID()),
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
		Values *[]TagContract `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// TagListByProductComplete retrieves all the results into a single object
func (c ProductTagClient) TagListByProductComplete(ctx context.Context, id ProductId, options TagListByProductOperationOptions) (TagListByProductCompleteResult, error) {
	return c.TagListByProductCompleteMatchingPredicate(ctx, id, options, TagContractOperationPredicate{})
}

// TagListByProductCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ProductTagClient) TagListByProductCompleteMatchingPredicate(ctx context.Context, id ProductId, options TagListByProductOperationOptions, predicate TagContractOperationPredicate) (result TagListByProductCompleteResult, err error) {
	items := make([]TagContract, 0)

	resp, err := c.TagListByProduct(ctx, id, options)
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

	result = TagListByProductCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
