package containerappspatches

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByContainerAppOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ContainerAppsPatchResource
}

type ListByContainerAppCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ContainerAppsPatchResource
}

type ListByContainerAppOperationOptions struct {
	Filter *string
}

func DefaultListByContainerAppOperationOptions() ListByContainerAppOperationOptions {
	return ListByContainerAppOperationOptions{}
}

func (o ListByContainerAppOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListByContainerAppOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ListByContainerAppOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	return &out
}

type ListByContainerAppCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByContainerAppCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByContainerApp ...
func (c ContainerAppsPatchesClient) ListByContainerApp(ctx context.Context, id ContainerAppId, options ListByContainerAppOperationOptions) (result ListByContainerAppOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListByContainerAppCustomPager{},
		Path:          fmt.Sprintf("%s/patches", id.ID()),
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
		Values *[]ContainerAppsPatchResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByContainerAppComplete retrieves all the results into a single object
func (c ContainerAppsPatchesClient) ListByContainerAppComplete(ctx context.Context, id ContainerAppId, options ListByContainerAppOperationOptions) (ListByContainerAppCompleteResult, error) {
	return c.ListByContainerAppCompleteMatchingPredicate(ctx, id, options, ContainerAppsPatchResourceOperationPredicate{})
}

// ListByContainerAppCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ContainerAppsPatchesClient) ListByContainerAppCompleteMatchingPredicate(ctx context.Context, id ContainerAppId, options ListByContainerAppOperationOptions, predicate ContainerAppsPatchResourceOperationPredicate) (result ListByContainerAppCompleteResult, err error) {
	items := make([]ContainerAppsPatchResource, 0)

	resp, err := c.ListByContainerApp(ctx, id, options)
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

	result = ListByContainerAppCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
