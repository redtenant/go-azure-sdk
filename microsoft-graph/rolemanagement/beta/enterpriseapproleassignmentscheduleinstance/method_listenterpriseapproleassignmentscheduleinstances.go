package enterpriseapproleassignmentscheduleinstance

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

type ListEnterpriseAppRoleAssignmentScheduleInstancesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UnifiedRoleAssignmentScheduleInstance
}

type ListEnterpriseAppRoleAssignmentScheduleInstancesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UnifiedRoleAssignmentScheduleInstance
}

type ListEnterpriseAppRoleAssignmentScheduleInstancesOperationOptions struct {
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

func DefaultListEnterpriseAppRoleAssignmentScheduleInstancesOperationOptions() ListEnterpriseAppRoleAssignmentScheduleInstancesOperationOptions {
	return ListEnterpriseAppRoleAssignmentScheduleInstancesOperationOptions{}
}

func (o ListEnterpriseAppRoleAssignmentScheduleInstancesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEnterpriseAppRoleAssignmentScheduleInstancesOperationOptions) ToOData() *odata.Query {
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

func (o ListEnterpriseAppRoleAssignmentScheduleInstancesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEnterpriseAppRoleAssignmentScheduleInstancesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEnterpriseAppRoleAssignmentScheduleInstancesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEnterpriseAppRoleAssignmentScheduleInstances - Get roleAssignmentScheduleInstances from roleManagement
func (c EnterpriseAppRoleAssignmentScheduleInstanceClient) ListEnterpriseAppRoleAssignmentScheduleInstances(ctx context.Context, id beta.RoleManagementEnterpriseAppId, options ListEnterpriseAppRoleAssignmentScheduleInstancesOperationOptions) (result ListEnterpriseAppRoleAssignmentScheduleInstancesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEnterpriseAppRoleAssignmentScheduleInstancesCustomPager{},
		Path:          fmt.Sprintf("%s/roleAssignmentScheduleInstances", id.ID()),
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
		Values *[]beta.UnifiedRoleAssignmentScheduleInstance `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEnterpriseAppRoleAssignmentScheduleInstancesComplete retrieves all the results into a single object
func (c EnterpriseAppRoleAssignmentScheduleInstanceClient) ListEnterpriseAppRoleAssignmentScheduleInstancesComplete(ctx context.Context, id beta.RoleManagementEnterpriseAppId, options ListEnterpriseAppRoleAssignmentScheduleInstancesOperationOptions) (ListEnterpriseAppRoleAssignmentScheduleInstancesCompleteResult, error) {
	return c.ListEnterpriseAppRoleAssignmentScheduleInstancesCompleteMatchingPredicate(ctx, id, options, UnifiedRoleAssignmentScheduleInstanceOperationPredicate{})
}

// ListEnterpriseAppRoleAssignmentScheduleInstancesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EnterpriseAppRoleAssignmentScheduleInstanceClient) ListEnterpriseAppRoleAssignmentScheduleInstancesCompleteMatchingPredicate(ctx context.Context, id beta.RoleManagementEnterpriseAppId, options ListEnterpriseAppRoleAssignmentScheduleInstancesOperationOptions, predicate UnifiedRoleAssignmentScheduleInstanceOperationPredicate) (result ListEnterpriseAppRoleAssignmentScheduleInstancesCompleteResult, err error) {
	items := make([]beta.UnifiedRoleAssignmentScheduleInstance, 0)

	resp, err := c.ListEnterpriseAppRoleAssignmentScheduleInstances(ctx, id, options)
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

	result = ListEnterpriseAppRoleAssignmentScheduleInstancesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
