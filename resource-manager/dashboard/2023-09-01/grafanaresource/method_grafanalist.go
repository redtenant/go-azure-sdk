package grafanaresource

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

type GrafanaListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ManagedGrafana
}

type GrafanaListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ManagedGrafana
}

type GrafanaListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *GrafanaListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// GrafanaList ...
func (c GrafanaResourceClient) GrafanaList(ctx context.Context, id commonids.SubscriptionId) (result GrafanaListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &GrafanaListCustomPager{},
		Path:       fmt.Sprintf("%s/providers/Microsoft.Dashboard/grafana", id.ID()),
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
		Values *[]ManagedGrafana `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// GrafanaListComplete retrieves all the results into a single object
func (c GrafanaResourceClient) GrafanaListComplete(ctx context.Context, id commonids.SubscriptionId) (GrafanaListCompleteResult, error) {
	return c.GrafanaListCompleteMatchingPredicate(ctx, id, ManagedGrafanaOperationPredicate{})
}

// GrafanaListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GrafanaResourceClient) GrafanaListCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, predicate ManagedGrafanaOperationPredicate) (result GrafanaListCompleteResult, err error) {
	items := make([]ManagedGrafana, 0)

	resp, err := c.GrafanaList(ctx, id)
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

	result = GrafanaListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
