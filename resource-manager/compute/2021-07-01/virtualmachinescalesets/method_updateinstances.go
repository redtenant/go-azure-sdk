package virtualmachinescalesets

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

type UpdateInstancesOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

// UpdateInstances ...
func (c VirtualMachineScaleSetsClient) UpdateInstances(ctx context.Context, id VirtualMachineScaleSetId, input VirtualMachineScaleSetVMInstanceRequiredIDs) (result UpdateInstancesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/manualupgrade", id.ID()),
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

// UpdateInstancesThenPoll performs UpdateInstances then polls until it's completed
func (c VirtualMachineScaleSetsClient) UpdateInstancesThenPoll(ctx context.Context, id VirtualMachineScaleSetId, input VirtualMachineScaleSetVMInstanceRequiredIDs) error {
	result, err := c.UpdateInstances(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing UpdateInstances: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after UpdateInstances: %+v", err)
	}

	return nil
}
