package permissionsanalyticgcppermissionscreepindexdistribution

import (
	"context"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreatePermissionsAnalyticGcpPermissionsCreepIndexDistributionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.PermissionsCreepIndexDistribution
}

type CreatePermissionsAnalyticGcpPermissionsCreepIndexDistributionOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreatePermissionsAnalyticGcpPermissionsCreepIndexDistributionOperationOptions() CreatePermissionsAnalyticGcpPermissionsCreepIndexDistributionOperationOptions {
	return CreatePermissionsAnalyticGcpPermissionsCreepIndexDistributionOperationOptions{}
}

func (o CreatePermissionsAnalyticGcpPermissionsCreepIndexDistributionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreatePermissionsAnalyticGcpPermissionsCreepIndexDistributionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreatePermissionsAnalyticGcpPermissionsCreepIndexDistributionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreatePermissionsAnalyticGcpPermissionsCreepIndexDistribution - Create new navigation property to
// permissionsCreepIndexDistributions for identityGovernance
func (c PermissionsAnalyticGcpPermissionsCreepIndexDistributionClient) CreatePermissionsAnalyticGcpPermissionsCreepIndexDistribution(ctx context.Context, input beta.PermissionsCreepIndexDistribution, options CreatePermissionsAnalyticGcpPermissionsCreepIndexDistributionOperationOptions) (result CreatePermissionsAnalyticGcpPermissionsCreepIndexDistributionOperationResponse, err error) {
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
		Path:          "/identityGovernance/permissionsAnalytics/gcp/permissionsCreepIndexDistributions",
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

	var model beta.PermissionsCreepIndexDistribution
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
