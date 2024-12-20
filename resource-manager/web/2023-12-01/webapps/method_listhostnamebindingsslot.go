package webapps

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListHostNameBindingsSlotOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]HostNameBinding
}

type ListHostNameBindingsSlotCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []HostNameBinding
}

type ListHostNameBindingsSlotCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListHostNameBindingsSlotCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListHostNameBindingsSlot ...
func (c WebAppsClient) ListHostNameBindingsSlot(ctx context.Context, id SlotId) (result ListHostNameBindingsSlotOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListHostNameBindingsSlotCustomPager{},
		Path:       fmt.Sprintf("%s/hostNameBindings", id.ID()),
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
		Values *[]HostNameBinding `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListHostNameBindingsSlotComplete retrieves all the results into a single object
func (c WebAppsClient) ListHostNameBindingsSlotComplete(ctx context.Context, id SlotId) (ListHostNameBindingsSlotCompleteResult, error) {
	return c.ListHostNameBindingsSlotCompleteMatchingPredicate(ctx, id, HostNameBindingOperationPredicate{})
}

// ListHostNameBindingsSlotCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c WebAppsClient) ListHostNameBindingsSlotCompleteMatchingPredicate(ctx context.Context, id SlotId, predicate HostNameBindingOperationPredicate) (result ListHostNameBindingsSlotCompleteResult, err error) {
	items := make([]HostNameBinding, 0)

	resp, err := c.ListHostNameBindingsSlot(ctx, id)
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

	result = ListHostNameBindingsSlotCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
