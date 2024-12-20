package tagoperationlink

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspaceTagOperationLinkListByProductOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]TagOperationLinkContract
}

type WorkspaceTagOperationLinkListByProductCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []TagOperationLinkContract
}

type WorkspaceTagOperationLinkListByProductOperationOptions struct {
	Filter *string
	Skip   *int64
	Top    *int64
}

func DefaultWorkspaceTagOperationLinkListByProductOperationOptions() WorkspaceTagOperationLinkListByProductOperationOptions {
	return WorkspaceTagOperationLinkListByProductOperationOptions{}
}

func (o WorkspaceTagOperationLinkListByProductOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o WorkspaceTagOperationLinkListByProductOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o WorkspaceTagOperationLinkListByProductOperationOptions) ToQuery() *client.QueryParams {
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

type WorkspaceTagOperationLinkListByProductCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *WorkspaceTagOperationLinkListByProductCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// WorkspaceTagOperationLinkListByProduct ...
func (c TagOperationLinkClient) WorkspaceTagOperationLinkListByProduct(ctx context.Context, id WorkspaceTagId, options WorkspaceTagOperationLinkListByProductOperationOptions) (result WorkspaceTagOperationLinkListByProductOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &WorkspaceTagOperationLinkListByProductCustomPager{},
		Path:          fmt.Sprintf("%s/operationLinks", id.ID()),
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
		Values *[]TagOperationLinkContract `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// WorkspaceTagOperationLinkListByProductComplete retrieves all the results into a single object
func (c TagOperationLinkClient) WorkspaceTagOperationLinkListByProductComplete(ctx context.Context, id WorkspaceTagId, options WorkspaceTagOperationLinkListByProductOperationOptions) (WorkspaceTagOperationLinkListByProductCompleteResult, error) {
	return c.WorkspaceTagOperationLinkListByProductCompleteMatchingPredicate(ctx, id, options, TagOperationLinkContractOperationPredicate{})
}

// WorkspaceTagOperationLinkListByProductCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TagOperationLinkClient) WorkspaceTagOperationLinkListByProductCompleteMatchingPredicate(ctx context.Context, id WorkspaceTagId, options WorkspaceTagOperationLinkListByProductOperationOptions, predicate TagOperationLinkContractOperationPredicate) (result WorkspaceTagOperationLinkListByProductCompleteResult, err error) {
	items := make([]TagOperationLinkContract, 0)

	resp, err := c.WorkspaceTagOperationLinkListByProduct(ctx, id, options)
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

	result = WorkspaceTagOperationLinkListByProductCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
