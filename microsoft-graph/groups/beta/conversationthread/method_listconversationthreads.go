package conversationthread

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

type ListConversationThreadsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ConversationThread
}

type ListConversationThreadsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ConversationThread
}

type ListConversationThreadsOperationOptions struct {
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

func DefaultListConversationThreadsOperationOptions() ListConversationThreadsOperationOptions {
	return ListConversationThreadsOperationOptions{}
}

func (o ListConversationThreadsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListConversationThreadsOperationOptions) ToOData() *odata.Query {
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

func (o ListConversationThreadsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListConversationThreadsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListConversationThreadsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListConversationThreads - List threads. Get all the threads in a group conversation. Note: You can also get all the
// threads of a group.
func (c ConversationThreadClient) ListConversationThreads(ctx context.Context, id beta.GroupIdConversationId, options ListConversationThreadsOperationOptions) (result ListConversationThreadsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListConversationThreadsCustomPager{},
		Path:          fmt.Sprintf("%s/threads", id.ID()),
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
		Values *[]beta.ConversationThread `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListConversationThreadsComplete retrieves all the results into a single object
func (c ConversationThreadClient) ListConversationThreadsComplete(ctx context.Context, id beta.GroupIdConversationId, options ListConversationThreadsOperationOptions) (ListConversationThreadsCompleteResult, error) {
	return c.ListConversationThreadsCompleteMatchingPredicate(ctx, id, options, ConversationThreadOperationPredicate{})
}

// ListConversationThreadsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ConversationThreadClient) ListConversationThreadsCompleteMatchingPredicate(ctx context.Context, id beta.GroupIdConversationId, options ListConversationThreadsOperationOptions, predicate ConversationThreadOperationPredicate) (result ListConversationThreadsCompleteResult, err error) {
	items := make([]beta.ConversationThread, 0)

	resp, err := c.ListConversationThreads(ctx, id, options)
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

	result = ListConversationThreadsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
