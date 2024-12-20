package devicecustomattributeshellscriptassignment

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

type ListDeviceCustomAttributeShellScriptAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementScriptAssignment
}

type ListDeviceCustomAttributeShellScriptAssignmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementScriptAssignment
}

type ListDeviceCustomAttributeShellScriptAssignmentsOperationOptions struct {
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

func DefaultListDeviceCustomAttributeShellScriptAssignmentsOperationOptions() ListDeviceCustomAttributeShellScriptAssignmentsOperationOptions {
	return ListDeviceCustomAttributeShellScriptAssignmentsOperationOptions{}
}

func (o ListDeviceCustomAttributeShellScriptAssignmentsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDeviceCustomAttributeShellScriptAssignmentsOperationOptions) ToOData() *odata.Query {
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

func (o ListDeviceCustomAttributeShellScriptAssignmentsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDeviceCustomAttributeShellScriptAssignmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceCustomAttributeShellScriptAssignmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceCustomAttributeShellScriptAssignments - Get assignments from deviceManagement. The list of group
// assignments for the device management script.
func (c DeviceCustomAttributeShellScriptAssignmentClient) ListDeviceCustomAttributeShellScriptAssignments(ctx context.Context, id beta.DeviceManagementDeviceCustomAttributeShellScriptId, options ListDeviceCustomAttributeShellScriptAssignmentsOperationOptions) (result ListDeviceCustomAttributeShellScriptAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDeviceCustomAttributeShellScriptAssignmentsCustomPager{},
		Path:          fmt.Sprintf("%s/assignments", id.ID()),
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
		Values *[]beta.DeviceManagementScriptAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDeviceCustomAttributeShellScriptAssignmentsComplete retrieves all the results into a single object
func (c DeviceCustomAttributeShellScriptAssignmentClient) ListDeviceCustomAttributeShellScriptAssignmentsComplete(ctx context.Context, id beta.DeviceManagementDeviceCustomAttributeShellScriptId, options ListDeviceCustomAttributeShellScriptAssignmentsOperationOptions) (ListDeviceCustomAttributeShellScriptAssignmentsCompleteResult, error) {
	return c.ListDeviceCustomAttributeShellScriptAssignmentsCompleteMatchingPredicate(ctx, id, options, DeviceManagementScriptAssignmentOperationPredicate{})
}

// ListDeviceCustomAttributeShellScriptAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceCustomAttributeShellScriptAssignmentClient) ListDeviceCustomAttributeShellScriptAssignmentsCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementDeviceCustomAttributeShellScriptId, options ListDeviceCustomAttributeShellScriptAssignmentsOperationOptions, predicate DeviceManagementScriptAssignmentOperationPredicate) (result ListDeviceCustomAttributeShellScriptAssignmentsCompleteResult, err error) {
	items := make([]beta.DeviceManagementScriptAssignment, 0)

	resp, err := c.ListDeviceCustomAttributeShellScriptAssignments(ctx, id, options)
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

	result = ListDeviceCustomAttributeShellScriptAssignmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
