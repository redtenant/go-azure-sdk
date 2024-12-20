package comanageddevicedevicecompliancepolicystate

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

type ListComanagedDeviceCompliancePolicyStatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceCompliancePolicyState
}

type ListComanagedDeviceCompliancePolicyStatesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceCompliancePolicyState
}

type ListComanagedDeviceCompliancePolicyStatesOperationOptions struct {
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

func DefaultListComanagedDeviceCompliancePolicyStatesOperationOptions() ListComanagedDeviceCompliancePolicyStatesOperationOptions {
	return ListComanagedDeviceCompliancePolicyStatesOperationOptions{}
}

func (o ListComanagedDeviceCompliancePolicyStatesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListComanagedDeviceCompliancePolicyStatesOperationOptions) ToOData() *odata.Query {
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

func (o ListComanagedDeviceCompliancePolicyStatesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListComanagedDeviceCompliancePolicyStatesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListComanagedDeviceCompliancePolicyStatesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListComanagedDeviceCompliancePolicyStates - Get deviceCompliancePolicyStates from deviceManagement. Device compliance
// policy states for this device.
func (c ComanagedDeviceDeviceCompliancePolicyStateClient) ListComanagedDeviceCompliancePolicyStates(ctx context.Context, id beta.DeviceManagementComanagedDeviceId, options ListComanagedDeviceCompliancePolicyStatesOperationOptions) (result ListComanagedDeviceCompliancePolicyStatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListComanagedDeviceCompliancePolicyStatesCustomPager{},
		Path:          fmt.Sprintf("%s/deviceCompliancePolicyStates", id.ID()),
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
		Values *[]beta.DeviceCompliancePolicyState `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListComanagedDeviceCompliancePolicyStatesComplete retrieves all the results into a single object
func (c ComanagedDeviceDeviceCompliancePolicyStateClient) ListComanagedDeviceCompliancePolicyStatesComplete(ctx context.Context, id beta.DeviceManagementComanagedDeviceId, options ListComanagedDeviceCompliancePolicyStatesOperationOptions) (ListComanagedDeviceCompliancePolicyStatesCompleteResult, error) {
	return c.ListComanagedDeviceCompliancePolicyStatesCompleteMatchingPredicate(ctx, id, options, DeviceCompliancePolicyStateOperationPredicate{})
}

// ListComanagedDeviceCompliancePolicyStatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ComanagedDeviceDeviceCompliancePolicyStateClient) ListComanagedDeviceCompliancePolicyStatesCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementComanagedDeviceId, options ListComanagedDeviceCompliancePolicyStatesOperationOptions, predicate DeviceCompliancePolicyStateOperationPredicate) (result ListComanagedDeviceCompliancePolicyStatesCompleteResult, err error) {
	items := make([]beta.DeviceCompliancePolicyState, 0)

	resp, err := c.ListComanagedDeviceCompliancePolicyStates(ctx, id, options)
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

	result = ListComanagedDeviceCompliancePolicyStatesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
