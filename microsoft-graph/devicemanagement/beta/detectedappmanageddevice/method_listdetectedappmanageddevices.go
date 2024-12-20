package detectedappmanageddevice

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListDetectedAppManagedDevicesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ManagedDevice
}

type ListDetectedAppManagedDevicesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ManagedDevice
}

type ListDetectedAppManagedDevicesOperationOptions struct {
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

func DefaultListDetectedAppManagedDevicesOperationOptions() ListDetectedAppManagedDevicesOperationOptions {
	return ListDetectedAppManagedDevicesOperationOptions{}
}

func (o ListDetectedAppManagedDevicesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDetectedAppManagedDevicesOperationOptions) ToOData() *odata.Query {
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

func (o ListDetectedAppManagedDevicesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDetectedAppManagedDevicesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDetectedAppManagedDevicesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDetectedAppManagedDevices - Get managedDevices from deviceManagement. The devices that have the discovered
// application installed
func (c DetectedAppManagedDeviceClient) ListDetectedAppManagedDevices(ctx context.Context, id beta.DeviceManagementDetectedAppId, options ListDetectedAppManagedDevicesOperationOptions) (result ListDetectedAppManagedDevicesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDetectedAppManagedDevicesCustomPager{},
		Path:          fmt.Sprintf("%s/managedDevices", id.ID()),
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
		Values *[]json.RawMessage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	temp := make([]beta.ManagedDevice, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := beta.UnmarshalManagedDeviceImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for beta.ManagedDevice (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListDetectedAppManagedDevicesComplete retrieves all the results into a single object
func (c DetectedAppManagedDeviceClient) ListDetectedAppManagedDevicesComplete(ctx context.Context, id beta.DeviceManagementDetectedAppId, options ListDetectedAppManagedDevicesOperationOptions) (ListDetectedAppManagedDevicesCompleteResult, error) {
	return c.ListDetectedAppManagedDevicesCompleteMatchingPredicate(ctx, id, options, ManagedDeviceOperationPredicate{})
}

// ListDetectedAppManagedDevicesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DetectedAppManagedDeviceClient) ListDetectedAppManagedDevicesCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementDetectedAppId, options ListDetectedAppManagedDevicesOperationOptions, predicate ManagedDeviceOperationPredicate) (result ListDetectedAppManagedDevicesCompleteResult, err error) {
	items := make([]beta.ManagedDevice, 0)

	resp, err := c.ListDetectedAppManagedDevices(ctx, id, options)
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

	result = ListDetectedAppManagedDevicesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
