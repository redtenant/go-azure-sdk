package apischema

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspaceApiSchemaDeleteOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type WorkspaceApiSchemaDeleteOperationOptions struct {
	Force   *bool
	IfMatch *string
}

func DefaultWorkspaceApiSchemaDeleteOperationOptions() WorkspaceApiSchemaDeleteOperationOptions {
	return WorkspaceApiSchemaDeleteOperationOptions{}
}

func (o WorkspaceApiSchemaDeleteOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o WorkspaceApiSchemaDeleteOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o WorkspaceApiSchemaDeleteOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Force != nil {
		out.Append("force", fmt.Sprintf("%v", *o.Force))
	}
	return &out
}

// WorkspaceApiSchemaDelete ...
func (c ApiSchemaClient) WorkspaceApiSchemaDelete(ctx context.Context, id WorkspaceApiSchemaId, options WorkspaceApiSchemaDeleteOperationOptions) (result WorkspaceApiSchemaDeleteOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          id.ID(),
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

	return
}
