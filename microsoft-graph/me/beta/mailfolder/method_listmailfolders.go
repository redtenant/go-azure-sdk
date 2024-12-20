package mailfolder

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListMailFoldersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.MailFolder
}

type ListMailFoldersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.MailFolder
}

type ListMailFoldersOperationOptions struct {
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

func DefaultListMailFoldersOperationOptions() ListMailFoldersOperationOptions {
	return ListMailFoldersOperationOptions{}
}

func (o ListMailFoldersOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListMailFoldersOperationOptions) ToOData() *odata.Query {
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

func (o ListMailFoldersOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.IncludeHiddenFolders != nil {
		out.Append("includeHiddenFolders", fmt.Sprintf("%v", *o.IncludeHiddenFolders))
	}
	return &out
}

type ListMailFoldersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListMailFoldersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListMailFolders - List mailFolders. Get the mail folder collection directly under the root folder of the signed-in
// user. The returned collection includes any mail search folders directly under the root. By default, this operation
// does not return hidden folders. Use a query parameter includeHiddenFolders to include them in the response. This
// operation does not return all mail folders in a mailbox, only the child folders of the root folder. To return all
// mail folders in a mailbox, each child folder must be traversed separately.
func (c MailFolderClient) ListMailFolders(ctx context.Context, options ListMailFoldersOperationOptions) (result ListMailFoldersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListMailFoldersCustomPager{},
		Path:          "/me/mailFolders",
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

	temp := make([]beta.MailFolder, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := beta.UnmarshalMailFolderImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for beta.MailFolder (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListMailFoldersComplete retrieves all the results into a single object
func (c MailFolderClient) ListMailFoldersComplete(ctx context.Context, options ListMailFoldersOperationOptions) (ListMailFoldersCompleteResult, error) {
	return c.ListMailFoldersCompleteMatchingPredicate(ctx, options, MailFolderOperationPredicate{})
}

// ListMailFoldersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MailFolderClient) ListMailFoldersCompleteMatchingPredicate(ctx context.Context, options ListMailFoldersOperationOptions, predicate MailFolderOperationPredicate) (result ListMailFoldersCompleteResult, err error) {
	items := make([]beta.MailFolder, 0)

	resp, err := c.ListMailFolders(ctx, options)
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

	result = ListMailFoldersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
