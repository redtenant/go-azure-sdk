package comanageddevice

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

type SendComanagedDeviceCustomNotificationToCompanyPortalOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SendComanagedDeviceCustomNotificationToCompanyPortalOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultSendComanagedDeviceCustomNotificationToCompanyPortalOperationOptions() SendComanagedDeviceCustomNotificationToCompanyPortalOperationOptions {
	return SendComanagedDeviceCustomNotificationToCompanyPortalOperationOptions{}
}

func (o SendComanagedDeviceCustomNotificationToCompanyPortalOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SendComanagedDeviceCustomNotificationToCompanyPortalOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SendComanagedDeviceCustomNotificationToCompanyPortalOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SendComanagedDeviceCustomNotificationToCompanyPortal - Invoke action sendCustomNotificationToCompanyPortal
func (c ComanagedDeviceClient) SendComanagedDeviceCustomNotificationToCompanyPortal(ctx context.Context, id beta.DeviceManagementComanagedDeviceId, input SendComanagedDeviceCustomNotificationToCompanyPortalRequest, options SendComanagedDeviceCustomNotificationToCompanyPortalOperationOptions) (result SendComanagedDeviceCustomNotificationToCompanyPortalOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/sendCustomNotificationToCompanyPortal", id.ID()),
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

	return
}
