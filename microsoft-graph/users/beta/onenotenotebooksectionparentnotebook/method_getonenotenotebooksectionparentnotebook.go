package onenotenotebooksectionparentnotebook

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

type GetOnenoteNotebookSectionParentNotebookOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.Notebook
}

type GetOnenoteNotebookSectionParentNotebookOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetOnenoteNotebookSectionParentNotebookOperationOptions() GetOnenoteNotebookSectionParentNotebookOperationOptions {
	return GetOnenoteNotebookSectionParentNotebookOperationOptions{}
}

func (o GetOnenoteNotebookSectionParentNotebookOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetOnenoteNotebookSectionParentNotebookOperationOptions) ToOData() *odata.Query {
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

func (o GetOnenoteNotebookSectionParentNotebookOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetOnenoteNotebookSectionParentNotebook - Get parentNotebook from users. The notebook that contains the section.
// Read-only.
func (c OnenoteNotebookSectionParentNotebookClient) GetOnenoteNotebookSectionParentNotebook(ctx context.Context, id beta.UserIdOnenoteNotebookIdSectionId, options GetOnenoteNotebookSectionParentNotebookOperationOptions) (result GetOnenoteNotebookSectionParentNotebookOperationResponse, err error) {
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
