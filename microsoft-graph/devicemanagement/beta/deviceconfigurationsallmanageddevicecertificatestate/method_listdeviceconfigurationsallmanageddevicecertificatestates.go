package deviceconfigurationsallmanageddevicecertificatestate

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

type ListDeviceConfigurationsAllManagedDeviceCertificateStatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ManagedAllDeviceCertificateState
}

type ListDeviceConfigurationsAllManagedDeviceCertificateStatesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ManagedAllDeviceCertificateState
}

type ListDeviceConfigurationsAllManagedDeviceCertificateStatesOperationOptions struct {
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

func DefaultListDeviceConfigurationsAllManagedDeviceCertificateStatesOperationOptions() ListDeviceConfigurationsAllManagedDeviceCertificateStatesOperationOptions {
	return ListDeviceConfigurationsAllManagedDeviceCertificateStatesOperationOptions{}
}

func (o ListDeviceConfigurationsAllManagedDeviceCertificateStatesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDeviceConfigurationsAllManagedDeviceCertificateStatesOperationOptions) ToOData() *odata.Query {
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

func (o ListDeviceConfigurationsAllManagedDeviceCertificateStatesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDeviceConfigurationsAllManagedDeviceCertificateStatesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceConfigurationsAllManagedDeviceCertificateStatesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceConfigurationsAllManagedDeviceCertificateStates - Get deviceConfigurationsAllManagedDeviceCertificateStates
// from deviceManagement. Summary of all certificates for all devices.
func (c DeviceConfigurationsAllManagedDeviceCertificateStateClient) ListDeviceConfigurationsAllManagedDeviceCertificateStates(ctx context.Context, options ListDeviceConfigurationsAllManagedDeviceCertificateStatesOperationOptions) (result ListDeviceConfigurationsAllManagedDeviceCertificateStatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDeviceConfigurationsAllManagedDeviceCertificateStatesCustomPager{},
		Path:          "/deviceManagement/deviceConfigurationsAllManagedDeviceCertificateStates",
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
		Values *[]beta.ManagedAllDeviceCertificateState `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDeviceConfigurationsAllManagedDeviceCertificateStatesComplete retrieves all the results into a single object
func (c DeviceConfigurationsAllManagedDeviceCertificateStateClient) ListDeviceConfigurationsAllManagedDeviceCertificateStatesComplete(ctx context.Context, options ListDeviceConfigurationsAllManagedDeviceCertificateStatesOperationOptions) (ListDeviceConfigurationsAllManagedDeviceCertificateStatesCompleteResult, error) {
	return c.ListDeviceConfigurationsAllManagedDeviceCertificateStatesCompleteMatchingPredicate(ctx, options, ManagedAllDeviceCertificateStateOperationPredicate{})
}

// ListDeviceConfigurationsAllManagedDeviceCertificateStatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceConfigurationsAllManagedDeviceCertificateStateClient) ListDeviceConfigurationsAllManagedDeviceCertificateStatesCompleteMatchingPredicate(ctx context.Context, options ListDeviceConfigurationsAllManagedDeviceCertificateStatesOperationOptions, predicate ManagedAllDeviceCertificateStateOperationPredicate) (result ListDeviceConfigurationsAllManagedDeviceCertificateStatesCompleteResult, err error) {
	items := make([]beta.ManagedAllDeviceCertificateState, 0)

	resp, err := c.ListDeviceConfigurationsAllManagedDeviceCertificateStates(ctx, options)
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

	result = ListDeviceConfigurationsAllManagedDeviceCertificateStatesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
