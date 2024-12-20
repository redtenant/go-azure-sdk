package virtualwans

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

type VpnSitesListByResourceGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]VpnSite
}

type VpnSitesListByResourceGroupCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []VpnSite
}

type VpnSitesListByResourceGroupCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *VpnSitesListByResourceGroupCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// VpnSitesListByResourceGroup ...
func (c VirtualWANsClient) VpnSitesListByResourceGroup(ctx context.Context, id commonids.ResourceGroupId) (result VpnSitesListByResourceGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &VpnSitesListByResourceGroupCustomPager{},
		Path:       fmt.Sprintf("%s/providers/Microsoft.Network/vpnSites", id.ID()),
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
		Values *[]VpnSite `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// VpnSitesListByResourceGroupComplete retrieves all the results into a single object
func (c VirtualWANsClient) VpnSitesListByResourceGroupComplete(ctx context.Context, id commonids.ResourceGroupId) (VpnSitesListByResourceGroupCompleteResult, error) {
	return c.VpnSitesListByResourceGroupCompleteMatchingPredicate(ctx, id, VpnSiteOperationPredicate{})
}

// VpnSitesListByResourceGroupCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c VirtualWANsClient) VpnSitesListByResourceGroupCompleteMatchingPredicate(ctx context.Context, id commonids.ResourceGroupId, predicate VpnSiteOperationPredicate) (result VpnSitesListByResourceGroupCompleteResult, err error) {
	items := make([]VpnSite, 0)

	resp, err := c.VpnSitesListByResourceGroup(ctx, id)
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

	result = VpnSitesListByResourceGroupCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
