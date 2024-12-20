package cloudpc

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

type ListCloudPCBulkResizesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.CloudPCRemoteActionResult
}

type ListCloudPCBulkResizesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.CloudPCRemoteActionResult
}

type ListCloudPCBulkResizesOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Skip      *int64
	Top       *int64
}

func DefaultListCloudPCBulkResizesOperationOptions() ListCloudPCBulkResizesOperationOptions {
	return ListCloudPCBulkResizesOperationOptions{}
}

func (o ListCloudPCBulkResizesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListCloudPCBulkResizesOperationOptions) ToOData() *odata.Query {
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

func (o ListCloudPCBulkResizesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListCloudPCBulkResizesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCloudPCBulkResizesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCloudPCBulkResizes - Invoke action bulkResize. Perform a bulk resize action to resize a group of cloudPCs that
// successfully pass validation. If any devices can't be resized, those devices indicate 'resize failed'. The remaining
// devices are provisioned for the resize process.
func (c CloudPCClient) ListCloudPCBulkResizes(ctx context.Context, input ListCloudPCBulkResizesRequest, options ListCloudPCBulkResizesOperationOptions) (result ListCloudPCBulkResizesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &ListCloudPCBulkResizesCustomPager{},
		Path:          "/me/cloudPCs/bulkResize",
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
		Values *[]beta.CloudPCRemoteActionResult `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListCloudPCBulkResizesComplete retrieves all the results into a single object
func (c CloudPCClient) ListCloudPCBulkResizesComplete(ctx context.Context, input ListCloudPCBulkResizesRequest, options ListCloudPCBulkResizesOperationOptions) (ListCloudPCBulkResizesCompleteResult, error) {
	return c.ListCloudPCBulkResizesCompleteMatchingPredicate(ctx, input, options, CloudPCRemoteActionResultOperationPredicate{})
}

// ListCloudPCBulkResizesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CloudPCClient) ListCloudPCBulkResizesCompleteMatchingPredicate(ctx context.Context, input ListCloudPCBulkResizesRequest, options ListCloudPCBulkResizesOperationOptions, predicate CloudPCRemoteActionResultOperationPredicate) (result ListCloudPCBulkResizesCompleteResult, err error) {
	items := make([]beta.CloudPCRemoteActionResult, 0)

	resp, err := c.ListCloudPCBulkResizes(ctx, input, options)
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

	result = ListCloudPCBulkResizesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
