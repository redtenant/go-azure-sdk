package delegatedpermissionclassification

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

type ListDelegatedPermissionClassificationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DelegatedPermissionClassification
}

type ListDelegatedPermissionClassificationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DelegatedPermissionClassification
}

type ListDelegatedPermissionClassificationsOperationOptions struct {
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

func DefaultListDelegatedPermissionClassificationsOperationOptions() ListDelegatedPermissionClassificationsOperationOptions {
	return ListDelegatedPermissionClassificationsOperationOptions{}
}

func (o ListDelegatedPermissionClassificationsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDelegatedPermissionClassificationsOperationOptions) ToOData() *odata.Query {
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

func (o ListDelegatedPermissionClassificationsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDelegatedPermissionClassificationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDelegatedPermissionClassificationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDelegatedPermissionClassifications - List delegatedPermissionClassifications collection of servicePrincipal.
// Retrieve the list of delegatedPermissionClassification currently configured for the delegated permissions exposed by
// an API.
func (c DelegatedPermissionClassificationClient) ListDelegatedPermissionClassifications(ctx context.Context, id beta.ServicePrincipalId, options ListDelegatedPermissionClassificationsOperationOptions) (result ListDelegatedPermissionClassificationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDelegatedPermissionClassificationsCustomPager{},
		Path:          fmt.Sprintf("%s/delegatedPermissionClassifications", id.ID()),
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
		Values *[]beta.DelegatedPermissionClassification `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDelegatedPermissionClassificationsComplete retrieves all the results into a single object
func (c DelegatedPermissionClassificationClient) ListDelegatedPermissionClassificationsComplete(ctx context.Context, id beta.ServicePrincipalId, options ListDelegatedPermissionClassificationsOperationOptions) (ListDelegatedPermissionClassificationsCompleteResult, error) {
	return c.ListDelegatedPermissionClassificationsCompleteMatchingPredicate(ctx, id, options, DelegatedPermissionClassificationOperationPredicate{})
}

// ListDelegatedPermissionClassificationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DelegatedPermissionClassificationClient) ListDelegatedPermissionClassificationsCompleteMatchingPredicate(ctx context.Context, id beta.ServicePrincipalId, options ListDelegatedPermissionClassificationsOperationOptions, predicate DelegatedPermissionClassificationOperationPredicate) (result ListDelegatedPermissionClassificationsCompleteResult, err error) {
	items := make([]beta.DelegatedPermissionClassification, 0)

	resp, err := c.ListDelegatedPermissionClassifications(ctx, id, options)
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

	result = ListDelegatedPermissionClassificationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
