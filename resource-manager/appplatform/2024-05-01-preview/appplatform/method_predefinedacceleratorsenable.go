package appplatform

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

type PredefinedAcceleratorsEnableOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

// PredefinedAcceleratorsEnable ...
func (c AppPlatformClient) PredefinedAcceleratorsEnable(ctx context.Context, id PredefinedAcceleratorId) (result PredefinedAcceleratorsEnableOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/enable", id.ID()),
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

// PredefinedAcceleratorsEnableThenPoll performs PredefinedAcceleratorsEnable then polls until it's completed
func (c AppPlatformClient) PredefinedAcceleratorsEnableThenPoll(ctx context.Context, id PredefinedAcceleratorId) error {
	result, err := c.PredefinedAcceleratorsEnable(ctx, id)
	if err != nil {
		return fmt.Errorf("performing PredefinedAcceleratorsEnable: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after PredefinedAcceleratorsEnable: %+v", err)
	}

	return nil
}
