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

type ListCloudPCSOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.CloudPC
}

type ListCloudPCSCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.CloudPC
}

type ListCloudPCSOperationOptions struct {
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

func DefaultListCloudPCSOperationOptions() ListCloudPCSOperationOptions {
	return ListCloudPCSOperationOptions{}
}

func (o ListCloudPCSOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListCloudPCSOperationOptions) ToOData() *odata.Query {
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

func (o ListCloudPCSOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListCloudPCSCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCloudPCSCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCloudPCS - List cloudPCs for user. List the cloudPC devices that are attributed to the signed-in user.
func (c CloudPCClient) ListCloudPCS(ctx context.Context, options ListCloudPCSOperationOptions) (result ListCloudPCSOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListCloudPCSCustomPager{},
		Path:          "/me/cloudPCs",
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
		Values *[]beta.CloudPC `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListCloudPCSComplete retrieves all the results into a single object
func (c CloudPCClient) ListCloudPCSComplete(ctx context.Context, options ListCloudPCSOperationOptions) (ListCloudPCSCompleteResult, error) {
	return c.ListCloudPCSCompleteMatchingPredicate(ctx, options, CloudPCOperationPredicate{})
}

// ListCloudPCSCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CloudPCClient) ListCloudPCSCompleteMatchingPredicate(ctx context.Context, options ListCloudPCSOperationOptions, predicate CloudPCOperationPredicate) (result ListCloudPCSCompleteResult, err error) {
	items := make([]beta.CloudPC, 0)

	resp, err := c.ListCloudPCS(ctx, options)
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

	result = ListCloudPCSCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
