package licensedetail

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

type ListLicenseDetailsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.LicenseDetails
}

type ListLicenseDetailsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.LicenseDetails
}

type ListLicenseDetailsOperationOptions struct {
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

func DefaultListLicenseDetailsOperationOptions() ListLicenseDetailsOperationOptions {
	return ListLicenseDetailsOperationOptions{}
}

func (o ListLicenseDetailsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListLicenseDetailsOperationOptions) ToOData() *odata.Query {
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

func (o ListLicenseDetailsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListLicenseDetailsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListLicenseDetailsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListLicenseDetails - Get licenseDetails from users. A collection of this user's license details. Read-only.
func (c LicenseDetailClient) ListLicenseDetails(ctx context.Context, id stable.UserId, options ListLicenseDetailsOperationOptions) (result ListLicenseDetailsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListLicenseDetailsCustomPager{},
		Path:          fmt.Sprintf("%s/licenseDetails", id.ID()),
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
		Values *[]stable.LicenseDetails `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListLicenseDetailsComplete retrieves all the results into a single object
func (c LicenseDetailClient) ListLicenseDetailsComplete(ctx context.Context, id stable.UserId, options ListLicenseDetailsOperationOptions) (ListLicenseDetailsCompleteResult, error) {
	return c.ListLicenseDetailsCompleteMatchingPredicate(ctx, id, options, LicenseDetailsOperationPredicate{})
}

// ListLicenseDetailsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c LicenseDetailClient) ListLicenseDetailsCompleteMatchingPredicate(ctx context.Context, id stable.UserId, options ListLicenseDetailsOperationOptions, predicate LicenseDetailsOperationPredicate) (result ListLicenseDetailsCompleteResult, err error) {
	items := make([]stable.LicenseDetails, 0)

	resp, err := c.ListLicenseDetails(ctx, id, options)
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

	result = ListLicenseDetailsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
