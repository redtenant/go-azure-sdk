package healthbots

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

type BotsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]HealthBot
}

type BotsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []HealthBot
}

type BotsListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *BotsListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// BotsList ...
func (c HealthbotsClient) BotsList(ctx context.Context, id commonids.SubscriptionId) (result BotsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &BotsListCustomPager{},
		Path:       fmt.Sprintf("%s/providers/Microsoft.HealthBot/healthBots", id.ID()),
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
		Values *[]HealthBot `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// BotsListComplete retrieves all the results into a single object
func (c HealthbotsClient) BotsListComplete(ctx context.Context, id commonids.SubscriptionId) (BotsListCompleteResult, error) {
	return c.BotsListCompleteMatchingPredicate(ctx, id, HealthBotOperationPredicate{})
}

// BotsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c HealthbotsClient) BotsListCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, predicate HealthBotOperationPredicate) (result BotsListCompleteResult, err error) {
	items := make([]HealthBot, 0)

	resp, err := c.BotsList(ctx, id)
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

	result = BotsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
