package deviceregisteredowner

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

type ListDeviceRegisteredOwnerRefsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DirectoryObject
}

type ListDeviceRegisteredOwnerRefsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DirectoryObject
}

type ListDeviceRegisteredOwnerRefsOperationOptions struct {
	ConsistencyLevel *odata.ConsistencyLevel
	Count            *bool
	Filter           *string
	Metadata         *odata.Metadata
	OrderBy          *odata.OrderBy
	RetryFunc        client.RequestRetryFunc
	Search           *string
	Skip             *int64
	Top              *int64
}

func DefaultListDeviceRegisteredOwnerRefsOperationOptions() ListDeviceRegisteredOwnerRefsOperationOptions {
	return ListDeviceRegisteredOwnerRefsOperationOptions{}
}

func (o ListDeviceRegisteredOwnerRefsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDeviceRegisteredOwnerRefsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.ConsistencyLevel != nil {
		out.ConsistencyLevel = *o.ConsistencyLevel
	}
	if o.Count != nil {
		out.Count = *o.Count
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
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListDeviceRegisteredOwnerRefsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDeviceRegisteredOwnerRefsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceRegisteredOwnerRefsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceRegisteredOwnerRefs - Get ref of registeredOwners from me. The user that cloud joined the device or
// registered their personal device. The registered owner is set at the time of registration. Read-only. Nullable.
// Supports $expand.
func (c DeviceRegisteredOwnerClient) ListDeviceRegisteredOwnerRefs(ctx context.Context, id beta.MeDeviceId, options ListDeviceRegisteredOwnerRefsOperationOptions) (result ListDeviceRegisteredOwnerRefsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDeviceRegisteredOwnerRefsCustomPager{},
		Path:          fmt.Sprintf("%s/registeredOwners/$ref", id.ID()),
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

	temp := make([]beta.DirectoryObject, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := beta.UnmarshalDirectoryObjectImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for beta.DirectoryObject (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListDeviceRegisteredOwnerRefsComplete retrieves all the results into a single object
func (c DeviceRegisteredOwnerClient) ListDeviceRegisteredOwnerRefsComplete(ctx context.Context, id beta.MeDeviceId, options ListDeviceRegisteredOwnerRefsOperationOptions) (ListDeviceRegisteredOwnerRefsCompleteResult, error) {
	return c.ListDeviceRegisteredOwnerRefsCompleteMatchingPredicate(ctx, id, options, DirectoryObjectOperationPredicate{})
}

// ListDeviceRegisteredOwnerRefsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceRegisteredOwnerClient) ListDeviceRegisteredOwnerRefsCompleteMatchingPredicate(ctx context.Context, id beta.MeDeviceId, options ListDeviceRegisteredOwnerRefsOperationOptions, predicate DirectoryObjectOperationPredicate) (result ListDeviceRegisteredOwnerRefsCompleteResult, err error) {
	items := make([]beta.DirectoryObject, 0)

	resp, err := c.ListDeviceRegisteredOwnerRefs(ctx, id, options)
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

	result = ListDeviceRegisteredOwnerRefsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
