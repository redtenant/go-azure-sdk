package entitlementmanagementaccesspackageincompatibleaccesspackage

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

type ListEntitlementManagementAccessPackageIncompatibleRefsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.DirectoryObject
}

type ListEntitlementManagementAccessPackageIncompatibleRefsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.DirectoryObject
}

type ListEntitlementManagementAccessPackageIncompatibleRefsOperationOptions struct {
	Count     *bool
	Filter    *string
	Metadata  *odata.Metadata
	OrderBy   *odata.OrderBy
	RetryFunc client.RequestRetryFunc
	Search    *string
	Skip      *int64
	Top       *int64
}

func DefaultListEntitlementManagementAccessPackageIncompatibleRefsOperationOptions() ListEntitlementManagementAccessPackageIncompatibleRefsOperationOptions {
	return ListEntitlementManagementAccessPackageIncompatibleRefsOperationOptions{}
}

func (o ListEntitlementManagementAccessPackageIncompatibleRefsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEntitlementManagementAccessPackageIncompatibleRefsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Count != nil {
		out.Count = *o.Count
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
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListEntitlementManagementAccessPackageIncompatibleRefsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEntitlementManagementAccessPackageIncompatibleRefsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementAccessPackageIncompatibleRefsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementAccessPackageIncompatibleRefs - List incompatibleAccessPackages. Retrieve a list of the
// accessPackage objects that have been marked as incompatible on an accessPackage.
func (c EntitlementManagementAccessPackageIncompatibleAccessPackageClient) ListEntitlementManagementAccessPackageIncompatibleRefs(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementAccessPackageId, options ListEntitlementManagementAccessPackageIncompatibleRefsOperationOptions) (result ListEntitlementManagementAccessPackageIncompatibleRefsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEntitlementManagementAccessPackageIncompatibleRefsCustomPager{},
		Path:          fmt.Sprintf("%s/incompatibleAccessPackages/$ref", id.ID()),
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

// ListEntitlementManagementAccessPackageIncompatibleRefsComplete retrieves all the results into a single object
func (c EntitlementManagementAccessPackageIncompatibleAccessPackageClient) ListEntitlementManagementAccessPackageIncompatibleRefsComplete(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementAccessPackageId, options ListEntitlementManagementAccessPackageIncompatibleRefsOperationOptions) (ListEntitlementManagementAccessPackageIncompatibleRefsCompleteResult, error) {
	return c.ListEntitlementManagementAccessPackageIncompatibleRefsCompleteMatchingPredicate(ctx, id, options, DirectoryObjectOperationPredicate{})
}

// ListEntitlementManagementAccessPackageIncompatibleRefsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementAccessPackageIncompatibleAccessPackageClient) ListEntitlementManagementAccessPackageIncompatibleRefsCompleteMatchingPredicate(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementAccessPackageId, options ListEntitlementManagementAccessPackageIncompatibleRefsOperationOptions, predicate DirectoryObjectOperationPredicate) (result ListEntitlementManagementAccessPackageIncompatibleRefsCompleteResult, err error) {
	items := make([]stable.DirectoryObject, 0)

	resp, err := c.ListEntitlementManagementAccessPackageIncompatibleRefs(ctx, id, options)
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

	result = ListEntitlementManagementAccessPackageIncompatibleRefsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
