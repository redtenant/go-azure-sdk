package informationprotectionsensitivitylabelsublabel

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

type CreateInformationProtectionSensitivityLabelSublabelOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.SensitivityLabel
}

type CreateInformationProtectionSensitivityLabelSublabelOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateInformationProtectionSensitivityLabelSublabelOperationOptions() CreateInformationProtectionSensitivityLabelSublabelOperationOptions {
	return CreateInformationProtectionSensitivityLabelSublabelOperationOptions{}
}

func (o CreateInformationProtectionSensitivityLabelSublabelOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateInformationProtectionSensitivityLabelSublabelOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateInformationProtectionSensitivityLabelSublabelOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateInformationProtectionSensitivityLabelSublabel - Create new navigation property to sublabels for users
func (c InformationProtectionSensitivityLabelSublabelClient) CreateInformationProtectionSensitivityLabelSublabel(ctx context.Context, id beta.UserIdInformationProtectionSensitivityLabelId, input beta.SensitivityLabel, options CreateInformationProtectionSensitivityLabelSublabelOperationOptions) (result CreateInformationProtectionSensitivityLabelSublabelOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/sublabels", id.ID()),
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

	var model beta.SensitivityLabel
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
