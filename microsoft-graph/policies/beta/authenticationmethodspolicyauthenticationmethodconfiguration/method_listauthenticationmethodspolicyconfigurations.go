package authenticationmethodspolicyauthenticationmethodconfiguration

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListAuthenticationMethodsPolicyConfigurationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AuthenticationMethodConfiguration
}

type ListAuthenticationMethodsPolicyConfigurationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AuthenticationMethodConfiguration
}

type ListAuthenticationMethodsPolicyConfigurationsOperationOptions struct {
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

func DefaultListAuthenticationMethodsPolicyConfigurationsOperationOptions() ListAuthenticationMethodsPolicyConfigurationsOperationOptions {
	return ListAuthenticationMethodsPolicyConfigurationsOperationOptions{}
}

func (o ListAuthenticationMethodsPolicyConfigurationsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListAuthenticationMethodsPolicyConfigurationsOperationOptions) ToOData() *odata.Query {
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

func (o ListAuthenticationMethodsPolicyConfigurationsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListAuthenticationMethodsPolicyConfigurationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAuthenticationMethodsPolicyConfigurationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAuthenticationMethodsPolicyConfigurations - Get externalAuthenticationMethodConfiguration. Read the properties
// and relationships of an externalAuthenticationMethodConfiguration object.
func (c AuthenticationMethodsPolicyAuthenticationMethodConfigurationClient) ListAuthenticationMethodsPolicyConfigurations(ctx context.Context, options ListAuthenticationMethodsPolicyConfigurationsOperationOptions) (result ListAuthenticationMethodsPolicyConfigurationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListAuthenticationMethodsPolicyConfigurationsCustomPager{},
		Path:          "/policies/authenticationMethodsPolicy/authenticationMethodConfigurations",
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

	temp := make([]beta.AuthenticationMethodConfiguration, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := beta.UnmarshalAuthenticationMethodConfigurationImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for beta.AuthenticationMethodConfiguration (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListAuthenticationMethodsPolicyConfigurationsComplete retrieves all the results into a single object
func (c AuthenticationMethodsPolicyAuthenticationMethodConfigurationClient) ListAuthenticationMethodsPolicyConfigurationsComplete(ctx context.Context, options ListAuthenticationMethodsPolicyConfigurationsOperationOptions) (ListAuthenticationMethodsPolicyConfigurationsCompleteResult, error) {
	return c.ListAuthenticationMethodsPolicyConfigurationsCompleteMatchingPredicate(ctx, options, AuthenticationMethodConfigurationOperationPredicate{})
}

// ListAuthenticationMethodsPolicyConfigurationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AuthenticationMethodsPolicyAuthenticationMethodConfigurationClient) ListAuthenticationMethodsPolicyConfigurationsCompleteMatchingPredicate(ctx context.Context, options ListAuthenticationMethodsPolicyConfigurationsOperationOptions, predicate AuthenticationMethodConfigurationOperationPredicate) (result ListAuthenticationMethodsPolicyConfigurationsCompleteResult, err error) {
	items := make([]beta.AuthenticationMethodConfiguration, 0)

	resp, err := c.ListAuthenticationMethodsPolicyConfigurations(ctx, options)
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

	result = ListAuthenticationMethodsPolicyConfigurationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
