package entitlementmanagementresourcerolescope

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

type ListEntitlementManagementResourceRoleScopesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AccessPackageResourceRoleScope
}

type ListEntitlementManagementResourceRoleScopesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AccessPackageResourceRoleScope
}

type ListEntitlementManagementResourceRoleScopesOperationOptions struct {
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

func DefaultListEntitlementManagementResourceRoleScopesOperationOptions() ListEntitlementManagementResourceRoleScopesOperationOptions {
	return ListEntitlementManagementResourceRoleScopesOperationOptions{}
}

func (o ListEntitlementManagementResourceRoleScopesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEntitlementManagementResourceRoleScopesOperationOptions) ToOData() *odata.Query {
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

func (o ListEntitlementManagementResourceRoleScopesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEntitlementManagementResourceRoleScopesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementResourceRoleScopesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementResourceRoleScopes - Get resourceRoleScopes from identityGovernance
func (c EntitlementManagementResourceRoleScopeClient) ListEntitlementManagementResourceRoleScopes(ctx context.Context, options ListEntitlementManagementResourceRoleScopesOperationOptions) (result ListEntitlementManagementResourceRoleScopesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEntitlementManagementResourceRoleScopesCustomPager{},
		Path:          "/identityGovernance/entitlementManagement/resourceRoleScopes",
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
		Values *[]stable.AccessPackageResourceRoleScope `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEntitlementManagementResourceRoleScopesComplete retrieves all the results into a single object
func (c EntitlementManagementResourceRoleScopeClient) ListEntitlementManagementResourceRoleScopesComplete(ctx context.Context, options ListEntitlementManagementResourceRoleScopesOperationOptions) (ListEntitlementManagementResourceRoleScopesCompleteResult, error) {
	return c.ListEntitlementManagementResourceRoleScopesCompleteMatchingPredicate(ctx, options, AccessPackageResourceRoleScopeOperationPredicate{})
}

// ListEntitlementManagementResourceRoleScopesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementResourceRoleScopeClient) ListEntitlementManagementResourceRoleScopesCompleteMatchingPredicate(ctx context.Context, options ListEntitlementManagementResourceRoleScopesOperationOptions, predicate AccessPackageResourceRoleScopeOperationPredicate) (result ListEntitlementManagementResourceRoleScopesCompleteResult, err error) {
	items := make([]stable.AccessPackageResourceRoleScope, 0)

	resp, err := c.ListEntitlementManagementResourceRoleScopes(ctx, options)
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

	result = ListEntitlementManagementResourceRoleScopesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
