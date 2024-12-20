package teammember

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

type ListTeamMembersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ConversationMember
}

type ListTeamMembersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ConversationMember
}

type ListTeamMembersOperationOptions struct {
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

func DefaultListTeamMembersOperationOptions() ListTeamMembersOperationOptions {
	return ListTeamMembersOperationOptions{}
}

func (o ListTeamMembersOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListTeamMembersOperationOptions) ToOData() *odata.Query {
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

func (o ListTeamMembersOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListTeamMembersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTeamMembersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTeamMembers - Get members from groups. Members and owners of the team.
func (c TeamMemberClient) ListTeamMembers(ctx context.Context, id stable.GroupId, options ListTeamMembersOperationOptions) (result ListTeamMembersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListTeamMembersCustomPager{},
		Path:          fmt.Sprintf("%s/team/members", id.ID()),
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

	temp := make([]stable.ConversationMember, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := stable.UnmarshalConversationMemberImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for stable.ConversationMember (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListTeamMembersComplete retrieves all the results into a single object
func (c TeamMemberClient) ListTeamMembersComplete(ctx context.Context, id stable.GroupId, options ListTeamMembersOperationOptions) (ListTeamMembersCompleteResult, error) {
	return c.ListTeamMembersCompleteMatchingPredicate(ctx, id, options, ConversationMemberOperationPredicate{})
}

// ListTeamMembersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TeamMemberClient) ListTeamMembersCompleteMatchingPredicate(ctx context.Context, id stable.GroupId, options ListTeamMembersOperationOptions, predicate ConversationMemberOperationPredicate) (result ListTeamMembersCompleteResult, err error) {
	items := make([]stable.ConversationMember, 0)

	resp, err := c.ListTeamMembers(ctx, id, options)
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

	result = ListTeamMembersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
