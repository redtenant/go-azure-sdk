package managedhsms

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListDeletedOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]DeletedManagedHsm
}

type ListDeletedCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []DeletedManagedHsm
}

type ListDeletedCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListDeletedCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeleted ...
func (c ManagedHsmsClient) ListDeleted(ctx context.Context, id commonids.SubscriptionId) (result ListDeletedOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListDeletedCustomPager{},
		Path:       fmt.Sprintf("%s/providers/Microsoft.KeyVault/deletedManagedHSMs", id.ID()),
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
		Values *[]DeletedManagedHsm `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDeletedComplete retrieves all the results into a single object
func (c ManagedHsmsClient) ListDeletedComplete(ctx context.Context, id commonids.SubscriptionId) (ListDeletedCompleteResult, error) {
	return c.ListDeletedCompleteMatchingPredicate(ctx, id, DeletedManagedHsmOperationPredicate{})
}

// ListDeletedCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ManagedHsmsClient) ListDeletedCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, predicate DeletedManagedHsmOperationPredicate) (result ListDeletedCompleteResult, err error) {
	items := make([]DeletedManagedHsm, 0)

	resp, err := c.ListDeleted(ctx, id)
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

	result = ListDeletedCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
