package workloadnetworks

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListVirtualMachinesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]WorkloadNetworkVirtualMachine
}

type ListVirtualMachinesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []WorkloadNetworkVirtualMachine
}

type ListVirtualMachinesCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListVirtualMachinesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListVirtualMachines ...
func (c WorkloadNetworksClient) ListVirtualMachines(ctx context.Context, id PrivateCloudId) (result ListVirtualMachinesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListVirtualMachinesCustomPager{},
		Path:       fmt.Sprintf("%s/workloadNetworks/default/virtualMachines", id.ID()),
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
		Values *[]WorkloadNetworkVirtualMachine `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListVirtualMachinesComplete retrieves all the results into a single object
func (c WorkloadNetworksClient) ListVirtualMachinesComplete(ctx context.Context, id PrivateCloudId) (ListVirtualMachinesCompleteResult, error) {
	return c.ListVirtualMachinesCompleteMatchingPredicate(ctx, id, WorkloadNetworkVirtualMachineOperationPredicate{})
}

// ListVirtualMachinesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c WorkloadNetworksClient) ListVirtualMachinesCompleteMatchingPredicate(ctx context.Context, id PrivateCloudId, predicate WorkloadNetworkVirtualMachineOperationPredicate) (result ListVirtualMachinesCompleteResult, err error) {
	items := make([]WorkloadNetworkVirtualMachine, 0)

	resp, err := c.ListVirtualMachines(ctx, id)
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

	result = ListVirtualMachinesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
