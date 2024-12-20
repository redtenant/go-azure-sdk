package incrementalrestorepoints

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

type DiskRestorePointRevokeAccessOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

// DiskRestorePointRevokeAccess ...
func (c IncrementalRestorePointsClient) DiskRestorePointRevokeAccess(ctx context.Context, id DiskRestorePointId) (result DiskRestorePointRevokeAccessOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/endGetAccess", id.ID()),
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

// DiskRestorePointRevokeAccessThenPoll performs DiskRestorePointRevokeAccess then polls until it's completed
func (c IncrementalRestorePointsClient) DiskRestorePointRevokeAccessThenPoll(ctx context.Context, id DiskRestorePointId) error {
	result, err := c.DiskRestorePointRevokeAccess(ctx, id)
	if err != nil {
		return fmt.Errorf("performing DiskRestorePointRevokeAccess: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after DiskRestorePointRevokeAccess: %+v", err)
	}

	return nil
}
