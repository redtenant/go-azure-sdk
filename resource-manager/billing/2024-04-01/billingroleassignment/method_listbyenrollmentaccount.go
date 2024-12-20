package billingroleassignment

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByEnrollmentAccountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]BillingRoleAssignment
}

type ListByEnrollmentAccountCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []BillingRoleAssignment
}

type ListByEnrollmentAccountCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByEnrollmentAccountCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByEnrollmentAccount ...
func (c BillingRoleAssignmentClient) ListByEnrollmentAccount(ctx context.Context, id EnrollmentAccountId) (result ListByEnrollmentAccountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListByEnrollmentAccountCustomPager{},
		Path:       fmt.Sprintf("%s/billingRoleAssignments", id.ID()),
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
		Values *[]BillingRoleAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByEnrollmentAccountComplete retrieves all the results into a single object
func (c BillingRoleAssignmentClient) ListByEnrollmentAccountComplete(ctx context.Context, id EnrollmentAccountId) (ListByEnrollmentAccountCompleteResult, error) {
	return c.ListByEnrollmentAccountCompleteMatchingPredicate(ctx, id, BillingRoleAssignmentOperationPredicate{})
}

// ListByEnrollmentAccountCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c BillingRoleAssignmentClient) ListByEnrollmentAccountCompleteMatchingPredicate(ctx context.Context, id EnrollmentAccountId, predicate BillingRoleAssignmentOperationPredicate) (result ListByEnrollmentAccountCompleteResult, err error) {
	items := make([]BillingRoleAssignment, 0)

	resp, err := c.ListByEnrollmentAccount(ctx, id)
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

	result = ListByEnrollmentAccountCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
