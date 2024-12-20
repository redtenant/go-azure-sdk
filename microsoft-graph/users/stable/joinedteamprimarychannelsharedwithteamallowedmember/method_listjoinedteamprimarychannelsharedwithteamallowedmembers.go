package joinedteamprimarychannelsharedwithteamallowedmember

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

type ListJoinedTeamPrimaryChannelSharedWithTeamAllowedMembersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ConversationMember
}

type ListJoinedTeamPrimaryChannelSharedWithTeamAllowedMembersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ConversationMember
}

type ListJoinedTeamPrimaryChannelSharedWithTeamAllowedMembersOperationOptions struct {
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

func DefaultListJoinedTeamPrimaryChannelSharedWithTeamAllowedMembersOperationOptions() ListJoinedTeamPrimaryChannelSharedWithTeamAllowedMembersOperationOptions {
	return ListJoinedTeamPrimaryChannelSharedWithTeamAllowedMembersOperationOptions{}
}

func (o ListJoinedTeamPrimaryChannelSharedWithTeamAllowedMembersOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListJoinedTeamPrimaryChannelSharedWithTeamAllowedMembersOperationOptions) ToOData() *odata.Query {
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

func (o ListJoinedTeamPrimaryChannelSharedWithTeamAllowedMembersOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListJoinedTeamPrimaryChannelSharedWithTeamAllowedMembersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListJoinedTeamPrimaryChannelSharedWithTeamAllowedMembersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListJoinedTeamPrimaryChannelSharedWithTeamAllowedMembers - Get allowedMembers from users. A collection of team
// members who have access to the shared channel.
func (c JoinedTeamPrimaryChannelSharedWithTeamAllowedMemberClient) ListJoinedTeamPrimaryChannelSharedWithTeamAllowedMembers(ctx context.Context, id stable.UserIdJoinedTeamIdPrimaryChannelSharedWithTeamId, options ListJoinedTeamPrimaryChannelSharedWithTeamAllowedMembersOperationOptions) (result ListJoinedTeamPrimaryChannelSharedWithTeamAllowedMembersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListJoinedTeamPrimaryChannelSharedWithTeamAllowedMembersCustomPager{},
		Path:          fmt.Sprintf("%s/allowedMembers", id.ID()),
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

// ListJoinedTeamPrimaryChannelSharedWithTeamAllowedMembersComplete retrieves all the results into a single object
func (c JoinedTeamPrimaryChannelSharedWithTeamAllowedMemberClient) ListJoinedTeamPrimaryChannelSharedWithTeamAllowedMembersComplete(ctx context.Context, id stable.UserIdJoinedTeamIdPrimaryChannelSharedWithTeamId, options ListJoinedTeamPrimaryChannelSharedWithTeamAllowedMembersOperationOptions) (ListJoinedTeamPrimaryChannelSharedWithTeamAllowedMembersCompleteResult, error) {
	return c.ListJoinedTeamPrimaryChannelSharedWithTeamAllowedMembersCompleteMatchingPredicate(ctx, id, options, ConversationMemberOperationPredicate{})
}

// ListJoinedTeamPrimaryChannelSharedWithTeamAllowedMembersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c JoinedTeamPrimaryChannelSharedWithTeamAllowedMemberClient) ListJoinedTeamPrimaryChannelSharedWithTeamAllowedMembersCompleteMatchingPredicate(ctx context.Context, id stable.UserIdJoinedTeamIdPrimaryChannelSharedWithTeamId, options ListJoinedTeamPrimaryChannelSharedWithTeamAllowedMembersOperationOptions, predicate ConversationMemberOperationPredicate) (result ListJoinedTeamPrimaryChannelSharedWithTeamAllowedMembersCompleteResult, err error) {
	items := make([]stable.ConversationMember, 0)

	resp, err := c.ListJoinedTeamPrimaryChannelSharedWithTeamAllowedMembers(ctx, id, options)
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

	result = ListJoinedTeamPrimaryChannelSharedWithTeamAllowedMembersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
