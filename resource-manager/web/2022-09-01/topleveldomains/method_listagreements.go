package topleveldomains

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListAgreementsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]TldLegalAgreement
}

type ListAgreementsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []TldLegalAgreement
}

type ListAgreementsCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListAgreementsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAgreements ...
func (c TopLevelDomainsClient) ListAgreements(ctx context.Context, id TopLevelDomainId, input TopLevelDomainAgreementOption) (result ListAgreementsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Pager:      &ListAgreementsCustomPager{},
		Path:       fmt.Sprintf("%s/listAgreements", id.ID()),
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
		Values *[]TldLegalAgreement `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAgreementsComplete retrieves all the results into a single object
func (c TopLevelDomainsClient) ListAgreementsComplete(ctx context.Context, id TopLevelDomainId, input TopLevelDomainAgreementOption) (ListAgreementsCompleteResult, error) {
	return c.ListAgreementsCompleteMatchingPredicate(ctx, id, input, TldLegalAgreementOperationPredicate{})
}

// ListAgreementsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TopLevelDomainsClient) ListAgreementsCompleteMatchingPredicate(ctx context.Context, id TopLevelDomainId, input TopLevelDomainAgreementOption, predicate TldLegalAgreementOperationPredicate) (result ListAgreementsCompleteResult, err error) {
	items := make([]TldLegalAgreement, 0)

	resp, err := c.ListAgreements(ctx, id, input)
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

	result = ListAgreementsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
