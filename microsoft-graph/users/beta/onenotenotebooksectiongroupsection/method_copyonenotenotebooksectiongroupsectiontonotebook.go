package onenotenotebooksectiongroupsection

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

type CopyOnenoteNotebookSectionGroupSectionToNotebookOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.OnenoteOperation
}

type CopyOnenoteNotebookSectionGroupSectionToNotebookOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCopyOnenoteNotebookSectionGroupSectionToNotebookOperationOptions() CopyOnenoteNotebookSectionGroupSectionToNotebookOperationOptions {
	return CopyOnenoteNotebookSectionGroupSectionToNotebookOperationOptions{}
}

func (o CopyOnenoteNotebookSectionGroupSectionToNotebookOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CopyOnenoteNotebookSectionGroupSectionToNotebookOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CopyOnenoteNotebookSectionGroupSectionToNotebookOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CopyOnenoteNotebookSectionGroupSectionToNotebook - Invoke action copyToNotebook. Copies a section to a specific
// notebook. For Copy operations, you follow an asynchronous calling pattern: First call the Copy action, and then poll
// the operation endpoint for the result.
func (c OnenoteNotebookSectionGroupSectionClient) CopyOnenoteNotebookSectionGroupSectionToNotebook(ctx context.Context, id beta.UserIdOnenoteNotebookIdSectionGroupIdSectionId, input CopyOnenoteNotebookSectionGroupSectionToNotebookRequest, options CopyOnenoteNotebookSectionGroupSectionToNotebookOperationOptions) (result CopyOnenoteNotebookSectionGroupSectionToNotebookOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/copyToNotebook", id.ID()),
		RetryFunc:     options.RetryFunc,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
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

	var model beta.OnenoteOperation
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
