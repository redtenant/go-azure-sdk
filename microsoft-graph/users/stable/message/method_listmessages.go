package message

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

type ListMessagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.Message
}

type ListMessagesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.Message
}

type ListMessagesOperationOptions struct {
	Count                 *bool
	Expand                *odata.Expand
	Filter                *string
	IncludeHiddenMessages *string
	Metadata              *odata.Metadata
	OrderBy               *odata.OrderBy
	RetryFunc             client.RequestRetryFunc
	Search                *string
	Select                *[]string
	Skip                  *int64
	Top                   *int64
}

func DefaultListMessagesOperationOptions() ListMessagesOperationOptions {
	return ListMessagesOperationOptions{}
}

func (o ListMessagesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListMessagesOperationOptions) ToOData() *odata.Query {
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

func (o ListMessagesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.IncludeHiddenMessages != nil {
		out.Append("includeHiddenMessages", fmt.Sprintf("%v", *o.IncludeHiddenMessages))
	}
	return &out
}

type ListMessagesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListMessagesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListMessages - Get messages from users. The messages in a mailbox or folder. Read-only. Nullable.
func (c MessageClient) ListMessages(ctx context.Context, id stable.UserId, options ListMessagesOperationOptions) (result ListMessagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListMessagesCustomPager{},
		Path:          fmt.Sprintf("%s/messages", id.ID()),
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

	temp := make([]stable.Message, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := stable.UnmarshalMessageImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for stable.Message (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListMessagesComplete retrieves all the results into a single object
func (c MessageClient) ListMessagesComplete(ctx context.Context, id stable.UserId, options ListMessagesOperationOptions) (ListMessagesCompleteResult, error) {
	return c.ListMessagesCompleteMatchingPredicate(ctx, id, options, MessageOperationPredicate{})
}

// ListMessagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MessageClient) ListMessagesCompleteMatchingPredicate(ctx context.Context, id stable.UserId, options ListMessagesOperationOptions, predicate MessageOperationPredicate) (result ListMessagesCompleteResult, err error) {
	items := make([]stable.Message, 0)

	resp, err := c.ListMessages(ctx, id, options)
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

	result = ListMessagesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
