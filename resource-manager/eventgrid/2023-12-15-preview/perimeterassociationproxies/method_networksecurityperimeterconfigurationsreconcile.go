package perimeterassociationproxies

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

type NetworkSecurityPerimeterConfigurationsReconcileOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *NetworkSecurityPerimeterConfiguration
}

// NetworkSecurityPerimeterConfigurationsReconcile ...
func (c PerimeterAssociationProxiesClient) NetworkSecurityPerimeterConfigurationsReconcile(ctx context.Context, id ScopedNetworkSecurityPerimeterConfigurationId) (result NetworkSecurityPerimeterConfigurationsReconcileOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/reconcile", id.ID()),
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

// NetworkSecurityPerimeterConfigurationsReconcileThenPoll performs NetworkSecurityPerimeterConfigurationsReconcile then polls until it's completed
func (c PerimeterAssociationProxiesClient) NetworkSecurityPerimeterConfigurationsReconcileThenPoll(ctx context.Context, id ScopedNetworkSecurityPerimeterConfigurationId) error {
	result, err := c.NetworkSecurityPerimeterConfigurationsReconcile(ctx, id)
	if err != nil {
		return fmt.Errorf("performing NetworkSecurityPerimeterConfigurationsReconcile: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after NetworkSecurityPerimeterConfigurationsReconcile: %+v", err)
	}

	return nil
}
