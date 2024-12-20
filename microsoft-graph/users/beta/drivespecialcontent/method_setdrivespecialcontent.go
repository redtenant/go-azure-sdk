package drivespecialcontent

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

type SetDriveSpecialContentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SetDriveSpecialContentOperationOptions struct {
	ContentType string
	Metadata    *odata.Metadata
	RetryFunc   client.RequestRetryFunc
}

func DefaultSetDriveSpecialContentOperationOptions() SetDriveSpecialContentOperationOptions {
	return SetDriveSpecialContentOperationOptions{}
}

func (o SetDriveSpecialContentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SetDriveSpecialContentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SetDriveSpecialContentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SetDriveSpecialContent - Update content for the navigation property special in users. The content stream, if the item
// represents a file. The content property will have a potentially breaking change in behavior in the future. It will
// stream content directly instead of redirecting. To proactively opt in to the new behavior ahead of time, use the
// contentStream property instead.
func (c DriveSpecialContentClient) SetDriveSpecialContent(ctx context.Context, id beta.UserIdDriveIdSpecialId, input []byte, options SetDriveSpecialContentOperationOptions) (result SetDriveSpecialContentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: options.ContentType,
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPut,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/content", id.ID()),
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

	return
}
