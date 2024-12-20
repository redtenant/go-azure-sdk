package microsofttunnelserverlogcollectionresponse

import (
	"context"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateMicrosoftTunnelServerLogCollectionResponseOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.MicrosoftTunnelServerLogCollectionResponse
}

type CreateMicrosoftTunnelServerLogCollectionResponseOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateMicrosoftTunnelServerLogCollectionResponseOperationOptions() CreateMicrosoftTunnelServerLogCollectionResponseOperationOptions {
	return CreateMicrosoftTunnelServerLogCollectionResponseOperationOptions{}
}

func (o CreateMicrosoftTunnelServerLogCollectionResponseOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateMicrosoftTunnelServerLogCollectionResponseOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateMicrosoftTunnelServerLogCollectionResponseOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateMicrosoftTunnelServerLogCollectionResponse - Create new navigation property to
// microsoftTunnelServerLogCollectionResponses for deviceManagement
func (c MicrosoftTunnelServerLogCollectionResponseClient) CreateMicrosoftTunnelServerLogCollectionResponse(ctx context.Context, input beta.MicrosoftTunnelServerLogCollectionResponse, options CreateMicrosoftTunnelServerLogCollectionResponseOperationOptions) (result CreateMicrosoftTunnelServerLogCollectionResponseOperationResponse, err error) {
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
		Path:          "/deviceManagement/microsoftTunnelServerLogCollectionResponses",
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
