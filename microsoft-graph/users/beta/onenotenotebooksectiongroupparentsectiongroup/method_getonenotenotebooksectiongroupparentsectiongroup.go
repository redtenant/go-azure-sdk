package onenotenotebooksectiongroupparentsectiongroup

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

type GetOnenoteNotebookSectionGroupParentSectionGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.SectionGroup
}

type GetOnenoteNotebookSectionGroupParentSectionGroupOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetOnenoteNotebookSectionGroupParentSectionGroupOperationOptions() GetOnenoteNotebookSectionGroupParentSectionGroupOperationOptions {
	return GetOnenoteNotebookSectionGroupParentSectionGroupOperationOptions{}
}

func (o GetOnenoteNotebookSectionGroupParentSectionGroupOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetOnenoteNotebookSectionGroupParentSectionGroupOperationOptions) ToOData() *odata.Query {
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

func (o GetOnenoteNotebookSectionGroupParentSectionGroupOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetOnenoteNotebookSectionGroupParentSectionGroup - Get parentSectionGroup from users. The section group that contains
// the section group. Read-only.
func (c OnenoteNotebookSectionGroupParentSectionGroupClient) GetOnenoteNotebookSectionGroupParentSectionGroup(ctx context.Context, id beta.UserIdOnenoteNotebookIdSectionGroupId, options GetOnenoteNotebookSectionGroupParentSectionGroupOperationOptions) (result GetOnenoteNotebookSectionGroupParentSectionGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/parentSectionGroup", id.ID()),
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

	var model beta.SectionGroup
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
