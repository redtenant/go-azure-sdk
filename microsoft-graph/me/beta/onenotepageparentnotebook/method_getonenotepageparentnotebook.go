package onenotepageparentnotebook

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

type GetOnenotePageParentNotebookOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.Notebook
}

type GetOnenotePageParentNotebookOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetOnenotePageParentNotebookOperationOptions() GetOnenotePageParentNotebookOperationOptions {
	return GetOnenotePageParentNotebookOperationOptions{}
}

func (o GetOnenotePageParentNotebookOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetOnenotePageParentNotebookOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetOnenotePageParentNotebookOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetOnenotePageParentNotebook - Get parentNotebook from me. The notebook that contains the page. Read-only.
func (c OnenotePageParentNotebookClient) GetOnenotePageParentNotebook(ctx context.Context, id beta.MeOnenotePageId, options GetOnenotePageParentNotebookOperationOptions) (result GetOnenotePageParentNotebookOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/parentNotebook", id.ID()),
		RetryFunc:     options.RetryFunc,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var model beta.Notebook
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
