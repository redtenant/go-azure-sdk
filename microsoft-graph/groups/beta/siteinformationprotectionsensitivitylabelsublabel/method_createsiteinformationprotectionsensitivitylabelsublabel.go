package siteinformationprotectionsensitivitylabelsublabel

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

type CreateSiteInformationProtectionSensitivityLabelSublabelOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.SensitivityLabel
}

type CreateSiteInformationProtectionSensitivityLabelSublabelOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateSiteInformationProtectionSensitivityLabelSublabelOperationOptions() CreateSiteInformationProtectionSensitivityLabelSublabelOperationOptions {
	return CreateSiteInformationProtectionSensitivityLabelSublabelOperationOptions{}
}

func (o CreateSiteInformationProtectionSensitivityLabelSublabelOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateSiteInformationProtectionSensitivityLabelSublabelOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateSiteInformationProtectionSensitivityLabelSublabelOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateSiteInformationProtectionSensitivityLabelSublabel - Create new navigation property to sublabels for groups
func (c SiteInformationProtectionSensitivityLabelSublabelClient) CreateSiteInformationProtectionSensitivityLabelSublabel(ctx context.Context, id beta.GroupIdSiteIdInformationProtectionSensitivityLabelId, input beta.SensitivityLabel, options CreateSiteInformationProtectionSensitivityLabelSublabelOperationOptions) (result CreateSiteInformationProtectionSensitivityLabelSublabelOperationResponse, err error) {
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
