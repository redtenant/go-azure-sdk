package driveitempermission

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

type ListDriveItemPermissionGrantsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.Permission
}

type ListDriveItemPermissionGrantsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.Permission
}

type ListDriveItemPermissionGrantsOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Skip      *int64
	Top       *int64
}

func DefaultListDriveItemPermissionGrantsOperationOptions() ListDriveItemPermissionGrantsOperationOptions {
	return ListDriveItemPermissionGrantsOperationOptions{}
}

func (o ListDriveItemPermissionGrantsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDriveItemPermissionGrantsOperationOptions) ToOData() *odata.Query {
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

func (o ListDriveItemPermissionGrantsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDriveItemPermissionGrantsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDriveItemPermissionGrantsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDriveItemPermissionGrants - Invoke action grant. Grant users access to a link represented by a permission.
func (c DriveItemPermissionClient) ListDriveItemPermissionGrants(ctx context.Context, id stable.MeDriveIdItemIdPermissionId, input ListDriveItemPermissionGrantsRequest, options ListDriveItemPermissionGrantsOperationOptions) (result ListDriveItemPermissionGrantsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &ListDriveItemPermissionGrantsCustomPager{},
		Path:          fmt.Sprintf("%s/grant", id.ID()),
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
		Values *[]stable.Permission `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDriveItemPermissionGrantsComplete retrieves all the results into a single object
func (c DriveItemPermissionClient) ListDriveItemPermissionGrantsComplete(ctx context.Context, id stable.MeDriveIdItemIdPermissionId, input ListDriveItemPermissionGrantsRequest, options ListDriveItemPermissionGrantsOperationOptions) (ListDriveItemPermissionGrantsCompleteResult, error) {
	return c.ListDriveItemPermissionGrantsCompleteMatchingPredicate(ctx, id, input, options, PermissionOperationPredicate{})
}

// ListDriveItemPermissionGrantsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DriveItemPermissionClient) ListDriveItemPermissionGrantsCompleteMatchingPredicate(ctx context.Context, id stable.MeDriveIdItemIdPermissionId, input ListDriveItemPermissionGrantsRequest, options ListDriveItemPermissionGrantsOperationOptions, predicate PermissionOperationPredicate) (result ListDriveItemPermissionGrantsCompleteResult, err error) {
	items := make([]stable.Permission, 0)

	resp, err := c.ListDriveItemPermissionGrants(ctx, id, input, options)
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

	result = ListDriveItemPermissionGrantsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
