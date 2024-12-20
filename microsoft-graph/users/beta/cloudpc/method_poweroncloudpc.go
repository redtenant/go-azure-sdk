package cloudpc

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

type PowerOnCloudPCOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type PowerOnCloudPCOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultPowerOnCloudPCOperationOptions() PowerOnCloudPCOperationOptions {
	return PowerOnCloudPCOperationOptions{}
}

func (o PowerOnCloudPCOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o PowerOnCloudPCOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o PowerOnCloudPCOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// PowerOnCloudPC - Invoke action powerOn. Power on a Windows 365 Frontline Cloud PC. This action supports Microsoft
// Endpoint Manager (MEM) admin scenarios. After a Windows 365 Frontline Cloud PC is powered on, it is allocated to a
// user, and licenses are assigned immediately. Only IT admin users can perform this action.
func (c CloudPCClient) PowerOnCloudPC(ctx context.Context, id beta.UserIdCloudPCId, options PowerOnCloudPCOperationOptions) (result PowerOnCloudPCOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/powerOn", id.ID()),
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
