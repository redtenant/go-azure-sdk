package eventsubscriptions

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

type SystemTopicEventSubscriptionsDeleteOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

// SystemTopicEventSubscriptionsDelete ...
func (c EventSubscriptionsClient) SystemTopicEventSubscriptionsDelete(ctx context.Context, id SystemTopicEventSubscriptionId) (result SystemTopicEventSubscriptionsDeleteOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
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

// SystemTopicEventSubscriptionsDeleteThenPoll performs SystemTopicEventSubscriptionsDelete then polls until it's completed
func (c EventSubscriptionsClient) SystemTopicEventSubscriptionsDeleteThenPoll(ctx context.Context, id SystemTopicEventSubscriptionId) error {
	result, err := c.SystemTopicEventSubscriptionsDelete(ctx, id)
	if err != nil {
		return fmt.Errorf("performing SystemTopicEventSubscriptionsDelete: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after SystemTopicEventSubscriptionsDelete: %+v", err)
	}

	return nil
}
