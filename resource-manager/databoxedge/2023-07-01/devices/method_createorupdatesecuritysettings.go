package devices

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

type CreateOrUpdateSecuritySettingsOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

// CreateOrUpdateSecuritySettings ...
func (c DevicesClient) CreateOrUpdateSecuritySettings(ctx context.Context, id DataBoxEdgeDeviceId, input SecuritySettings) (result CreateOrUpdateSecuritySettingsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/securitySettings/default/update", id.ID()),
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

// CreateOrUpdateSecuritySettingsThenPoll performs CreateOrUpdateSecuritySettings then polls until it's completed
func (c DevicesClient) CreateOrUpdateSecuritySettingsThenPoll(ctx context.Context, id DataBoxEdgeDeviceId, input SecuritySettings) error {
	result, err := c.CreateOrUpdateSecuritySettings(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing CreateOrUpdateSecuritySettings: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after CreateOrUpdateSecuritySettings: %+v", err)
	}

	return nil
}
