package rbacs

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

type SqlResourcesCreateUpdateSqlRoleDefinitionOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *SqlRoleDefinitionGetResults
}

// SqlResourcesCreateUpdateSqlRoleDefinition ...
func (c RbacsClient) SqlResourcesCreateUpdateSqlRoleDefinition(ctx context.Context, id SqlRoleDefinitionId, input SqlRoleDefinitionCreateUpdateParameters) (result SqlResourcesCreateUpdateSqlRoleDefinitionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod: http.MethodPut,
		Path:       id.ID(),
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

// SqlResourcesCreateUpdateSqlRoleDefinitionThenPoll performs SqlResourcesCreateUpdateSqlRoleDefinition then polls until it's completed
func (c RbacsClient) SqlResourcesCreateUpdateSqlRoleDefinitionThenPoll(ctx context.Context, id SqlRoleDefinitionId, input SqlRoleDefinitionCreateUpdateParameters) error {
	result, err := c.SqlResourcesCreateUpdateSqlRoleDefinition(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing SqlResourcesCreateUpdateSqlRoleDefinition: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after SqlResourcesCreateUpdateSqlRoleDefinition: %+v", err)
	}

	return nil
}
