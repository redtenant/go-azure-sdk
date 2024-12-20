package assessmentsmetadata

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

type SubscriptionListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]SecurityAssessmentMetadata
}

type SubscriptionListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []SecurityAssessmentMetadata
}

type SubscriptionListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *SubscriptionListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// SubscriptionList ...
func (c AssessmentsMetadataClient) SubscriptionList(ctx context.Context, id commonids.SubscriptionId) (result SubscriptionListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &SubscriptionListCustomPager{},
		Path:       fmt.Sprintf("%s/providers/Microsoft.Security/assessmentMetadata", id.ID()),
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
		Values *[]SecurityAssessmentMetadata `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// SubscriptionListComplete retrieves all the results into a single object
func (c AssessmentsMetadataClient) SubscriptionListComplete(ctx context.Context, id commonids.SubscriptionId) (SubscriptionListCompleteResult, error) {
	return c.SubscriptionListCompleteMatchingPredicate(ctx, id, SecurityAssessmentMetadataOperationPredicate{})
}

// SubscriptionListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AssessmentsMetadataClient) SubscriptionListCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, predicate SecurityAssessmentMetadataOperationPredicate) (result SubscriptionListCompleteResult, err error) {
	items := make([]SecurityAssessmentMetadata, 0)

	resp, err := c.SubscriptionList(ctx, id)
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

	result = SubscriptionListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
