package deponboardingsetting

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

type CreateDepOnboardingSettingGenerateEncryptionPublicKeyOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *CreateDepOnboardingSettingGenerateEncryptionPublicKeyResult
}

type CreateDepOnboardingSettingGenerateEncryptionPublicKeyOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateDepOnboardingSettingGenerateEncryptionPublicKeyOperationOptions() CreateDepOnboardingSettingGenerateEncryptionPublicKeyOperationOptions {
	return CreateDepOnboardingSettingGenerateEncryptionPublicKeyOperationOptions{}
}

func (o CreateDepOnboardingSettingGenerateEncryptionPublicKeyOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateDepOnboardingSettingGenerateEncryptionPublicKeyOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateDepOnboardingSettingGenerateEncryptionPublicKeyOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateDepOnboardingSettingGenerateEncryptionPublicKey - Invoke action generateEncryptionPublicKey. Generate a public
// key to use to encrypt the Apple device enrollment program token
func (c DepOnboardingSettingClient) CreateDepOnboardingSettingGenerateEncryptionPublicKey(ctx context.Context, id beta.DeviceManagementDepOnboardingSettingId, options CreateDepOnboardingSettingGenerateEncryptionPublicKeyOperationOptions) (result CreateDepOnboardingSettingGenerateEncryptionPublicKeyOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/generateEncryptionPublicKey", id.ID()),
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

	var model CreateDepOnboardingSettingGenerateEncryptionPublicKeyResult
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
