package iscsitargets

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByDiskPoolOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]IscsiTarget
}

type ListByDiskPoolCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []IscsiTarget
}

type ListByDiskPoolCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByDiskPoolCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByDiskPool ...
func (c IscsiTargetsClient) ListByDiskPool(ctx context.Context, id DiskPoolId) (result ListByDiskPoolOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListByDiskPoolCustomPager{},
		Path:       fmt.Sprintf("%s/iscsiTargets", id.ID()),
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
		Values *[]IscsiTarget `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByDiskPoolComplete retrieves all the results into a single object
func (c IscsiTargetsClient) ListByDiskPoolComplete(ctx context.Context, id DiskPoolId) (ListByDiskPoolCompleteResult, error) {
	return c.ListByDiskPoolCompleteMatchingPredicate(ctx, id, IscsiTargetOperationPredicate{})
}

// ListByDiskPoolCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IscsiTargetsClient) ListByDiskPoolCompleteMatchingPredicate(ctx context.Context, id DiskPoolId, predicate IscsiTargetOperationPredicate) (result ListByDiskPoolCompleteResult, err error) {
	items := make([]IscsiTarget, 0)

	resp, err := c.ListByDiskPool(ctx, id)
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

	result = ListByDiskPoolCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
