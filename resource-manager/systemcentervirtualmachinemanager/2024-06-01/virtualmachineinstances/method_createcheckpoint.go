package virtualmachineinstances

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/client/pollers"
	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateCheckpointOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

// CreateCheckpoint ...
func (c VirtualMachineInstancesClient) CreateCheckpoint(ctx context.Context, id commonids.ScopeId, input VirtualMachineCreateCheckpoint) (result CreateCheckpointOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/providers/Microsoft.ScVmm/virtualMachineInstances/default/createCheckpoint", id.ID()),
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

// CreateCheckpointThenPoll performs CreateCheckpoint then polls until it's completed
func (c VirtualMachineInstancesClient) CreateCheckpointThenPoll(ctx context.Context, id commonids.ScopeId, input VirtualMachineCreateCheckpoint) error {
	result, err := c.CreateCheckpoint(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing CreateCheckpoint: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after CreateCheckpoint: %+v", err)
	}

	return nil
}
