package domainnamereference

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

type ListDomainNameReferencesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.DirectoryObject
}

type ListDomainNameReferencesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.DirectoryObject
}

type ListDomainNameReferencesOperationOptions struct {
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

func DefaultListDomainNameReferencesOperationOptions() ListDomainNameReferencesOperationOptions {
	return ListDomainNameReferencesOperationOptions{}
}

func (o ListDomainNameReferencesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDomainNameReferencesOperationOptions) ToOData() *odata.Query {
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

func (o ListDomainNameReferencesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDomainNameReferencesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDomainNameReferencesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDomainNameReferences - List domainNameReferences. Retrieve a list of directoryObject with a reference to the
// domain. The returned list will contain all directory objects that have a dependency on the domain.
func (c DomainNameReferenceClient) ListDomainNameReferences(ctx context.Context, id stable.DomainId, options ListDomainNameReferencesOperationOptions) (result ListDomainNameReferencesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDomainNameReferencesCustomPager{},
		Path:          fmt.Sprintf("%s/domainNameReferences", id.ID()),
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

	temp := make([]stable.DirectoryObject, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := stable.UnmarshalDirectoryObjectImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for stable.DirectoryObject (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListDomainNameReferencesComplete retrieves all the results into a single object
func (c DomainNameReferenceClient) ListDomainNameReferencesComplete(ctx context.Context, id stable.DomainId, options ListDomainNameReferencesOperationOptions) (ListDomainNameReferencesCompleteResult, error) {
	return c.ListDomainNameReferencesCompleteMatchingPredicate(ctx, id, options, DirectoryObjectOperationPredicate{})
}

// ListDomainNameReferencesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DomainNameReferenceClient) ListDomainNameReferencesCompleteMatchingPredicate(ctx context.Context, id stable.DomainId, options ListDomainNameReferencesOperationOptions, predicate DirectoryObjectOperationPredicate) (result ListDomainNameReferencesCompleteResult, err error) {
	items := make([]stable.DirectoryObject, 0)

	resp, err := c.ListDomainNameReferences(ctx, id, options)
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

	result = ListDomainNameReferencesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
