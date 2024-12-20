package webapps

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

type ListConfigurationSnapshotInfoOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]SiteConfigurationSnapshotInfo
}

type ListConfigurationSnapshotInfoCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []SiteConfigurationSnapshotInfo
}

type ListConfigurationSnapshotInfoCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListConfigurationSnapshotInfoCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListConfigurationSnapshotInfo ...
func (c WebAppsClient) ListConfigurationSnapshotInfo(ctx context.Context, id commonids.AppServiceId) (result ListConfigurationSnapshotInfoOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListConfigurationSnapshotInfoCustomPager{},
		Path:       fmt.Sprintf("%s/config/web/snapshots", id.ID()),
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
		Values *[]SiteConfigurationSnapshotInfo `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListConfigurationSnapshotInfoComplete retrieves all the results into a single object
func (c WebAppsClient) ListConfigurationSnapshotInfoComplete(ctx context.Context, id commonids.AppServiceId) (ListConfigurationSnapshotInfoCompleteResult, error) {
	return c.ListConfigurationSnapshotInfoCompleteMatchingPredicate(ctx, id, SiteConfigurationSnapshotInfoOperationPredicate{})
}

// ListConfigurationSnapshotInfoCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c WebAppsClient) ListConfigurationSnapshotInfoCompleteMatchingPredicate(ctx context.Context, id commonids.AppServiceId, predicate SiteConfigurationSnapshotInfoOperationPredicate) (result ListConfigurationSnapshotInfoCompleteResult, err error) {
	items := make([]SiteConfigurationSnapshotInfo, 0)

	resp, err := c.ListConfigurationSnapshotInfo(ctx, id)
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

	result = ListConfigurationSnapshotInfoCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
