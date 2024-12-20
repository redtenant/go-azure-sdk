package integrationruntime

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

type EnableInteractiveQueryOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

// EnableInteractiveQuery ...
func (c IntegrationRuntimeClient) EnableInteractiveQuery(ctx context.Context, id IntegrationRuntimeId) (result EnableInteractiveQueryOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/enableInteractiveQuery", id.ID()),
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

// EnableInteractiveQueryThenPoll performs EnableInteractiveQuery then polls until it's completed
func (c IntegrationRuntimeClient) EnableInteractiveQueryThenPoll(ctx context.Context, id IntegrationRuntimeId) error {
	result, err := c.EnableInteractiveQuery(ctx, id)
	if err != nil {
		return fmt.Errorf("performing EnableInteractiveQuery: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after EnableInteractiveQuery: %+v", err)
	}

	return nil
}
