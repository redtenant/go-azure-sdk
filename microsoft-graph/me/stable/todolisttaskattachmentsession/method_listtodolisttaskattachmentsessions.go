package todolisttaskattachmentsession

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

type ListTodoListTaskAttachmentSessionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AttachmentSession
}

type ListTodoListTaskAttachmentSessionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AttachmentSession
}

type ListTodoListTaskAttachmentSessionsOperationOptions struct {
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

func DefaultListTodoListTaskAttachmentSessionsOperationOptions() ListTodoListTaskAttachmentSessionsOperationOptions {
	return ListTodoListTaskAttachmentSessionsOperationOptions{}
}

func (o ListTodoListTaskAttachmentSessionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListTodoListTaskAttachmentSessionsOperationOptions) ToOData() *odata.Query {
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

func (o ListTodoListTaskAttachmentSessionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListTodoListTaskAttachmentSessionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTodoListTaskAttachmentSessionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTodoListTaskAttachmentSessions - Get attachmentSessions from me
func (c TodoListTaskAttachmentSessionClient) ListTodoListTaskAttachmentSessions(ctx context.Context, id stable.MeTodoListIdTaskId, options ListTodoListTaskAttachmentSessionsOperationOptions) (result ListTodoListTaskAttachmentSessionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListTodoListTaskAttachmentSessionsCustomPager{},
		Path:          fmt.Sprintf("%s/attachmentSessions", id.ID()),
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
		Values *[]stable.AttachmentSession `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTodoListTaskAttachmentSessionsComplete retrieves all the results into a single object
func (c TodoListTaskAttachmentSessionClient) ListTodoListTaskAttachmentSessionsComplete(ctx context.Context, id stable.MeTodoListIdTaskId, options ListTodoListTaskAttachmentSessionsOperationOptions) (ListTodoListTaskAttachmentSessionsCompleteResult, error) {
	return c.ListTodoListTaskAttachmentSessionsCompleteMatchingPredicate(ctx, id, options, AttachmentSessionOperationPredicate{})
}

// ListTodoListTaskAttachmentSessionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TodoListTaskAttachmentSessionClient) ListTodoListTaskAttachmentSessionsCompleteMatchingPredicate(ctx context.Context, id stable.MeTodoListIdTaskId, options ListTodoListTaskAttachmentSessionsOperationOptions, predicate AttachmentSessionOperationPredicate) (result ListTodoListTaskAttachmentSessionsCompleteResult, err error) {
	items := make([]stable.AttachmentSession, 0)

	resp, err := c.ListTodoListTaskAttachmentSessions(ctx, id, options)
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

	result = ListTodoListTaskAttachmentSessionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
