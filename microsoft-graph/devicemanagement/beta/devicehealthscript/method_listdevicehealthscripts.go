package devicehealthscript

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

type ListDeviceHealthScriptsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceHealthScript
}

type ListDeviceHealthScriptsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceHealthScript
}

type ListDeviceHealthScriptsOperationOptions struct {
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

func DefaultListDeviceHealthScriptsOperationOptions() ListDeviceHealthScriptsOperationOptions {
	return ListDeviceHealthScriptsOperationOptions{}
}

func (o ListDeviceHealthScriptsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDeviceHealthScriptsOperationOptions) ToOData() *odata.Query {
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

func (o ListDeviceHealthScriptsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDeviceHealthScriptsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceHealthScriptsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceHealthScripts - Get deviceHealthScripts from deviceManagement. The list of device health scripts associated
// with the tenant.
func (c DeviceHealthScriptClient) ListDeviceHealthScripts(ctx context.Context, options ListDeviceHealthScriptsOperationOptions) (result ListDeviceHealthScriptsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDeviceHealthScriptsCustomPager{},
		Path:          "/deviceManagement/deviceHealthScripts",
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
		Values *[]beta.DeviceHealthScript `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDeviceHealthScriptsComplete retrieves all the results into a single object
func (c DeviceHealthScriptClient) ListDeviceHealthScriptsComplete(ctx context.Context, options ListDeviceHealthScriptsOperationOptions) (ListDeviceHealthScriptsCompleteResult, error) {
	return c.ListDeviceHealthScriptsCompleteMatchingPredicate(ctx, options, DeviceHealthScriptOperationPredicate{})
}

// ListDeviceHealthScriptsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceHealthScriptClient) ListDeviceHealthScriptsCompleteMatchingPredicate(ctx context.Context, options ListDeviceHealthScriptsOperationOptions, predicate DeviceHealthScriptOperationPredicate) (result ListDeviceHealthScriptsCompleteResult, err error) {
	items := make([]beta.DeviceHealthScript, 0)

	resp, err := c.ListDeviceHealthScripts(ctx, options)
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

	result = ListDeviceHealthScriptsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
