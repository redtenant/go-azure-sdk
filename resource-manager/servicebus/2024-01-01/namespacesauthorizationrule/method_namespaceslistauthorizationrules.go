package namespacesauthorizationrule

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NamespacesListAuthorizationRulesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]SBAuthorizationRule
}

type NamespacesListAuthorizationRulesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []SBAuthorizationRule
}

type NamespacesListAuthorizationRulesCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *NamespacesListAuthorizationRulesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// NamespacesListAuthorizationRules ...
func (c NamespacesAuthorizationRuleClient) NamespacesListAuthorizationRules(ctx context.Context, id NamespaceId) (result NamespacesListAuthorizationRulesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &NamespacesListAuthorizationRulesCustomPager{},
		Path:       fmt.Sprintf("%s/authorizationRules", id.ID()),
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
		Values *[]SBAuthorizationRule `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// NamespacesListAuthorizationRulesComplete retrieves all the results into a single object
func (c NamespacesAuthorizationRuleClient) NamespacesListAuthorizationRulesComplete(ctx context.Context, id NamespaceId) (NamespacesListAuthorizationRulesCompleteResult, error) {
	return c.NamespacesListAuthorizationRulesCompleteMatchingPredicate(ctx, id, SBAuthorizationRuleOperationPredicate{})
}

// NamespacesListAuthorizationRulesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c NamespacesAuthorizationRuleClient) NamespacesListAuthorizationRulesCompleteMatchingPredicate(ctx context.Context, id NamespaceId, predicate SBAuthorizationRuleOperationPredicate) (result NamespacesListAuthorizationRulesCompleteResult, err error) {
	items := make([]SBAuthorizationRule, 0)

	resp, err := c.NamespacesListAuthorizationRules(ctx, id)
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

	result = NamespacesListAuthorizationRulesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
