package credentialuserregistrationdetail

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

type ListCredentialUserRegistrationDetailsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.CredentialUserRegistrationDetails
}

type ListCredentialUserRegistrationDetailsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.CredentialUserRegistrationDetails
}

type ListCredentialUserRegistrationDetailsOperationOptions struct {
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

func DefaultListCredentialUserRegistrationDetailsOperationOptions() ListCredentialUserRegistrationDetailsOperationOptions {
	return ListCredentialUserRegistrationDetailsOperationOptions{}
}

func (o ListCredentialUserRegistrationDetailsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListCredentialUserRegistrationDetailsOperationOptions) ToOData() *odata.Query {
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

func (o ListCredentialUserRegistrationDetailsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListCredentialUserRegistrationDetailsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCredentialUserRegistrationDetailsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCredentialUserRegistrationDetails - List credentialUserRegistrationDetails. Get a list of
// credentialUserRegistrationDetails objects for a given tenant.
func (c CredentialUserRegistrationDetailClient) ListCredentialUserRegistrationDetails(ctx context.Context, options ListCredentialUserRegistrationDetailsOperationOptions) (result ListCredentialUserRegistrationDetailsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListCredentialUserRegistrationDetailsCustomPager{},
		Path:          "/reports/credentialUserRegistrationDetails",
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
		Values *[]beta.CredentialUserRegistrationDetails `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListCredentialUserRegistrationDetailsComplete retrieves all the results into a single object
func (c CredentialUserRegistrationDetailClient) ListCredentialUserRegistrationDetailsComplete(ctx context.Context, options ListCredentialUserRegistrationDetailsOperationOptions) (ListCredentialUserRegistrationDetailsCompleteResult, error) {
	return c.ListCredentialUserRegistrationDetailsCompleteMatchingPredicate(ctx, options, CredentialUserRegistrationDetailsOperationPredicate{})
}

// ListCredentialUserRegistrationDetailsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CredentialUserRegistrationDetailClient) ListCredentialUserRegistrationDetailsCompleteMatchingPredicate(ctx context.Context, options ListCredentialUserRegistrationDetailsOperationOptions, predicate CredentialUserRegistrationDetailsOperationPredicate) (result ListCredentialUserRegistrationDetailsCompleteResult, err error) {
	items := make([]beta.CredentialUserRegistrationDetails, 0)

	resp, err := c.ListCredentialUserRegistrationDetails(ctx, options)
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

	result = ListCredentialUserRegistrationDetailsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
