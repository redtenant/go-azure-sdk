package windowsupdatecatalogitem

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateWindowsUpdateCatalogItemOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.WindowsUpdateCatalogItem
}

type CreateWindowsUpdateCatalogItemOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateWindowsUpdateCatalogItemOperationOptions() CreateWindowsUpdateCatalogItemOperationOptions {
	return CreateWindowsUpdateCatalogItemOperationOptions{}
}

func (o CreateWindowsUpdateCatalogItemOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateWindowsUpdateCatalogItemOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateWindowsUpdateCatalogItemOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateWindowsUpdateCatalogItem - Create new navigation property to windowsUpdateCatalogItems for deviceManagement
func (c WindowsUpdateCatalogItemClient) CreateWindowsUpdateCatalogItem(ctx context.Context, input beta.WindowsUpdateCatalogItem, options CreateWindowsUpdateCatalogItemOperationOptions) (result CreateWindowsUpdateCatalogItemOperationResponse, err error) {
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
		Path:          "/deviceManagement/windowsUpdateCatalogItems",
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

	var respObj json.RawMessage
	if err = resp.Unmarshal(&respObj); err != nil {
		return
	}
	model, err := beta.UnmarshalWindowsUpdateCatalogItemImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
