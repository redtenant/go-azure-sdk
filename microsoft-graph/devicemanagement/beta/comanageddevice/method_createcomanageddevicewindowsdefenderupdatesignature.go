package comanageddevice

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

type CreateComanagedDeviceWindowsDefenderUpdateSignatureOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateComanagedDeviceWindowsDefenderUpdateSignatureOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateComanagedDeviceWindowsDefenderUpdateSignatureOperationOptions() CreateComanagedDeviceWindowsDefenderUpdateSignatureOperationOptions {
	return CreateComanagedDeviceWindowsDefenderUpdateSignatureOperationOptions{}
}

func (o CreateComanagedDeviceWindowsDefenderUpdateSignatureOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateComanagedDeviceWindowsDefenderUpdateSignatureOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateComanagedDeviceWindowsDefenderUpdateSignatureOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateComanagedDeviceWindowsDefenderUpdateSignature - Invoke action windowsDefenderUpdateSignatures
func (c ComanagedDeviceClient) CreateComanagedDeviceWindowsDefenderUpdateSignature(ctx context.Context, id beta.DeviceManagementComanagedDeviceId, options CreateComanagedDeviceWindowsDefenderUpdateSignatureOperationOptions) (result CreateComanagedDeviceWindowsDefenderUpdateSignatureOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/windowsDefenderUpdateSignatures", id.ID()),
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
