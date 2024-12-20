package synchronizationjob

import (
	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ValidateSynchronizationJobsCredentialsRequest struct {
	ApplicationIdentifier nullable.Type[string]                           `json:"applicationIdentifier,omitempty"`
	Credentials           *[]beta.SynchronizationSecretKeyStringValuePair `json:"credentials,omitempty"`
	TemplateId            nullable.Type[string]                           `json:"templateId,omitempty"`
	UseSavedCredentials   nullable.Type[bool]                             `json:"useSavedCredentials,omitempty"`
}
