package entitlementmanagementaccesspackageassignmentaccesspackageaccesspackageresourcerolescopeaccesspackageresourcescopeaccesspackageresourceaccesspackageresourcescope

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

type ListEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopeResourceAccessPackageResourceScopesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AccessPackageResourceScope
}

type ListEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopeResourceAccessPackageResourceScopesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AccessPackageResourceScope
}

type ListEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopeResourceAccessPackageResourceScopesOperationOptions struct {
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

func DefaultListEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopeResourceAccessPackageResourceScopesOperationOptions() ListEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopeResourceAccessPackageResourceScopesOperationOptions {
	return ListEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopeResourceAccessPackageResourceScopesOperationOptions{}
}

func (o ListEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopeResourceAccessPackageResourceScopesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopeResourceAccessPackageResourceScopesOperationOptions) ToOData() *odata.Query {
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

func (o ListEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopeResourceAccessPackageResourceScopesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopeResourceAccessPackageResourceScopesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopeResourceAccessPackageResourceScopesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopeResourceAccessPackageResourceScopes -
// Get accessPackageResourceScopes from identityGovernance. Read-only. Nullable. Supports $expand.
func (c EntitlementManagementAccessPackageAssignmentAccessPackageAccessPackageResourceRoleScopeAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopeClient) ListEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopeResourceAccessPackageResourceScopes(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageResourceRoleScopeId, options ListEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopeResourceAccessPackageResourceScopesOperationOptions) (result ListEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopeResourceAccessPackageResourceScopesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopeResourceAccessPackageResourceScopesCustomPager{},
		Path:          fmt.Sprintf("%s/accessPackageResourceScope/accessPackageResource/accessPackageResourceScopes", id.ID()),
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
		Values *[]beta.AccessPackageResourceScope `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopeResourceAccessPackageResourceScopesComplete retrieves all the results into a single object
func (c EntitlementManagementAccessPackageAssignmentAccessPackageAccessPackageResourceRoleScopeAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopeClient) ListEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopeResourceAccessPackageResourceScopesComplete(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageResourceRoleScopeId, options ListEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopeResourceAccessPackageResourceScopesOperationOptions) (ListEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopeResourceAccessPackageResourceScopesCompleteResult, error) {
	return c.ListEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopeResourceAccessPackageResourceScopesCompleteMatchingPredicate(ctx, id, options, AccessPackageResourceScopeOperationPredicate{})
}

// ListEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopeResourceAccessPackageResourceScopesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementAccessPackageAssignmentAccessPackageAccessPackageResourceRoleScopeAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopeClient) ListEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopeResourceAccessPackageResourceScopesCompleteMatchingPredicate(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageResourceRoleScopeId, options ListEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopeResourceAccessPackageResourceScopesOperationOptions, predicate AccessPackageResourceScopeOperationPredicate) (result ListEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopeResourceAccessPackageResourceScopesCompleteResult, err error) {
	items := make([]beta.AccessPackageResourceScope, 0)

	resp, err := c.ListEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopeResourceAccessPackageResourceScopes(ctx, id, options)
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

	result = ListEntitlementManagementAccessPackageAssignmentAccessPackageResourceRoleScopeResourceAccessPackageResourceScopesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
