package entitlementmanagementresourcescoperesourcerole

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

type ListEntitlementManagementResourceScopeResourceRolesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AccessPackageResourceRole
}

type ListEntitlementManagementResourceScopeResourceRolesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AccessPackageResourceRole
}

type ListEntitlementManagementResourceScopeResourceRolesOperationOptions struct {
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

func DefaultListEntitlementManagementResourceScopeResourceRolesOperationOptions() ListEntitlementManagementResourceScopeResourceRolesOperationOptions {
	return ListEntitlementManagementResourceScopeResourceRolesOperationOptions{}
}

func (o ListEntitlementManagementResourceScopeResourceRolesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEntitlementManagementResourceScopeResourceRolesOperationOptions) ToOData() *odata.Query {
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

func (o ListEntitlementManagementResourceScopeResourceRolesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEntitlementManagementResourceScopeResourceRolesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementResourceScopeResourceRolesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementResourceScopeResourceRoles - Get roles from identityGovernance. Read-only. Nullable.
// Supports $expand.
func (c EntitlementManagementResourceScopeResourceRoleClient) ListEntitlementManagementResourceScopeResourceRoles(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementResourceIdScopeId, options ListEntitlementManagementResourceScopeResourceRolesOperationOptions) (result ListEntitlementManagementResourceScopeResourceRolesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEntitlementManagementResourceScopeResourceRolesCustomPager{},
		Path:          fmt.Sprintf("%s/resource/roles", id.ID()),
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
		Values *[]stable.AccessPackageResourceRole `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEntitlementManagementResourceScopeResourceRolesComplete retrieves all the results into a single object
func (c EntitlementManagementResourceScopeResourceRoleClient) ListEntitlementManagementResourceScopeResourceRolesComplete(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementResourceIdScopeId, options ListEntitlementManagementResourceScopeResourceRolesOperationOptions) (ListEntitlementManagementResourceScopeResourceRolesCompleteResult, error) {
	return c.ListEntitlementManagementResourceScopeResourceRolesCompleteMatchingPredicate(ctx, id, options, AccessPackageResourceRoleOperationPredicate{})
}

// ListEntitlementManagementResourceScopeResourceRolesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementResourceScopeResourceRoleClient) ListEntitlementManagementResourceScopeResourceRolesCompleteMatchingPredicate(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementResourceIdScopeId, options ListEntitlementManagementResourceScopeResourceRolesOperationOptions, predicate AccessPackageResourceRoleOperationPredicate) (result ListEntitlementManagementResourceScopeResourceRolesCompleteResult, err error) {
	items := make([]stable.AccessPackageResourceRole, 0)

	resp, err := c.ListEntitlementManagementResourceScopeResourceRoles(ctx, id, options)
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

	result = ListEntitlementManagementResourceScopeResourceRolesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
