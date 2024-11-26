package settingshiftpreference

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

type UpdateSettingShiftPreferenceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateSettingShiftPreferenceOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateSettingShiftPreferenceOperationOptions() UpdateSettingShiftPreferenceOperationOptions {
	return UpdateSettingShiftPreferenceOperationOptions{}
}

func (o UpdateSettingShiftPreferenceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateSettingShiftPreferenceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateSettingShiftPreferenceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateSettingShiftPreference - Update shiftPreferences. Update the properties and relationships of a shiftPreferences
// object.
func (c SettingShiftPreferenceClient) UpdateSettingShiftPreference(ctx context.Context, id stable.UserId, input stable.ShiftPreferences, options UpdateSettingShiftPreferenceOperationOptions) (result UpdateSettingShiftPreferenceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/settings/shiftPreferences", id.ID()),
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
