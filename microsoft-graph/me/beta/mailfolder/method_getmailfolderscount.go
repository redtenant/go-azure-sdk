package mailfolder

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetMailFoldersCountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetMailFoldersCountOperationOptions struct {
	Filter               *string
	IncludeHiddenFolders *string
	Metadata             *odata.Metadata
	RetryFunc            client.RequestRetryFunc
	Search               *string
}

func DefaultGetMailFoldersCountOperationOptions() GetMailFoldersCountOperationOptions {
	return GetMailFoldersCountOperationOptions{}
}

func (o GetMailFoldersCountOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetMailFoldersCountOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	return &out
}

func (o GetMailFoldersCountOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.IncludeHiddenFolders != nil {
		out.Append("includeHiddenFolders", fmt.Sprintf("%v", *o.IncludeHiddenFolders))
	}
	return &out
}

// GetMailFoldersCount - Get the number of the resource
func (c MailFolderClient) GetMailFoldersCount(ctx context.Context, options GetMailFoldersCountOperationOptions) (result GetMailFoldersCountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "text/plain",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          "/me/mailFolders/$count",
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

	var model []byte
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
