package mailfolderchildfolder

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

type ListMailFolderChildFoldersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.MailFolder
}

type ListMailFolderChildFoldersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.MailFolder
}

type ListMailFolderChildFoldersOperationOptions struct {
	Count                *bool
	Expand               *odata.Expand
	Filter               *string
	IncludeHiddenFolders *string
	Metadata             *odata.Metadata
	OrderBy              *odata.OrderBy
	RetryFunc            client.RequestRetryFunc
	Search               *string
	Select               *[]string
	Skip                 *int64
	Top                  *int64
}

func DefaultListMailFolderChildFoldersOperationOptions() ListMailFolderChildFoldersOperationOptions {
	return ListMailFolderChildFoldersOperationOptions{}
}

func (o ListMailFolderChildFoldersOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListMailFolderChildFoldersOperationOptions) ToOData() *odata.Query {
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

func (o ListMailFolderChildFoldersOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.IncludeHiddenFolders != nil {
		out.Append("includeHiddenFolders", fmt.Sprintf("%v", *o.IncludeHiddenFolders))
	}
	return &out
}

type ListMailFolderChildFoldersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListMailFolderChildFoldersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListMailFolderChildFolders - Get childFolders from users. The collection of child folders in the mailFolder.
func (c MailFolderChildFolderClient) ListMailFolderChildFolders(ctx context.Context, id stable.UserIdMailFolderId, options ListMailFolderChildFoldersOperationOptions) (result ListMailFolderChildFoldersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListMailFolderChildFoldersCustomPager{},
		Path:          fmt.Sprintf("%s/childFolders", id.ID()),
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

	temp := make([]stable.MailFolder, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := stable.UnmarshalMailFolderImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for stable.MailFolder (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListMailFolderChildFoldersComplete retrieves all the results into a single object
func (c MailFolderChildFolderClient) ListMailFolderChildFoldersComplete(ctx context.Context, id stable.UserIdMailFolderId, options ListMailFolderChildFoldersOperationOptions) (ListMailFolderChildFoldersCompleteResult, error) {
	return c.ListMailFolderChildFoldersCompleteMatchingPredicate(ctx, id, options, MailFolderOperationPredicate{})
}

// ListMailFolderChildFoldersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MailFolderChildFolderClient) ListMailFolderChildFoldersCompleteMatchingPredicate(ctx context.Context, id stable.UserIdMailFolderId, options ListMailFolderChildFoldersOperationOptions, predicate MailFolderOperationPredicate) (result ListMailFolderChildFoldersCompleteResult, err error) {
	items := make([]stable.MailFolder, 0)

	resp, err := c.ListMailFolderChildFolders(ctx, id, options)
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

	result = ListMailFolderChildFoldersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
