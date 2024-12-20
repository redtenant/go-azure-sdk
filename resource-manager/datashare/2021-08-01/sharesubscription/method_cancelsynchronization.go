package sharesubscription

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

type CancelSynchronizationOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *ShareSubscriptionSynchronization
}

// CancelSynchronization ...
func (c ShareSubscriptionClient) CancelSynchronization(ctx context.Context, id ShareSubscriptionId, input ShareSubscriptionSynchronization) (result CancelSynchronizationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/cancelSynchronization", id.ID()),
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

// CancelSynchronizationThenPoll performs CancelSynchronization then polls until it's completed
func (c ShareSubscriptionClient) CancelSynchronizationThenPoll(ctx context.Context, id ShareSubscriptionId, input ShareSubscriptionSynchronization) error {
	result, err := c.CancelSynchronization(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing CancelSynchronization: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after CancelSynchronization: %+v", err)
	}

	return nil
}
