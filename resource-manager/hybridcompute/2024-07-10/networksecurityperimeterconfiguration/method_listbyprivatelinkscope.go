package networksecurityperimeterconfiguration

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByPrivateLinkScopeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]NetworkSecurityPerimeterConfiguration
}

type ListByPrivateLinkScopeCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []NetworkSecurityPerimeterConfiguration
}

type ListByPrivateLinkScopeCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByPrivateLinkScopeCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByPrivateLinkScope ...
func (c NetworkSecurityPerimeterConfigurationClient) ListByPrivateLinkScope(ctx context.Context, id ProviderPrivateLinkScopeId) (result ListByPrivateLinkScopeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListByPrivateLinkScopeCustomPager{},
		Path:       fmt.Sprintf("%s/networkSecurityPerimeterConfigurations", id.ID()),
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
		Values *[]NetworkSecurityPerimeterConfiguration `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByPrivateLinkScopeComplete retrieves all the results into a single object
func (c NetworkSecurityPerimeterConfigurationClient) ListByPrivateLinkScopeComplete(ctx context.Context, id ProviderPrivateLinkScopeId) (ListByPrivateLinkScopeCompleteResult, error) {
	return c.ListByPrivateLinkScopeCompleteMatchingPredicate(ctx, id, NetworkSecurityPerimeterConfigurationOperationPredicate{})
}

// ListByPrivateLinkScopeCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c NetworkSecurityPerimeterConfigurationClient) ListByPrivateLinkScopeCompleteMatchingPredicate(ctx context.Context, id ProviderPrivateLinkScopeId, predicate NetworkSecurityPerimeterConfigurationOperationPredicate) (result ListByPrivateLinkScopeCompleteResult, err error) {
	items := make([]NetworkSecurityPerimeterConfiguration, 0)

	resp, err := c.ListByPrivateLinkScope(ctx, id)
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

	result = ListByPrivateLinkScopeCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
