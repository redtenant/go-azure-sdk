package androiddeviceownerenrollmentprofile

import (
	"context"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateAndroidDeviceOwnerEnrollmentProfileOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.AndroidDeviceOwnerEnrollmentProfile
}

type CreateAndroidDeviceOwnerEnrollmentProfileOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateAndroidDeviceOwnerEnrollmentProfileOperationOptions() CreateAndroidDeviceOwnerEnrollmentProfileOperationOptions {
	return CreateAndroidDeviceOwnerEnrollmentProfileOperationOptions{}
}

func (o CreateAndroidDeviceOwnerEnrollmentProfileOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateAndroidDeviceOwnerEnrollmentProfileOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateAndroidDeviceOwnerEnrollmentProfileOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateAndroidDeviceOwnerEnrollmentProfile - Create new navigation property to androidDeviceOwnerEnrollmentProfiles
// for deviceManagement
func (c AndroidDeviceOwnerEnrollmentProfileClient) CreateAndroidDeviceOwnerEnrollmentProfile(ctx context.Context, input beta.AndroidDeviceOwnerEnrollmentProfile, options CreateAndroidDeviceOwnerEnrollmentProfileOperationOptions) (result CreateAndroidDeviceOwnerEnrollmentProfileOperationResponse, err error) {
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
		Path:          "/deviceManagement/androidDeviceOwnerEnrollmentProfiles",
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

	var model beta.AndroidDeviceOwnerEnrollmentProfile
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
