package privateendpointconnections

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/client/pollers"
	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateEndpointConnectionDeleteOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

// PrivateEndpointConnectionDelete ...
func (c PrivateEndpointConnectionsClient) PrivateEndpointConnectionDelete(ctx context.Context, id PrivateEndpointConnectionId) (result PrivateEndpointConnectionDeleteOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
		},
		HttpMethod: http.MethodDelete,
		Path:       id.ID(),
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

	result.Poller, err = resourcemanager.PollerFromResponse(resp, c.Client)
	if err != nil {
		return
	}

	return
}

// PrivateEndpointConnectionDeleteThenPoll performs PrivateEndpointConnectionDelete then polls until it's completed
func (c PrivateEndpointConnectionsClient) PrivateEndpointConnectionDeleteThenPoll(ctx context.Context, id PrivateEndpointConnectionId) error {
	result, err := c.PrivateEndpointConnectionDelete(ctx, id)
	if err != nil {
		return fmt.Errorf("performing PrivateEndpointConnectionDelete: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after PrivateEndpointConnectionDelete: %+v", err)
	}

	return nil
}
