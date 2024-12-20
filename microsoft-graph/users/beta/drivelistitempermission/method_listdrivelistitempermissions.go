package drivelistitempermission

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

type ListDriveListItemPermissionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Permission
}

type ListDriveListItemPermissionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Permission
}

type ListDriveListItemPermissionsOperationOptions struct {
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

func DefaultListDriveListItemPermissionsOperationOptions() ListDriveListItemPermissionsOperationOptions {
	return ListDriveListItemPermissionsOperationOptions{}
}

func (o ListDriveListItemPermissionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDriveListItemPermissionsOperationOptions) ToOData() *odata.Query {
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

func (o ListDriveListItemPermissionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDriveListItemPermissionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDriveListItemPermissionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDriveListItemPermissions - Get permissions from users. The set of permissions for the item. Read-only. Nullable.
func (c DriveListItemPermissionClient) ListDriveListItemPermissions(ctx context.Context, id beta.UserIdDriveIdListItemId, options ListDriveListItemPermissionsOperationOptions) (result ListDriveListItemPermissionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDriveListItemPermissionsCustomPager{},
		Path:          fmt.Sprintf("%s/permissions", id.ID()),
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

// ListDriveListItemPermissionsComplete retrieves all the results into a single object
func (c DriveListItemPermissionClient) ListDriveListItemPermissionsComplete(ctx context.Context, id beta.UserIdDriveIdListItemId, options ListDriveListItemPermissionsOperationOptions) (ListDriveListItemPermissionsCompleteResult, error) {
	return c.ListDriveListItemPermissionsCompleteMatchingPredicate(ctx, id, options, PermissionOperationPredicate{})
}

// ListDriveListItemPermissionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DriveListItemPermissionClient) ListDriveListItemPermissionsCompleteMatchingPredicate(ctx context.Context, id beta.UserIdDriveIdListItemId, options ListDriveListItemPermissionsOperationOptions, predicate PermissionOperationPredicate) (result ListDriveListItemPermissionsCompleteResult, err error) {
	items := make([]beta.Permission, 0)

	resp, err := c.ListDriveListItemPermissions(ctx, id, options)
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

	result = ListDriveListItemPermissionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
