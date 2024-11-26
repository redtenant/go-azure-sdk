package virtualendpointonpremisesconnection

import (
	"context"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateVirtualEndpointOnPremisesConnectionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateVirtualEndpointOnPremisesConnectionOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateVirtualEndpointOnPremisesConnectionOperationOptions() UpdateVirtualEndpointOnPremisesConnectionOperationOptions {
	return UpdateVirtualEndpointOnPremisesConnectionOperationOptions{}
}

func (o UpdateVirtualEndpointOnPremisesConnectionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateVirtualEndpointOnPremisesConnectionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateVirtualEndpointOnPremisesConnectionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateVirtualEndpointOnPremisesConnection - Update cloudPcOnPremisesConnection. Update the properties of a
// cloudPcOnPremisesConnection object.
func (c VirtualEndpointOnPremisesConnectionClient) UpdateVirtualEndpointOnPremisesConnection(ctx context.Context, id beta.DeviceManagementVirtualEndpointOnPremisesConnectionId, input beta.CloudPCOnPremisesConnection, options UpdateVirtualEndpointOnPremisesConnectionOperationOptions) (result UpdateVirtualEndpointOnPremisesConnectionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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
