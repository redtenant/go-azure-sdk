package hybrididentitymetadata

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByMachinesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]HybridIdentityMetadata
}

type ListByMachinesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []HybridIdentityMetadata
}

type ListByMachinesCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByMachinesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByMachines ...
func (c HybridIdentityMetadataClient) ListByMachines(ctx context.Context, id MachineId) (result ListByMachinesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListByMachinesCustomPager{},
		Path:       fmt.Sprintf("%s/hybridIdentityMetadata", id.ID()),
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
		Values *[]HybridIdentityMetadata `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByMachinesComplete retrieves all the results into a single object
func (c HybridIdentityMetadataClient) ListByMachinesComplete(ctx context.Context, id MachineId) (ListByMachinesCompleteResult, error) {
	return c.ListByMachinesCompleteMatchingPredicate(ctx, id, HybridIdentityMetadataOperationPredicate{})
}

// ListByMachinesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c HybridIdentityMetadataClient) ListByMachinesCompleteMatchingPredicate(ctx context.Context, id MachineId, predicate HybridIdentityMetadataOperationPredicate) (result ListByMachinesCompleteResult, err error) {
	items := make([]HybridIdentityMetadata, 0)

	resp, err := c.ListByMachines(ctx, id)
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

	result = ListByMachinesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
