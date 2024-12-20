package remotedesktopsecurityconfigurationtargetdevicegroup

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

type CreateRemoteDesktopSecurityConfigurationTargetDeviceGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.TargetDeviceGroup
}

type CreateRemoteDesktopSecurityConfigurationTargetDeviceGroupOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateRemoteDesktopSecurityConfigurationTargetDeviceGroupOperationOptions() CreateRemoteDesktopSecurityConfigurationTargetDeviceGroupOperationOptions {
	return CreateRemoteDesktopSecurityConfigurationTargetDeviceGroupOperationOptions{}
}

func (o CreateRemoteDesktopSecurityConfigurationTargetDeviceGroupOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateRemoteDesktopSecurityConfigurationTargetDeviceGroupOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateRemoteDesktopSecurityConfigurationTargetDeviceGroupOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateRemoteDesktopSecurityConfigurationTargetDeviceGroup - Create targetDeviceGroup. Create a new targetDeviceGroup
// object for the remoteDesktopSecurityConfiguration object on the servicePrincipal. You can configure a maximum of 10
// target device groups for the remoteDesktopSecurityConfiguraiton object on the servicePrincipal.
func (c RemoteDesktopSecurityConfigurationTargetDeviceGroupClient) CreateRemoteDesktopSecurityConfigurationTargetDeviceGroup(ctx context.Context, id beta.ServicePrincipalId, input beta.TargetDeviceGroup, options CreateRemoteDesktopSecurityConfigurationTargetDeviceGroupOperationOptions) (result CreateRemoteDesktopSecurityConfigurationTargetDeviceGroupOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/remoteDesktopSecurityConfiguration/targetDeviceGroups", id.ID()),
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

	var model beta.TargetDeviceGroup
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
