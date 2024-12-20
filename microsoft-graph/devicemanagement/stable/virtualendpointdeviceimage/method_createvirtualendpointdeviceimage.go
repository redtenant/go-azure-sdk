package virtualendpointdeviceimage

import (
	"context"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateVirtualEndpointDeviceImageOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.CloudPCDeviceImage
}

type CreateVirtualEndpointDeviceImageOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateVirtualEndpointDeviceImageOperationOptions() CreateVirtualEndpointDeviceImageOperationOptions {
	return CreateVirtualEndpointDeviceImageOperationOptions{}
}

func (o CreateVirtualEndpointDeviceImageOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateVirtualEndpointDeviceImageOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateVirtualEndpointDeviceImageOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateVirtualEndpointDeviceImage - Create cloudPcDeviceImage. Create a new cloudPcDeviceImage object. Upload a custom
// OS image that you can later provision on Cloud PCs.
func (c VirtualEndpointDeviceImageClient) CreateVirtualEndpointDeviceImage(ctx context.Context, input stable.CloudPCDeviceImage, options CreateVirtualEndpointDeviceImageOperationOptions) (result CreateVirtualEndpointDeviceImageOperationResponse, err error) {
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
		Path:          "/deviceManagement/virtualEndpoint/deviceImages",
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

	var model stable.CloudPCDeviceImage
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
