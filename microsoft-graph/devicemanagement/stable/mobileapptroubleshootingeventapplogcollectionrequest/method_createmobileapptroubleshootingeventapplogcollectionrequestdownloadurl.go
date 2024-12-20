package mobileapptroubleshootingeventapplogcollectionrequest

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateMobileAppTroubleshootingEventAppLogCollectionRequestDownloadUrlOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.AppLogCollectionDownloadDetails
}

type CreateMobileAppTroubleshootingEventAppLogCollectionRequestDownloadUrlOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateMobileAppTroubleshootingEventAppLogCollectionRequestDownloadUrlOperationOptions() CreateMobileAppTroubleshootingEventAppLogCollectionRequestDownloadUrlOperationOptions {
	return CreateMobileAppTroubleshootingEventAppLogCollectionRequestDownloadUrlOperationOptions{}
}

func (o CreateMobileAppTroubleshootingEventAppLogCollectionRequestDownloadUrlOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateMobileAppTroubleshootingEventAppLogCollectionRequestDownloadUrlOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateMobileAppTroubleshootingEventAppLogCollectionRequestDownloadUrlOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateMobileAppTroubleshootingEventAppLogCollectionRequestDownloadUrl - Invoke action createDownloadUrl. Not yet
// documented
func (c MobileAppTroubleshootingEventAppLogCollectionRequestClient) CreateMobileAppTroubleshootingEventAppLogCollectionRequestDownloadUrl(ctx context.Context, id stable.DeviceManagementMobileAppTroubleshootingEventIdAppLogCollectionRequestId, options CreateMobileAppTroubleshootingEventAppLogCollectionRequestDownloadUrlOperationOptions) (result CreateMobileAppTroubleshootingEventAppLogCollectionRequestDownloadUrlOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/createDownloadUrl", id.ID()),
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

	var model stable.AppLogCollectionDownloadDetails
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
