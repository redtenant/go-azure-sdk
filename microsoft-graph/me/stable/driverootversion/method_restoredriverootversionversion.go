package driverootversion

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

type RestoreDriveRootVersionVersionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type RestoreDriveRootVersionVersionOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultRestoreDriveRootVersionVersionOperationOptions() RestoreDriveRootVersionVersionOperationOptions {
	return RestoreDriveRootVersionVersionOperationOptions{}
}

func (o RestoreDriveRootVersionVersionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o RestoreDriveRootVersionVersionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o RestoreDriveRootVersionVersionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// RestoreDriveRootVersionVersion - Invoke action restoreVersion. Restore a previous version of a DriveItem to be the
// current version. This will create a new version with the contents of the previous version, but preserves all existing
// versions of the file.
func (c DriveRootVersionClient) RestoreDriveRootVersionVersion(ctx context.Context, id stable.MeDriveIdRootVersionId, options RestoreDriveRootVersionVersionOperationOptions) (result RestoreDriveRootVersionVersionOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/restoreVersion", id.ID()),
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
