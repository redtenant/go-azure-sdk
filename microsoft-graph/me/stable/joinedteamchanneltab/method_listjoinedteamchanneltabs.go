package joinedteamchanneltab

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

type ListJoinedTeamChannelTabsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.TeamsTab
}

type ListJoinedTeamChannelTabsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.TeamsTab
}

type ListJoinedTeamChannelTabsOperationOptions struct {
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

func DefaultListJoinedTeamChannelTabsOperationOptions() ListJoinedTeamChannelTabsOperationOptions {
	return ListJoinedTeamChannelTabsOperationOptions{}
}

func (o ListJoinedTeamChannelTabsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListJoinedTeamChannelTabsOperationOptions) ToOData() *odata.Query {
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

func (o ListJoinedTeamChannelTabsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListJoinedTeamChannelTabsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListJoinedTeamChannelTabsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListJoinedTeamChannelTabs - Get tabs from me. A collection of all the tabs in the channel. A navigation property.
func (c JoinedTeamChannelTabClient) ListJoinedTeamChannelTabs(ctx context.Context, id stable.MeJoinedTeamIdChannelId, options ListJoinedTeamChannelTabsOperationOptions) (result ListJoinedTeamChannelTabsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListJoinedTeamChannelTabsCustomPager{},
		Path:          fmt.Sprintf("%s/tabs", id.ID()),
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
		Values *[]stable.TeamsTab `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListJoinedTeamChannelTabsComplete retrieves all the results into a single object
func (c JoinedTeamChannelTabClient) ListJoinedTeamChannelTabsComplete(ctx context.Context, id stable.MeJoinedTeamIdChannelId, options ListJoinedTeamChannelTabsOperationOptions) (ListJoinedTeamChannelTabsCompleteResult, error) {
	return c.ListJoinedTeamChannelTabsCompleteMatchingPredicate(ctx, id, options, TeamsTabOperationPredicate{})
}

// ListJoinedTeamChannelTabsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c JoinedTeamChannelTabClient) ListJoinedTeamChannelTabsCompleteMatchingPredicate(ctx context.Context, id stable.MeJoinedTeamIdChannelId, options ListJoinedTeamChannelTabsOperationOptions, predicate TeamsTabOperationPredicate) (result ListJoinedTeamChannelTabsCompleteResult, err error) {
	items := make([]stable.TeamsTab, 0)

	resp, err := c.ListJoinedTeamChannelTabs(ctx, id, options)
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

	result = ListJoinedTeamChannelTabsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
