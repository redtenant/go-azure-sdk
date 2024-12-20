package fhirservices

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByWorkspaceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]FhirService
}

type ListByWorkspaceCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []FhirService
}

type ListByWorkspaceCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByWorkspaceCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByWorkspace ...
func (c FhirServicesClient) ListByWorkspace(ctx context.Context, id WorkspaceId) (result ListByWorkspaceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListByWorkspaceCustomPager{},
		Path:       fmt.Sprintf("%s/fhirServices", id.ID()),
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
		Values *[]FhirService `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByWorkspaceComplete retrieves all the results into a single object
func (c FhirServicesClient) ListByWorkspaceComplete(ctx context.Context, id WorkspaceId) (ListByWorkspaceCompleteResult, error) {
	return c.ListByWorkspaceCompleteMatchingPredicate(ctx, id, FhirServiceOperationPredicate{})
}

// ListByWorkspaceCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c FhirServicesClient) ListByWorkspaceCompleteMatchingPredicate(ctx context.Context, id WorkspaceId, predicate FhirServiceOperationPredicate) (result ListByWorkspaceCompleteResult, err error) {
	items := make([]FhirService, 0)

	resp, err := c.ListByWorkspace(ctx, id)
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

	result = ListByWorkspaceCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
