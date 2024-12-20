package microsofttunnelsitemicrosofttunnelserver

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

type CreateMicrosoftTunnelSiteServerLogCollectionRequestOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.MicrosoftTunnelServerLogCollectionResponse
}

type CreateMicrosoftTunnelSiteServerLogCollectionRequestOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateMicrosoftTunnelSiteServerLogCollectionRequestOperationOptions() CreateMicrosoftTunnelSiteServerLogCollectionRequestOperationOptions {
	return CreateMicrosoftTunnelSiteServerLogCollectionRequestOperationOptions{}
}

func (o CreateMicrosoftTunnelSiteServerLogCollectionRequestOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateMicrosoftTunnelSiteServerLogCollectionRequestOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateMicrosoftTunnelSiteServerLogCollectionRequestOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateMicrosoftTunnelSiteServerLogCollectionRequest - Invoke action createServerLogCollectionRequest
func (c MicrosoftTunnelSiteMicrosoftTunnelServerClient) CreateMicrosoftTunnelSiteServerLogCollectionRequest(ctx context.Context, id beta.DeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerId, input CreateMicrosoftTunnelSiteServerLogCollectionRequestRequest, options CreateMicrosoftTunnelSiteServerLogCollectionRequestOperationOptions) (result CreateMicrosoftTunnelSiteServerLogCollectionRequestOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/createServerLogCollectionRequest", id.ID()),
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

	var model beta.MicrosoftTunnelServerLogCollectionResponse
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
