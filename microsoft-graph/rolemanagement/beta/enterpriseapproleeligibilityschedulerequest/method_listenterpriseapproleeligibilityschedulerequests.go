package enterpriseapproleeligibilityschedulerequest

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

type ListEnterpriseAppRoleEligibilityScheduleRequestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UnifiedRoleEligibilityScheduleRequest
}

type ListEnterpriseAppRoleEligibilityScheduleRequestsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UnifiedRoleEligibilityScheduleRequest
}

type ListEnterpriseAppRoleEligibilityScheduleRequestsOperationOptions struct {
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

func DefaultListEnterpriseAppRoleEligibilityScheduleRequestsOperationOptions() ListEnterpriseAppRoleEligibilityScheduleRequestsOperationOptions {
	return ListEnterpriseAppRoleEligibilityScheduleRequestsOperationOptions{}
}

func (o ListEnterpriseAppRoleEligibilityScheduleRequestsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEnterpriseAppRoleEligibilityScheduleRequestsOperationOptions) ToOData() *odata.Query {
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

func (o ListEnterpriseAppRoleEligibilityScheduleRequestsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEnterpriseAppRoleEligibilityScheduleRequestsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEnterpriseAppRoleEligibilityScheduleRequestsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEnterpriseAppRoleEligibilityScheduleRequests - Get roleEligibilityScheduleRequests from roleManagement
func (c EnterpriseAppRoleEligibilityScheduleRequestClient) ListEnterpriseAppRoleEligibilityScheduleRequests(ctx context.Context, id beta.RoleManagementEnterpriseAppId, options ListEnterpriseAppRoleEligibilityScheduleRequestsOperationOptions) (result ListEnterpriseAppRoleEligibilityScheduleRequestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEnterpriseAppRoleEligibilityScheduleRequestsCustomPager{},
		Path:          fmt.Sprintf("%s/roleEligibilityScheduleRequests", id.ID()),
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
		Values *[]beta.UnifiedRoleEligibilityScheduleRequest `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEnterpriseAppRoleEligibilityScheduleRequestsComplete retrieves all the results into a single object
func (c EnterpriseAppRoleEligibilityScheduleRequestClient) ListEnterpriseAppRoleEligibilityScheduleRequestsComplete(ctx context.Context, id beta.RoleManagementEnterpriseAppId, options ListEnterpriseAppRoleEligibilityScheduleRequestsOperationOptions) (ListEnterpriseAppRoleEligibilityScheduleRequestsCompleteResult, error) {
	return c.ListEnterpriseAppRoleEligibilityScheduleRequestsCompleteMatchingPredicate(ctx, id, options, UnifiedRoleEligibilityScheduleRequestOperationPredicate{})
}

// ListEnterpriseAppRoleEligibilityScheduleRequestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EnterpriseAppRoleEligibilityScheduleRequestClient) ListEnterpriseAppRoleEligibilityScheduleRequestsCompleteMatchingPredicate(ctx context.Context, id beta.RoleManagementEnterpriseAppId, options ListEnterpriseAppRoleEligibilityScheduleRequestsOperationOptions, predicate UnifiedRoleEligibilityScheduleRequestOperationPredicate) (result ListEnterpriseAppRoleEligibilityScheduleRequestsCompleteResult, err error) {
	items := make([]beta.UnifiedRoleEligibilityScheduleRequest, 0)

	resp, err := c.ListEnterpriseAppRoleEligibilityScheduleRequests(ctx, id, options)
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

	result = ListEnterpriseAppRoleEligibilityScheduleRequestsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
