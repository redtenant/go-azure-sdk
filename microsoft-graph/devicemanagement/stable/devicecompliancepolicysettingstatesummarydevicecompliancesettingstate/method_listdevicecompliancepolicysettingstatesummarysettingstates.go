package devicecompliancepolicysettingstatesummarydevicecompliancesettingstate

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

type ListDeviceCompliancePolicySettingStateSummarySettingStatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.DeviceComplianceSettingState
}

type ListDeviceCompliancePolicySettingStateSummarySettingStatesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.DeviceComplianceSettingState
}

type ListDeviceCompliancePolicySettingStateSummarySettingStatesOperationOptions struct {
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

func DefaultListDeviceCompliancePolicySettingStateSummarySettingStatesOperationOptions() ListDeviceCompliancePolicySettingStateSummarySettingStatesOperationOptions {
	return ListDeviceCompliancePolicySettingStateSummarySettingStatesOperationOptions{}
}

func (o ListDeviceCompliancePolicySettingStateSummarySettingStatesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDeviceCompliancePolicySettingStateSummarySettingStatesOperationOptions) ToOData() *odata.Query {
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

func (o ListDeviceCompliancePolicySettingStateSummarySettingStatesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDeviceCompliancePolicySettingStateSummarySettingStatesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceCompliancePolicySettingStateSummarySettingStatesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceCompliancePolicySettingStateSummarySettingStates - List deviceComplianceSettingStates. List properties and
// relationships of the deviceComplianceSettingState objects.
func (c DeviceCompliancePolicySettingStateSummaryDeviceComplianceSettingStateClient) ListDeviceCompliancePolicySettingStateSummarySettingStates(ctx context.Context, id stable.DeviceManagementDeviceCompliancePolicySettingStateSummaryId, options ListDeviceCompliancePolicySettingStateSummarySettingStatesOperationOptions) (result ListDeviceCompliancePolicySettingStateSummarySettingStatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDeviceCompliancePolicySettingStateSummarySettingStatesCustomPager{},
		Path:          fmt.Sprintf("%s/deviceComplianceSettingStates", id.ID()),
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
		Values *[]stable.DeviceComplianceSettingState `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDeviceCompliancePolicySettingStateSummarySettingStatesComplete retrieves all the results into a single object
func (c DeviceCompliancePolicySettingStateSummaryDeviceComplianceSettingStateClient) ListDeviceCompliancePolicySettingStateSummarySettingStatesComplete(ctx context.Context, id stable.DeviceManagementDeviceCompliancePolicySettingStateSummaryId, options ListDeviceCompliancePolicySettingStateSummarySettingStatesOperationOptions) (ListDeviceCompliancePolicySettingStateSummarySettingStatesCompleteResult, error) {
	return c.ListDeviceCompliancePolicySettingStateSummarySettingStatesCompleteMatchingPredicate(ctx, id, options, DeviceComplianceSettingStateOperationPredicate{})
}

// ListDeviceCompliancePolicySettingStateSummarySettingStatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceCompliancePolicySettingStateSummaryDeviceComplianceSettingStateClient) ListDeviceCompliancePolicySettingStateSummarySettingStatesCompleteMatchingPredicate(ctx context.Context, id stable.DeviceManagementDeviceCompliancePolicySettingStateSummaryId, options ListDeviceCompliancePolicySettingStateSummarySettingStatesOperationOptions, predicate DeviceComplianceSettingStateOperationPredicate) (result ListDeviceCompliancePolicySettingStateSummarySettingStatesCompleteResult, err error) {
	items := make([]stable.DeviceComplianceSettingState, 0)

	resp, err := c.ListDeviceCompliancePolicySettingStateSummarySettingStates(ctx, id, options)
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

	result = ListDeviceCompliancePolicySettingStateSummarySettingStatesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
