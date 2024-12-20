package webapps

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

type RestoreFromBackupBlobSlotOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

// RestoreFromBackupBlobSlot ...
func (c WebAppsClient) RestoreFromBackupBlobSlot(ctx context.Context, id SlotId, input RestoreRequest) (result RestoreFromBackupBlobSlotOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/restoreFromBackupBlob", id.ID()),
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

// RestoreFromBackupBlobSlotThenPoll performs RestoreFromBackupBlobSlot then polls until it's completed
func (c WebAppsClient) RestoreFromBackupBlobSlotThenPoll(ctx context.Context, id SlotId, input RestoreRequest) error {
	result, err := c.RestoreFromBackupBlobSlot(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing RestoreFromBackupBlobSlot: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after RestoreFromBackupBlobSlot: %+v", err)
	}

	return nil
}
