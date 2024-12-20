package entitlementmanagementconnectedorganizationinternalsponsor

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

type ListEntitlementManagementConnectedOrganizationInternalSponsorRefsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.DirectoryObject
}

type ListEntitlementManagementConnectedOrganizationInternalSponsorRefsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.DirectoryObject
}

type ListEntitlementManagementConnectedOrganizationInternalSponsorRefsOperationOptions struct {
	Count     *bool
	Filter    *string
	Metadata  *odata.Metadata
	OrderBy   *odata.OrderBy
	RetryFunc client.RequestRetryFunc
	Search    *string
	Skip      *int64
	Top       *int64
}

func DefaultListEntitlementManagementConnectedOrganizationInternalSponsorRefsOperationOptions() ListEntitlementManagementConnectedOrganizationInternalSponsorRefsOperationOptions {
	return ListEntitlementManagementConnectedOrganizationInternalSponsorRefsOperationOptions{}
}

func (o ListEntitlementManagementConnectedOrganizationInternalSponsorRefsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEntitlementManagementConnectedOrganizationInternalSponsorRefsOperationOptions) ToOData() *odata.Query {
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

func (o ListEntitlementManagementConnectedOrganizationInternalSponsorRefsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEntitlementManagementConnectedOrganizationInternalSponsorRefsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementConnectedOrganizationInternalSponsorRefsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementConnectedOrganizationInternalSponsorRefs - Get ref of internalSponsors from
// identityGovernance
func (c EntitlementManagementConnectedOrganizationInternalSponsorClient) ListEntitlementManagementConnectedOrganizationInternalSponsorRefs(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementConnectedOrganizationId, options ListEntitlementManagementConnectedOrganizationInternalSponsorRefsOperationOptions) (result ListEntitlementManagementConnectedOrganizationInternalSponsorRefsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEntitlementManagementConnectedOrganizationInternalSponsorRefsCustomPager{},
		Path:          fmt.Sprintf("%s/internalSponsors/$ref", id.ID()),
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

// ListEntitlementManagementConnectedOrganizationInternalSponsorRefsComplete retrieves all the results into a single object
func (c EntitlementManagementConnectedOrganizationInternalSponsorClient) ListEntitlementManagementConnectedOrganizationInternalSponsorRefsComplete(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementConnectedOrganizationId, options ListEntitlementManagementConnectedOrganizationInternalSponsorRefsOperationOptions) (ListEntitlementManagementConnectedOrganizationInternalSponsorRefsCompleteResult, error) {
	return c.ListEntitlementManagementConnectedOrganizationInternalSponsorRefsCompleteMatchingPredicate(ctx, id, options, DirectoryObjectOperationPredicate{})
}

// ListEntitlementManagementConnectedOrganizationInternalSponsorRefsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementConnectedOrganizationInternalSponsorClient) ListEntitlementManagementConnectedOrganizationInternalSponsorRefsCompleteMatchingPredicate(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementConnectedOrganizationId, options ListEntitlementManagementConnectedOrganizationInternalSponsorRefsOperationOptions, predicate DirectoryObjectOperationPredicate) (result ListEntitlementManagementConnectedOrganizationInternalSponsorRefsCompleteResult, err error) {
	items := make([]stable.DirectoryObject, 0)

	resp, err := c.ListEntitlementManagementConnectedOrganizationInternalSponsorRefs(ctx, id, options)
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

	result = ListEntitlementManagementConnectedOrganizationInternalSponsorRefsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
