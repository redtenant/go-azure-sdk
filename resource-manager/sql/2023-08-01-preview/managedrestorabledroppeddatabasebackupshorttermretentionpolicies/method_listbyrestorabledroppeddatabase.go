package managedrestorabledroppeddatabasebackupshorttermretentionpolicies

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByRestorableDroppedDatabaseOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ManagedBackupShortTermRetentionPolicy
}

type ListByRestorableDroppedDatabaseCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ManagedBackupShortTermRetentionPolicy
}

type ListByRestorableDroppedDatabaseCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByRestorableDroppedDatabaseCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByRestorableDroppedDatabase ...
func (c ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient) ListByRestorableDroppedDatabase(ctx context.Context, id ManagedInstanceRestorableDroppedDatabaseId) (result ListByRestorableDroppedDatabaseOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListByRestorableDroppedDatabaseCustomPager{},
		Path:       fmt.Sprintf("%s/backupShortTermRetentionPolicies", id.ID()),
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
		Values *[]ManagedBackupShortTermRetentionPolicy `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByRestorableDroppedDatabaseComplete retrieves all the results into a single object
func (c ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient) ListByRestorableDroppedDatabaseComplete(ctx context.Context, id ManagedInstanceRestorableDroppedDatabaseId) (ListByRestorableDroppedDatabaseCompleteResult, error) {
	return c.ListByRestorableDroppedDatabaseCompleteMatchingPredicate(ctx, id, ManagedBackupShortTermRetentionPolicyOperationPredicate{})
}

// ListByRestorableDroppedDatabaseCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient) ListByRestorableDroppedDatabaseCompleteMatchingPredicate(ctx context.Context, id ManagedInstanceRestorableDroppedDatabaseId, predicate ManagedBackupShortTermRetentionPolicyOperationPredicate) (result ListByRestorableDroppedDatabaseCompleteResult, err error) {
	items := make([]ManagedBackupShortTermRetentionPolicy, 0)

	resp, err := c.ListByRestorableDroppedDatabase(ctx, id)
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

	result = ListByRestorableDroppedDatabaseCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
