package stable

import (
	"encoding/json"
	"fmt"

	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ItemActionStat struct {
	// The number of times the action took place. Read-only.
	ActionCount nullable.Type[int64] `json:"actionCount,omitempty"`

	// The number of distinct actors that performed the action. Read-only.
	ActorCount nullable.Type[int64] `json:"actorCount,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

var _ json.Marshaler = ItemActionStat{}

func (s ItemActionStat) MarshalJSON() ([]byte, error) {
	type wrapper ItemActionStat
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ItemActionStat: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ItemActionStat: %+v", err)
	}

	delete(decoded, "actionCount")
	delete(decoded, "actorCount")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ItemActionStat: %+v", err)
	}

	return encoded, nil
}
