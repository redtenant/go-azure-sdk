package androidforworkappconfigurationschema

import (
	"context"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetAndroidForWorkAppConfigurationSchemaOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.AndroidForWorkAppConfigurationSchema
}

type GetAndroidForWorkAppConfigurationSchemaOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetAndroidForWorkAppConfigurationSchemaOperationOptions() GetAndroidForWorkAppConfigurationSchemaOperationOptions {
	return GetAndroidForWorkAppConfigurationSchemaOperationOptions{}
}

func (o GetAndroidForWorkAppConfigurationSchemaOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetAndroidForWorkAppConfigurationSchemaOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetAndroidForWorkAppConfigurationSchemaOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetAndroidForWorkAppConfigurationSchema - Get androidForWorkAppConfigurationSchemas from deviceManagement. Android
// for Work app configuration schema entities.
func (c AndroidForWorkAppConfigurationSchemaClient) GetAndroidForWorkAppConfigurationSchema(ctx context.Context, id beta.DeviceManagementAndroidForWorkAppConfigurationSchemaId, options GetAndroidForWorkAppConfigurationSchemaOperationOptions) (result GetAndroidForWorkAppConfigurationSchemaOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
		RetryFunc:     options.RetryFunc,
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

	var model beta.AndroidForWorkAppConfigurationSchema
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
