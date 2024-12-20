package manageddevicelogcollectionrequest

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

type ListManagedDeviceLogCollectionRequestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.DeviceLogCollectionResponse
}

type ListManagedDeviceLogCollectionRequestsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.DeviceLogCollectionResponse
}

type ListManagedDeviceLogCollectionRequestsOperationOptions struct {
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

func DefaultListManagedDeviceLogCollectionRequestsOperationOptions() ListManagedDeviceLogCollectionRequestsOperationOptions {
	return ListManagedDeviceLogCollectionRequestsOperationOptions{}
}

func (o ListManagedDeviceLogCollectionRequestsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListManagedDeviceLogCollectionRequestsOperationOptions) ToOData() *odata.Query {
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

func (o ListManagedDeviceLogCollectionRequestsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListManagedDeviceLogCollectionRequestsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListManagedDeviceLogCollectionRequestsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListManagedDeviceLogCollectionRequests - Get logCollectionRequests from users. List of log collection requests
func (c ManagedDeviceLogCollectionRequestClient) ListManagedDeviceLogCollectionRequests(ctx context.Context, id stable.UserIdManagedDeviceId, options ListManagedDeviceLogCollectionRequestsOperationOptions) (result ListManagedDeviceLogCollectionRequestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListManagedDeviceLogCollectionRequestsCustomPager{},
		Path:          fmt.Sprintf("%s/logCollectionRequests", id.ID()),
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
		Values *[]stable.DeviceLogCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListManagedDeviceLogCollectionRequestsComplete retrieves all the results into a single object
func (c ManagedDeviceLogCollectionRequestClient) ListManagedDeviceLogCollectionRequestsComplete(ctx context.Context, id stable.UserIdManagedDeviceId, options ListManagedDeviceLogCollectionRequestsOperationOptions) (ListManagedDeviceLogCollectionRequestsCompleteResult, error) {
	return c.ListManagedDeviceLogCollectionRequestsCompleteMatchingPredicate(ctx, id, options, DeviceLogCollectionResponseOperationPredicate{})
}

// ListManagedDeviceLogCollectionRequestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ManagedDeviceLogCollectionRequestClient) ListManagedDeviceLogCollectionRequestsCompleteMatchingPredicate(ctx context.Context, id stable.UserIdManagedDeviceId, options ListManagedDeviceLogCollectionRequestsOperationOptions, predicate DeviceLogCollectionResponseOperationPredicate) (result ListManagedDeviceLogCollectionRequestsCompleteResult, err error) {
	items := make([]stable.DeviceLogCollectionResponse, 0)

	resp, err := c.ListManagedDeviceLogCollectionRequests(ctx, id, options)
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

	result = ListManagedDeviceLogCollectionRequestsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
