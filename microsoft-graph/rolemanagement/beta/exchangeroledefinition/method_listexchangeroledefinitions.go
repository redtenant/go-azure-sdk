package exchangeroledefinition

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

type ListExchangeRoleDefinitionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UnifiedRoleDefinition
}

type ListExchangeRoleDefinitionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UnifiedRoleDefinition
}

type ListExchangeRoleDefinitionsOperationOptions struct {
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

func DefaultListExchangeRoleDefinitionsOperationOptions() ListExchangeRoleDefinitionsOperationOptions {
	return ListExchangeRoleDefinitionsOperationOptions{}
}

func (o ListExchangeRoleDefinitionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListExchangeRoleDefinitionsOperationOptions) ToOData() *odata.Query {
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

func (o ListExchangeRoleDefinitionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListExchangeRoleDefinitionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListExchangeRoleDefinitionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListExchangeRoleDefinitions - List roleDefinitions. Get a list of unifiedRoleDefinition objects for an RBAC provider.
// The following RBAC providers are currently supported: - Cloud PC - device management (Intune) - directory (Microsoft
// Entra ID) - entitlement management (Microsoft Entra ID) - Exchange Online
func (c ExchangeRoleDefinitionClient) ListExchangeRoleDefinitions(ctx context.Context, options ListExchangeRoleDefinitionsOperationOptions) (result ListExchangeRoleDefinitionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListExchangeRoleDefinitionsCustomPager{},
		Path:          "/roleManagement/exchange/roleDefinitions",
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
		Values *[]beta.UnifiedRoleDefinition `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListExchangeRoleDefinitionsComplete retrieves all the results into a single object
func (c ExchangeRoleDefinitionClient) ListExchangeRoleDefinitionsComplete(ctx context.Context, options ListExchangeRoleDefinitionsOperationOptions) (ListExchangeRoleDefinitionsCompleteResult, error) {
	return c.ListExchangeRoleDefinitionsCompleteMatchingPredicate(ctx, options, UnifiedRoleDefinitionOperationPredicate{})
}

// ListExchangeRoleDefinitionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ExchangeRoleDefinitionClient) ListExchangeRoleDefinitionsCompleteMatchingPredicate(ctx context.Context, options ListExchangeRoleDefinitionsOperationOptions, predicate UnifiedRoleDefinitionOperationPredicate) (result ListExchangeRoleDefinitionsCompleteResult, err error) {
	items := make([]beta.UnifiedRoleDefinition, 0)

	resp, err := c.ListExchangeRoleDefinitions(ctx, options)
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

	result = ListExchangeRoleDefinitionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
