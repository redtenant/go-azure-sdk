package grouppolicymigrationreportgrouppolicysettingmapping

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateGroupPolicyMigrationReportSettingMappingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.GroupPolicySettingMapping
}

type CreateGroupPolicyMigrationReportSettingMappingOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateGroupPolicyMigrationReportSettingMappingOperationOptions() CreateGroupPolicyMigrationReportSettingMappingOperationOptions {
	return CreateGroupPolicyMigrationReportSettingMappingOperationOptions{}
}

func (o CreateGroupPolicyMigrationReportSettingMappingOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateGroupPolicyMigrationReportSettingMappingOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateGroupPolicyMigrationReportSettingMappingOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateGroupPolicyMigrationReportSettingMapping - Create new navigation property to groupPolicySettingMappings for
// deviceManagement
func (c GroupPolicyMigrationReportGroupPolicySettingMappingClient) CreateGroupPolicyMigrationReportSettingMapping(ctx context.Context, id beta.DeviceManagementGroupPolicyMigrationReportId, input beta.GroupPolicySettingMapping, options CreateGroupPolicyMigrationReportSettingMappingOperationOptions) (result CreateGroupPolicyMigrationReportSettingMappingOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/groupPolicySettingMappings", id.ID()),
		RetryFunc:     options.RetryFunc,
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

	var model beta.GroupPolicySettingMapping
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
