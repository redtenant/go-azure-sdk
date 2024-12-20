package compliancepolicy

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

type SetCompliancePolicyScheduledActionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementComplianceScheduledActionForRule
}

type SetCompliancePolicyScheduledActionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementComplianceScheduledActionForRule
}

type SetCompliancePolicyScheduledActionsOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Skip      *int64
	Top       *int64
}

func DefaultSetCompliancePolicyScheduledActionsOperationOptions() SetCompliancePolicyScheduledActionsOperationOptions {
	return SetCompliancePolicyScheduledActionsOperationOptions{}
}

func (o SetCompliancePolicyScheduledActionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SetCompliancePolicyScheduledActionsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o SetCompliancePolicyScheduledActionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type SetCompliancePolicyScheduledActionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *SetCompliancePolicyScheduledActionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// SetCompliancePolicyScheduledActions - Invoke action setScheduledActions
func (c CompliancePolicyClient) SetCompliancePolicyScheduledActions(ctx context.Context, id beta.DeviceManagementCompliancePolicyId, input SetCompliancePolicyScheduledActionsRequest, options SetCompliancePolicyScheduledActionsOperationOptions) (result SetCompliancePolicyScheduledActionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &SetCompliancePolicyScheduledActionsCustomPager{},
		Path:          fmt.Sprintf("%s/setScheduledActions", id.ID()),
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
		Values *[]beta.DeviceManagementComplianceScheduledActionForRule `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// SetCompliancePolicyScheduledActionsComplete retrieves all the results into a single object
func (c CompliancePolicyClient) SetCompliancePolicyScheduledActionsComplete(ctx context.Context, id beta.DeviceManagementCompliancePolicyId, input SetCompliancePolicyScheduledActionsRequest, options SetCompliancePolicyScheduledActionsOperationOptions) (SetCompliancePolicyScheduledActionsCompleteResult, error) {
	return c.SetCompliancePolicyScheduledActionsCompleteMatchingPredicate(ctx, id, input, options, DeviceManagementComplianceScheduledActionForRuleOperationPredicate{})
}

// SetCompliancePolicyScheduledActionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CompliancePolicyClient) SetCompliancePolicyScheduledActionsCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementCompliancePolicyId, input SetCompliancePolicyScheduledActionsRequest, options SetCompliancePolicyScheduledActionsOperationOptions, predicate DeviceManagementComplianceScheduledActionForRuleOperationPredicate) (result SetCompliancePolicyScheduledActionsCompleteResult, err error) {
	items := make([]beta.DeviceManagementComplianceScheduledActionForRule, 0)

	resp, err := c.SetCompliancePolicyScheduledActions(ctx, id, input, options)
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

	result = SetCompliancePolicyScheduledActionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
