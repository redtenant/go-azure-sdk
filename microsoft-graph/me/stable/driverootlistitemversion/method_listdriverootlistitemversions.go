package driverootlistitemversion

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListDriveRootListItemVersionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ListItemVersion
}

type ListDriveRootListItemVersionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ListItemVersion
}

type ListDriveRootListItemVersionsOperationOptions struct {
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

func DefaultListDriveRootListItemVersionsOperationOptions() ListDriveRootListItemVersionsOperationOptions {
	return ListDriveRootListItemVersionsOperationOptions{}
}

func (o ListDriveRootListItemVersionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDriveRootListItemVersionsOperationOptions) ToOData() *odata.Query {
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

func (o ListDriveRootListItemVersionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDriveRootListItemVersionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDriveRootListItemVersionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDriveRootListItemVersions - Get versions from me. The list of previous versions of the list item.
func (c DriveRootListItemVersionClient) ListDriveRootListItemVersions(ctx context.Context, id stable.MeDriveId, options ListDriveRootListItemVersionsOperationOptions) (result ListDriveRootListItemVersionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDriveRootListItemVersionsCustomPager{},
		Path:          fmt.Sprintf("%s/root/listItem/versions", id.ID()),
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

	temp := make([]stable.ListItemVersion, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := stable.UnmarshalListItemVersionImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for stable.ListItemVersion (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListDriveRootListItemVersionsComplete retrieves all the results into a single object
func (c DriveRootListItemVersionClient) ListDriveRootListItemVersionsComplete(ctx context.Context, id stable.MeDriveId, options ListDriveRootListItemVersionsOperationOptions) (ListDriveRootListItemVersionsCompleteResult, error) {
	return c.ListDriveRootListItemVersionsCompleteMatchingPredicate(ctx, id, options, ListItemVersionOperationPredicate{})
}

// ListDriveRootListItemVersionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DriveRootListItemVersionClient) ListDriveRootListItemVersionsCompleteMatchingPredicate(ctx context.Context, id stable.MeDriveId, options ListDriveRootListItemVersionsOperationOptions, predicate ListItemVersionOperationPredicate) (result ListDriveRootListItemVersionsCompleteResult, err error) {
	items := make([]stable.ListItemVersion, 0)

	resp, err := c.ListDriveRootListItemVersions(ctx, id, options)
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

	result = ListDriveRootListItemVersionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
