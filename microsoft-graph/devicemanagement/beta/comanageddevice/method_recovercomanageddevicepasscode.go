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

type RecoverComanagedDevicePasscodeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type RecoverComanagedDevicePasscodeOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultRecoverComanagedDevicePasscodeOperationOptions() RecoverComanagedDevicePasscodeOperationOptions {
	return RecoverComanagedDevicePasscodeOperationOptions{}
}

func (o RecoverComanagedDevicePasscodeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o RecoverComanagedDevicePasscodeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o RecoverComanagedDevicePasscodeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// RecoverComanagedDevicePasscode - Invoke action recoverPasscode. Recover passcode
func (c ComanagedDeviceClient) RecoverComanagedDevicePasscode(ctx context.Context, id beta.DeviceManagementComanagedDeviceId, options RecoverComanagedDevicePasscodeOperationOptions) (result RecoverComanagedDevicePasscodeOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/recoverPasscode", id.ID()),
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
