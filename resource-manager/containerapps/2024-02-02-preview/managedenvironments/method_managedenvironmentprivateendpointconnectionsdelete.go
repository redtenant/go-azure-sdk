package managedenvironments

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

type ManagedEnvironmentPrivateEndpointConnectionsDeleteOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

// ManagedEnvironmentPrivateEndpointConnectionsDelete ...
func (c ManagedEnvironmentsClient) ManagedEnvironmentPrivateEndpointConnectionsDelete(ctx context.Context, id PrivateEndpointConnectionId) (result ManagedEnvironmentPrivateEndpointConnectionsDeleteOperationResponse, err error) {
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

// ManagedEnvironmentPrivateEndpointConnectionsDeleteThenPoll performs ManagedEnvironmentPrivateEndpointConnectionsDelete then polls until it's completed
func (c ManagedEnvironmentsClient) ManagedEnvironmentPrivateEndpointConnectionsDeleteThenPoll(ctx context.Context, id PrivateEndpointConnectionId) error {
	result, err := c.ManagedEnvironmentPrivateEndpointConnectionsDelete(ctx, id)
	if err != nil {
		return fmt.Errorf("performing ManagedEnvironmentPrivateEndpointConnectionsDelete: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after ManagedEnvironmentPrivateEndpointConnectionsDelete: %+v", err)
	}

	return nil
}
