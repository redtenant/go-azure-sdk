package drivelistpermission

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

type ListDriveListPermissionGrantsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Permission
}

type ListDriveListPermissionGrantsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Permission
}

type ListDriveListPermissionGrantsOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Skip      *int64
	Top       *int64
}

func DefaultListDriveListPermissionGrantsOperationOptions() ListDriveListPermissionGrantsOperationOptions {
	return ListDriveListPermissionGrantsOperationOptions{}
}

func (o ListDriveListPermissionGrantsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDriveListPermissionGrantsOperationOptions) ToOData() *odata.Query {
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

func (o ListDriveListPermissionGrantsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDriveListPermissionGrantsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDriveListPermissionGrantsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDriveListPermissionGrants - Invoke action grant. Grant users access to a link represented by a permission.
func (c DriveListPermissionClient) ListDriveListPermissionGrants(ctx context.Context, id beta.GroupIdDriveIdListPermissionId, input ListDriveListPermissionGrantsRequest, options ListDriveListPermissionGrantsOperationOptions) (result ListDriveListPermissionGrantsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &ListDriveListPermissionGrantsCustomPager{},
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
		Values *[]beta.Permission `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDriveListPermissionGrantsComplete retrieves all the results into a single object
func (c DriveListPermissionClient) ListDriveListPermissionGrantsComplete(ctx context.Context, id beta.GroupIdDriveIdListPermissionId, input ListDriveListPermissionGrantsRequest, options ListDriveListPermissionGrantsOperationOptions) (ListDriveListPermissionGrantsCompleteResult, error) {
	return c.ListDriveListPermissionGrantsCompleteMatchingPredicate(ctx, id, input, options, PermissionOperationPredicate{})
}

// ListDriveListPermissionGrantsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DriveListPermissionClient) ListDriveListPermissionGrantsCompleteMatchingPredicate(ctx context.Context, id beta.GroupIdDriveIdListPermissionId, input ListDriveListPermissionGrantsRequest, options ListDriveListPermissionGrantsOperationOptions, predicate PermissionOperationPredicate) (result ListDriveListPermissionGrantsCompleteResult, err error) {
	items := make([]beta.Permission, 0)

	resp, err := c.ListDriveListPermissionGrants(ctx, id, input, options)
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

	result = ListDriveListPermissionGrantsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
