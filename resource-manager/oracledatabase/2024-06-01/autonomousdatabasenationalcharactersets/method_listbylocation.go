package autonomousdatabasenationalcharactersets

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByLocationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]AutonomousDatabaseNationalCharacterSet
}

type ListByLocationCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []AutonomousDatabaseNationalCharacterSet
}

type ListByLocationCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByLocationCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByLocation ...
func (c AutonomousDatabaseNationalCharacterSetsClient) ListByLocation(ctx context.Context, id LocationId) (result ListByLocationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListByLocationCustomPager{},
		Path:       fmt.Sprintf("%s/autonomousDatabaseNationalCharacterSets", id.ID()),
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
		Values *[]AutonomousDatabaseNationalCharacterSet `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByLocationComplete retrieves all the results into a single object
func (c AutonomousDatabaseNationalCharacterSetsClient) ListByLocationComplete(ctx context.Context, id LocationId) (ListByLocationCompleteResult, error) {
	return c.ListByLocationCompleteMatchingPredicate(ctx, id, AutonomousDatabaseNationalCharacterSetOperationPredicate{})
}

// ListByLocationCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AutonomousDatabaseNationalCharacterSetsClient) ListByLocationCompleteMatchingPredicate(ctx context.Context, id LocationId, predicate AutonomousDatabaseNationalCharacterSetOperationPredicate) (result ListByLocationCompleteResult, err error) {
	items := make([]AutonomousDatabaseNationalCharacterSet, 0)

	resp, err := c.ListByLocation(ctx, id)
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

	result = ListByLocationCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
