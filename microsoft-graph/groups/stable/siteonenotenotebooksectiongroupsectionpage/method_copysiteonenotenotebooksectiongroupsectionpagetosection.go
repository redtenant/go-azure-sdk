package siteonenotenotebooksectiongroupsectionpage

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

type CopySiteOnenoteNotebookSectionGroupSectionPageToSectionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.OnenoteOperation
}

type CopySiteOnenoteNotebookSectionGroupSectionPageToSectionOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCopySiteOnenoteNotebookSectionGroupSectionPageToSectionOperationOptions() CopySiteOnenoteNotebookSectionGroupSectionPageToSectionOperationOptions {
	return CopySiteOnenoteNotebookSectionGroupSectionPageToSectionOperationOptions{}
}

func (o CopySiteOnenoteNotebookSectionGroupSectionPageToSectionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CopySiteOnenoteNotebookSectionGroupSectionPageToSectionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CopySiteOnenoteNotebookSectionGroupSectionPageToSectionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CopySiteOnenoteNotebookSectionGroupSectionPageToSection - Invoke action copyToSection. Copy a page to a specific
// section. For copy operations, you follow an asynchronous calling pattern: First call the Copy action, and then poll
// the operation endpoint for the result.
func (c SiteOnenoteNotebookSectionGroupSectionPageClient) CopySiteOnenoteNotebookSectionGroupSectionPageToSection(ctx context.Context, id stable.GroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionIdPageId, input CopySiteOnenoteNotebookSectionGroupSectionPageToSectionRequest, options CopySiteOnenoteNotebookSectionGroupSectionPageToSectionOperationOptions) (result CopySiteOnenoteNotebookSectionGroupSectionPageToSectionOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/copyToSection", id.ID()),
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

	var model stable.OnenoteOperation
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
