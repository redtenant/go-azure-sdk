package siteonenotenotebooksectiongroup

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

type ListSiteOnenoteNotebookSectionGroupsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.SectionGroup
}

type ListSiteOnenoteNotebookSectionGroupsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.SectionGroup
}

type ListSiteOnenoteNotebookSectionGroupsOperationOptions struct {
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

func DefaultListSiteOnenoteNotebookSectionGroupsOperationOptions() ListSiteOnenoteNotebookSectionGroupsOperationOptions {
	return ListSiteOnenoteNotebookSectionGroupsOperationOptions{}
}

func (o ListSiteOnenoteNotebookSectionGroupsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListSiteOnenoteNotebookSectionGroupsOperationOptions) ToOData() *odata.Query {
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

func (o ListSiteOnenoteNotebookSectionGroupsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListSiteOnenoteNotebookSectionGroupsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSiteOnenoteNotebookSectionGroupsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSiteOnenoteNotebookSectionGroups - Get sectionGroups from groups. The section groups in the notebook. Read-only.
// Nullable.
func (c SiteOnenoteNotebookSectionGroupClient) ListSiteOnenoteNotebookSectionGroups(ctx context.Context, id stable.GroupIdSiteIdOnenoteNotebookId, options ListSiteOnenoteNotebookSectionGroupsOperationOptions) (result ListSiteOnenoteNotebookSectionGroupsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListSiteOnenoteNotebookSectionGroupsCustomPager{},
		Path:          fmt.Sprintf("%s/sectionGroups", id.ID()),
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
		Values *[]stable.SectionGroup `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSiteOnenoteNotebookSectionGroupsComplete retrieves all the results into a single object
func (c SiteOnenoteNotebookSectionGroupClient) ListSiteOnenoteNotebookSectionGroupsComplete(ctx context.Context, id stable.GroupIdSiteIdOnenoteNotebookId, options ListSiteOnenoteNotebookSectionGroupsOperationOptions) (ListSiteOnenoteNotebookSectionGroupsCompleteResult, error) {
	return c.ListSiteOnenoteNotebookSectionGroupsCompleteMatchingPredicate(ctx, id, options, SectionGroupOperationPredicate{})
}

// ListSiteOnenoteNotebookSectionGroupsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteOnenoteNotebookSectionGroupClient) ListSiteOnenoteNotebookSectionGroupsCompleteMatchingPredicate(ctx context.Context, id stable.GroupIdSiteIdOnenoteNotebookId, options ListSiteOnenoteNotebookSectionGroupsOperationOptions, predicate SectionGroupOperationPredicate) (result ListSiteOnenoteNotebookSectionGroupsCompleteResult, err error) {
	items := make([]stable.SectionGroup, 0)

	resp, err := c.ListSiteOnenoteNotebookSectionGroups(ctx, id, options)
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

	result = ListSiteOnenoteNotebookSectionGroupsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
