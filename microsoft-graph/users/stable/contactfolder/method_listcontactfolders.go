package contactfolder

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

type ListContactFoldersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ContactFolder
}

type ListContactFoldersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ContactFolder
}

type ListContactFoldersOperationOptions struct {
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

func DefaultListContactFoldersOperationOptions() ListContactFoldersOperationOptions {
	return ListContactFoldersOperationOptions{}
}

func (o ListContactFoldersOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListContactFoldersOperationOptions) ToOData() *odata.Query {
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

func (o ListContactFoldersOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListContactFoldersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListContactFoldersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListContactFolders - Get contactFolders from users. The user's contacts folders. Read-only. Nullable.
func (c ContactFolderClient) ListContactFolders(ctx context.Context, id stable.UserId, options ListContactFoldersOperationOptions) (result ListContactFoldersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListContactFoldersCustomPager{},
		Path:          fmt.Sprintf("%s/contactFolders", id.ID()),
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
		Values *[]stable.ContactFolder `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListContactFoldersComplete retrieves all the results into a single object
func (c ContactFolderClient) ListContactFoldersComplete(ctx context.Context, id stable.UserId, options ListContactFoldersOperationOptions) (ListContactFoldersCompleteResult, error) {
	return c.ListContactFoldersCompleteMatchingPredicate(ctx, id, options, ContactFolderOperationPredicate{})
}

// ListContactFoldersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ContactFolderClient) ListContactFoldersCompleteMatchingPredicate(ctx context.Context, id stable.UserId, options ListContactFoldersOperationOptions, predicate ContactFolderOperationPredicate) (result ListContactFoldersCompleteResult, err error) {
	items := make([]stable.ContactFolder, 0)

	resp, err := c.ListContactFolders(ctx, id, options)
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

	result = ListContactFoldersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
