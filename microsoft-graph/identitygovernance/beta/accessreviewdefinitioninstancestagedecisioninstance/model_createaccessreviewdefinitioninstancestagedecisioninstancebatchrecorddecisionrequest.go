package accessreviewdefinitioninstancestagedecisioninstance

import (
	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateAccessReviewDefinitionInstanceStageDecisionInstanceBatchRecordDecisionRequest struct {
	Decision      nullable.Type[string] `json:"decision,omitempty"`
	Justification nullable.Type[string] `json:"justification,omitempty"`
	PrincipalId   nullable.Type[string] `json:"principalId,omitempty"`
	ResourceId    nullable.Type[string] `json:"resourceId,omitempty"`
}
