package quotainformation

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

type QuotaListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]CurrentQuotaLimitBase
}

type QuotaListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []CurrentQuotaLimitBase
}

type QuotaListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *QuotaListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// QuotaList ...
func (c QuotaInformationClient) QuotaList(ctx context.Context, id commonids.ScopeId) (result QuotaListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &QuotaListCustomPager{},
		Path:       fmt.Sprintf("%s/providers/Microsoft.Quota/quotas", id.ID()),
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
		Values *[]CurrentQuotaLimitBase `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// QuotaListComplete retrieves all the results into a single object
func (c QuotaInformationClient) QuotaListComplete(ctx context.Context, id commonids.ScopeId) (QuotaListCompleteResult, error) {
	return c.QuotaListCompleteMatchingPredicate(ctx, id, CurrentQuotaLimitBaseOperationPredicate{})
}

// QuotaListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c QuotaInformationClient) QuotaListCompleteMatchingPredicate(ctx context.Context, id commonids.ScopeId, predicate CurrentQuotaLimitBaseOperationPredicate) (result QuotaListCompleteResult, err error) {
	items := make([]CurrentQuotaLimitBase, 0)

	resp, err := c.QuotaList(ctx, id)
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

	result = QuotaListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
