package provisionedclusterinstances

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

type ListAdminKubeconfigOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *ListCredentialResponse
}

// ListAdminKubeconfig ...
func (c ProvisionedClusterInstancesClient) ListAdminKubeconfig(ctx context.Context, id commonids.ScopeId) (result ListAdminKubeconfigOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/providers/Microsoft.HybridContainerService/provisionedClusterInstances/default/listAdminKubeconfig", id.ID()),
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

// ListAdminKubeconfigThenPoll performs ListAdminKubeconfig then polls until it's completed
func (c ProvisionedClusterInstancesClient) ListAdminKubeconfigThenPoll(ctx context.Context, id commonids.ScopeId) error {
	result, err := c.ListAdminKubeconfig(ctx, id)
	if err != nil {
		return fmt.Errorf("performing ListAdminKubeconfig: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after ListAdminKubeconfig: %+v", err)
	}

	return nil
}
