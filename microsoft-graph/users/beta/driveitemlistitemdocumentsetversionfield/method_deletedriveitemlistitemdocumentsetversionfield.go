package driveitemlistitemdocumentsetversionfield

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

type DeleteDriveItemListItemDocumentSetVersionFieldOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteDriveItemListItemDocumentSetVersionFieldOperationOptions struct {
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultDeleteDriveItemListItemDocumentSetVersionFieldOperationOptions() DeleteDriveItemListItemDocumentSetVersionFieldOperationOptions {
	return DeleteDriveItemListItemDocumentSetVersionFieldOperationOptions{}
}

func (o DeleteDriveItemListItemDocumentSetVersionFieldOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteDriveItemListItemDocumentSetVersionFieldOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeleteDriveItemListItemDocumentSetVersionFieldOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteDriveItemListItemDocumentSetVersionField - Delete navigation property fields for users
func (c DriveItemListItemDocumentSetVersionFieldClient) DeleteDriveItemListItemDocumentSetVersionField(ctx context.Context, id beta.UserIdDriveIdItemIdListItemDocumentSetVersionId, options DeleteDriveItemListItemDocumentSetVersionFieldOperationOptions) (result DeleteDriveItemListItemDocumentSetVersionFieldOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/fields", id.ID()),
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

	return
}
