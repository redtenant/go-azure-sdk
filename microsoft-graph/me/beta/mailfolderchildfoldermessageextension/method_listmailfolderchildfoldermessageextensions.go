package mailfolderchildfoldermessageextension

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

type ListMailFolderChildFolderMessageExtensionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Extension
}

type ListMailFolderChildFolderMessageExtensionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Extension
}

type ListMailFolderChildFolderMessageExtensionsOperationOptions struct {
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

func DefaultListMailFolderChildFolderMessageExtensionsOperationOptions() ListMailFolderChildFolderMessageExtensionsOperationOptions {
	return ListMailFolderChildFolderMessageExtensionsOperationOptions{}
}

func (o ListMailFolderChildFolderMessageExtensionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListMailFolderChildFolderMessageExtensionsOperationOptions) ToOData() *odata.Query {
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

func (o ListMailFolderChildFolderMessageExtensionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListMailFolderChildFolderMessageExtensionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListMailFolderChildFolderMessageExtensionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListMailFolderChildFolderMessageExtensions - Get extensions from me. The collection of open extensions defined for
// the message. Nullable.
func (c MailFolderChildFolderMessageExtensionClient) ListMailFolderChildFolderMessageExtensions(ctx context.Context, id beta.MeMailFolderIdChildFolderIdMessageId, options ListMailFolderChildFolderMessageExtensionsOperationOptions) (result ListMailFolderChildFolderMessageExtensionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListMailFolderChildFolderMessageExtensionsCustomPager{},
		Path:          fmt.Sprintf("%s/extensions", id.ID()),
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

	temp := make([]beta.Extension, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := beta.UnmarshalExtensionImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for beta.Extension (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListMailFolderChildFolderMessageExtensionsComplete retrieves all the results into a single object
func (c MailFolderChildFolderMessageExtensionClient) ListMailFolderChildFolderMessageExtensionsComplete(ctx context.Context, id beta.MeMailFolderIdChildFolderIdMessageId, options ListMailFolderChildFolderMessageExtensionsOperationOptions) (ListMailFolderChildFolderMessageExtensionsCompleteResult, error) {
	return c.ListMailFolderChildFolderMessageExtensionsCompleteMatchingPredicate(ctx, id, options, ExtensionOperationPredicate{})
}

// ListMailFolderChildFolderMessageExtensionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MailFolderChildFolderMessageExtensionClient) ListMailFolderChildFolderMessageExtensionsCompleteMatchingPredicate(ctx context.Context, id beta.MeMailFolderIdChildFolderIdMessageId, options ListMailFolderChildFolderMessageExtensionsOperationOptions, predicate ExtensionOperationPredicate) (result ListMailFolderChildFolderMessageExtensionsCompleteResult, err error) {
	items := make([]beta.Extension, 0)

	resp, err := c.ListMailFolderChildFolderMessageExtensions(ctx, id, options)
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

	result = ListMailFolderChildFolderMessageExtensionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
