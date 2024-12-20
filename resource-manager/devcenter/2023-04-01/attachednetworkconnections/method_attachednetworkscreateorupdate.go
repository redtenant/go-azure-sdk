package attachednetworkconnections

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

type AttachedNetworksCreateOrUpdateOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *AttachedNetworkConnection
}

// AttachedNetworksCreateOrUpdate ...
func (c AttachedNetworkConnectionsClient) AttachedNetworksCreateOrUpdate(ctx context.Context, id DevCenterAttachedNetworkId, input AttachedNetworkConnection) (result AttachedNetworksCreateOrUpdateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod: http.MethodPut,
		Path:       id.ID(),
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

	result.Poller, err = resourcemanager.PollerFromResponse(resp, c.Client)
	if err != nil {
		return
	}

	return
}

// AttachedNetworksCreateOrUpdateThenPoll performs AttachedNetworksCreateOrUpdate then polls until it's completed
func (c AttachedNetworkConnectionsClient) AttachedNetworksCreateOrUpdateThenPoll(ctx context.Context, id DevCenterAttachedNetworkId, input AttachedNetworkConnection) error {
	result, err := c.AttachedNetworksCreateOrUpdate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing AttachedNetworksCreateOrUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after AttachedNetworksCreateOrUpdate: %+v", err)
	}

	return nil
}
