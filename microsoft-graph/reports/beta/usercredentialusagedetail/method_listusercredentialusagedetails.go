package usercredentialusagedetail

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

type ListUserCredentialUsageDetailsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UserCredentialUsageDetails
}

type ListUserCredentialUsageDetailsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UserCredentialUsageDetails
}

type ListUserCredentialUsageDetailsOperationOptions struct {
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

func DefaultListUserCredentialUsageDetailsOperationOptions() ListUserCredentialUsageDetailsOperationOptions {
	return ListUserCredentialUsageDetailsOperationOptions{}
}

func (o ListUserCredentialUsageDetailsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListUserCredentialUsageDetailsOperationOptions) ToOData() *odata.Query {
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

func (o ListUserCredentialUsageDetailsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListUserCredentialUsageDetailsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserCredentialUsageDetailsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserCredentialUsageDetails - List userCredentialUsageDetails. Get a list of userCredentialUsageDetails objects
// for a given tenant. Details include user information, status of the reset, and the reason for failure.
func (c UserCredentialUsageDetailClient) ListUserCredentialUsageDetails(ctx context.Context, options ListUserCredentialUsageDetailsOperationOptions) (result ListUserCredentialUsageDetailsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListUserCredentialUsageDetailsCustomPager{},
		Path:          "/reports/userCredentialUsageDetails",
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
		Values *[]beta.UserCredentialUsageDetails `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserCredentialUsageDetailsComplete retrieves all the results into a single object
func (c UserCredentialUsageDetailClient) ListUserCredentialUsageDetailsComplete(ctx context.Context, options ListUserCredentialUsageDetailsOperationOptions) (ListUserCredentialUsageDetailsCompleteResult, error) {
	return c.ListUserCredentialUsageDetailsCompleteMatchingPredicate(ctx, options, UserCredentialUsageDetailsOperationPredicate{})
}

// ListUserCredentialUsageDetailsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserCredentialUsageDetailClient) ListUserCredentialUsageDetailsCompleteMatchingPredicate(ctx context.Context, options ListUserCredentialUsageDetailsOperationOptions, predicate UserCredentialUsageDetailsOperationPredicate) (result ListUserCredentialUsageDetailsCompleteResult, err error) {
	items := make([]beta.UserCredentialUsageDetails, 0)

	resp, err := c.ListUserCredentialUsageDetails(ctx, options)
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

	result = ListUserCredentialUsageDetailsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
