package entitlementmanagementroleassignmentschedule

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListEntitlementManagementRoleAssignmentSchedulesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UnifiedRoleAssignmentSchedule
}

type ListEntitlementManagementRoleAssignmentSchedulesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UnifiedRoleAssignmentSchedule
}

type ListEntitlementManagementRoleAssignmentSchedulesOperationOptions struct {
	Count     *bool
	Expand    *odata.Expand
	Filter    *string
	Metadata  *odata.Metadata
	OrderBy   *odata.OrderBy
	RetryFunc client.RequestRetryFunc
	Search    *string
	Select    *[]string
	Skip      *int64
	Top       *int64
}

func DefaultListEntitlementManagementRoleAssignmentSchedulesOperationOptions() ListEntitlementManagementRoleAssignmentSchedulesOperationOptions {
	return ListEntitlementManagementRoleAssignmentSchedulesOperationOptions{}
}

func (o ListEntitlementManagementRoleAssignmentSchedulesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEntitlementManagementRoleAssignmentSchedulesOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Count != nil {
		out.Count = *o.Count
	}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.OrderBy != nil {
		out.OrderBy = *o.OrderBy
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListEntitlementManagementRoleAssignmentSchedulesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEntitlementManagementRoleAssignmentSchedulesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementRoleAssignmentSchedulesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementRoleAssignmentSchedules - Get roleAssignmentSchedules from roleManagement
func (c EntitlementManagementRoleAssignmentScheduleClient) ListEntitlementManagementRoleAssignmentSchedules(ctx context.Context, options ListEntitlementManagementRoleAssignmentSchedulesOperationOptions) (result ListEntitlementManagementRoleAssignmentSchedulesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEntitlementManagementRoleAssignmentSchedulesCustomPager{},
		Path:          "/roleManagement/entitlementManagement/roleAssignmentSchedules",
		RetryFunc:     options.RetryFunc,
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
		Values *[]beta.UnifiedRoleAssignmentSchedule `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEntitlementManagementRoleAssignmentSchedulesComplete retrieves all the results into a single object
func (c EntitlementManagementRoleAssignmentScheduleClient) ListEntitlementManagementRoleAssignmentSchedulesComplete(ctx context.Context, options ListEntitlementManagementRoleAssignmentSchedulesOperationOptions) (ListEntitlementManagementRoleAssignmentSchedulesCompleteResult, error) {
	return c.ListEntitlementManagementRoleAssignmentSchedulesCompleteMatchingPredicate(ctx, options, UnifiedRoleAssignmentScheduleOperationPredicate{})
}

// ListEntitlementManagementRoleAssignmentSchedulesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementRoleAssignmentScheduleClient) ListEntitlementManagementRoleAssignmentSchedulesCompleteMatchingPredicate(ctx context.Context, options ListEntitlementManagementRoleAssignmentSchedulesOperationOptions, predicate UnifiedRoleAssignmentScheduleOperationPredicate) (result ListEntitlementManagementRoleAssignmentSchedulesCompleteResult, err error) {
	items := make([]beta.UnifiedRoleAssignmentSchedule, 0)

	resp, err := c.ListEntitlementManagementRoleAssignmentSchedules(ctx, options)
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

	result = ListEntitlementManagementRoleAssignmentSchedulesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
