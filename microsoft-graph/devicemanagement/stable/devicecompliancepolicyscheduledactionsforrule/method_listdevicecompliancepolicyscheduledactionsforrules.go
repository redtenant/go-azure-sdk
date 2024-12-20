package devicecompliancepolicyscheduledactionsforrule

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListDeviceCompliancePolicyScheduledActionsForRulesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.DeviceComplianceScheduledActionForRule
}

type ListDeviceCompliancePolicyScheduledActionsForRulesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.DeviceComplianceScheduledActionForRule
}

type ListDeviceCompliancePolicyScheduledActionsForRulesOperationOptions struct {
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

func DefaultListDeviceCompliancePolicyScheduledActionsForRulesOperationOptions() ListDeviceCompliancePolicyScheduledActionsForRulesOperationOptions {
	return ListDeviceCompliancePolicyScheduledActionsForRulesOperationOptions{}
}

func (o ListDeviceCompliancePolicyScheduledActionsForRulesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDeviceCompliancePolicyScheduledActionsForRulesOperationOptions) ToOData() *odata.Query {
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

func (o ListDeviceCompliancePolicyScheduledActionsForRulesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDeviceCompliancePolicyScheduledActionsForRulesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceCompliancePolicyScheduledActionsForRulesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceCompliancePolicyScheduledActionsForRules - List deviceComplianceScheduledActionForRules. List properties
// and relationships of the deviceComplianceScheduledActionForRule objects.
func (c DeviceCompliancePolicyScheduledActionsForRuleClient) ListDeviceCompliancePolicyScheduledActionsForRules(ctx context.Context, id stable.DeviceManagementDeviceCompliancePolicyId, options ListDeviceCompliancePolicyScheduledActionsForRulesOperationOptions) (result ListDeviceCompliancePolicyScheduledActionsForRulesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDeviceCompliancePolicyScheduledActionsForRulesCustomPager{},
		Path:          fmt.Sprintf("%s/scheduledActionsForRule", id.ID()),
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
		Values *[]stable.DeviceComplianceScheduledActionForRule `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDeviceCompliancePolicyScheduledActionsForRulesComplete retrieves all the results into a single object
func (c DeviceCompliancePolicyScheduledActionsForRuleClient) ListDeviceCompliancePolicyScheduledActionsForRulesComplete(ctx context.Context, id stable.DeviceManagementDeviceCompliancePolicyId, options ListDeviceCompliancePolicyScheduledActionsForRulesOperationOptions) (ListDeviceCompliancePolicyScheduledActionsForRulesCompleteResult, error) {
	return c.ListDeviceCompliancePolicyScheduledActionsForRulesCompleteMatchingPredicate(ctx, id, options, DeviceComplianceScheduledActionForRuleOperationPredicate{})
}

// ListDeviceCompliancePolicyScheduledActionsForRulesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceCompliancePolicyScheduledActionsForRuleClient) ListDeviceCompliancePolicyScheduledActionsForRulesCompleteMatchingPredicate(ctx context.Context, id stable.DeviceManagementDeviceCompliancePolicyId, options ListDeviceCompliancePolicyScheduledActionsForRulesOperationOptions, predicate DeviceComplianceScheduledActionForRuleOperationPredicate) (result ListDeviceCompliancePolicyScheduledActionsForRulesCompleteResult, err error) {
	items := make([]stable.DeviceComplianceScheduledActionForRule, 0)

	resp, err := c.ListDeviceCompliancePolicyScheduledActionsForRules(ctx, id, options)
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

	result = ListDeviceCompliancePolicyScheduledActionsForRulesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
