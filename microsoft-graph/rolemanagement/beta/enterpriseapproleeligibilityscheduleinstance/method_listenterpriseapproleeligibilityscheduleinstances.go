package enterpriseapproleeligibilityscheduleinstance

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

type ListEnterpriseAppRoleEligibilityScheduleInstancesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UnifiedRoleEligibilityScheduleInstance
}

type ListEnterpriseAppRoleEligibilityScheduleInstancesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UnifiedRoleEligibilityScheduleInstance
}

type ListEnterpriseAppRoleEligibilityScheduleInstancesOperationOptions struct {
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

func DefaultListEnterpriseAppRoleEligibilityScheduleInstancesOperationOptions() ListEnterpriseAppRoleEligibilityScheduleInstancesOperationOptions {
	return ListEnterpriseAppRoleEligibilityScheduleInstancesOperationOptions{}
}

func (o ListEnterpriseAppRoleEligibilityScheduleInstancesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEnterpriseAppRoleEligibilityScheduleInstancesOperationOptions) ToOData() *odata.Query {
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

func (o ListEnterpriseAppRoleEligibilityScheduleInstancesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEnterpriseAppRoleEligibilityScheduleInstancesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEnterpriseAppRoleEligibilityScheduleInstancesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEnterpriseAppRoleEligibilityScheduleInstances - Get roleEligibilityScheduleInstances from roleManagement
func (c EnterpriseAppRoleEligibilityScheduleInstanceClient) ListEnterpriseAppRoleEligibilityScheduleInstances(ctx context.Context, id beta.RoleManagementEnterpriseAppId, options ListEnterpriseAppRoleEligibilityScheduleInstancesOperationOptions) (result ListEnterpriseAppRoleEligibilityScheduleInstancesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEnterpriseAppRoleEligibilityScheduleInstancesCustomPager{},
		Path:          fmt.Sprintf("%s/roleEligibilityScheduleInstances", id.ID()),
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
		Values *[]beta.UnifiedRoleEligibilityScheduleInstance `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEnterpriseAppRoleEligibilityScheduleInstancesComplete retrieves all the results into a single object
func (c EnterpriseAppRoleEligibilityScheduleInstanceClient) ListEnterpriseAppRoleEligibilityScheduleInstancesComplete(ctx context.Context, id beta.RoleManagementEnterpriseAppId, options ListEnterpriseAppRoleEligibilityScheduleInstancesOperationOptions) (ListEnterpriseAppRoleEligibilityScheduleInstancesCompleteResult, error) {
	return c.ListEnterpriseAppRoleEligibilityScheduleInstancesCompleteMatchingPredicate(ctx, id, options, UnifiedRoleEligibilityScheduleInstanceOperationPredicate{})
}

// ListEnterpriseAppRoleEligibilityScheduleInstancesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EnterpriseAppRoleEligibilityScheduleInstanceClient) ListEnterpriseAppRoleEligibilityScheduleInstancesCompleteMatchingPredicate(ctx context.Context, id beta.RoleManagementEnterpriseAppId, options ListEnterpriseAppRoleEligibilityScheduleInstancesOperationOptions, predicate UnifiedRoleEligibilityScheduleInstanceOperationPredicate) (result ListEnterpriseAppRoleEligibilityScheduleInstancesCompleteResult, err error) {
	items := make([]beta.UnifiedRoleEligibilityScheduleInstance, 0)

	resp, err := c.ListEnterpriseAppRoleEligibilityScheduleInstances(ctx, id, options)
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

	result = ListEnterpriseAppRoleEligibilityScheduleInstancesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
