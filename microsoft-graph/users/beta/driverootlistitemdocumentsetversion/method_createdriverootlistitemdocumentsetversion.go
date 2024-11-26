package driverootlistitemdocumentsetversion

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

type CreateDriveRootListItemDocumentSetVersionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.DocumentSetVersion
}

type CreateDriveRootListItemDocumentSetVersionOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateDriveRootListItemDocumentSetVersionOperationOptions() CreateDriveRootListItemDocumentSetVersionOperationOptions {
	return CreateDriveRootListItemDocumentSetVersionOperationOptions{}
}

func (o CreateDriveRootListItemDocumentSetVersionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateDriveRootListItemDocumentSetVersionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateDriveRootListItemDocumentSetVersionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateDriveRootListItemDocumentSetVersion - Create new navigation property to documentSetVersions for users
func (c DriveRootListItemDocumentSetVersionClient) CreateDriveRootListItemDocumentSetVersion(ctx context.Context, id beta.UserIdDriveId, input beta.DocumentSetVersion, options CreateDriveRootListItemDocumentSetVersionOperationOptions) (result CreateDriveRootListItemDocumentSetVersionOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/root/listItem/documentSetVersions", id.ID()),
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

	var model beta.DocumentSetVersion
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
