package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedMobileLobApp interface {
	Entity
	MobileApp
	ManagedApp
	ManagedMobileLobApp() BaseManagedMobileLobAppImpl
}

var _ ManagedMobileLobApp = BaseManagedMobileLobAppImpl{}

type BaseManagedMobileLobAppImpl struct {
	// The internal committed content version.
	CommittedContentVersion nullable.Type[string] `json:"committedContentVersion,omitempty"`

	// The list of content versions for this app.
	ContentVersions *[]MobileAppContent `json:"contentVersions,omitempty"`

	// The name of the main Lob application file.
	FileName nullable.Type[string] `json:"fileName,omitempty"`

	// The total size, including all uploaded files.
	Size *int64 `json:"size,omitempty"`

	// Fields inherited from ManagedApp

	// A managed (MAM) application's availability.
	AppAvailability *ManagedAppAvailability `json:"appAvailability,omitempty"`

	// The Application's version.
	Version nullable.Type[string] `json:"version,omitempty"`

	// Fields inherited from MobileApp

	// The list of group assignments for this mobile app.
	Assignments *[]MobileAppAssignment `json:"assignments,omitempty"`

	// The list of categories for this app.
	Categories *[]MobileAppCategory `json:"categories,omitempty"`

	// The date and time the app was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The description of the app.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The developer of the app.
	Developer nullable.Type[string] `json:"developer,omitempty"`

	// The admin provided or imported title of the app.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The more information Url.
	InformationUrl nullable.Type[string] `json:"informationUrl,omitempty"`

	// The value indicating whether the app is marked as featured by the admin.
	IsFeatured *bool `json:"isFeatured,omitempty"`

	// The large icon, to be displayed in the app details and used for upload of the icon.
	LargeIcon *MimeContent `json:"largeIcon,omitempty"`

	// The date and time the app was last modified.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// Notes for the app.
	Notes nullable.Type[string] `json:"notes,omitempty"`

	// The owner of the app.
	Owner nullable.Type[string] `json:"owner,omitempty"`

	// The privacy statement Url.
	PrivacyInformationUrl nullable.Type[string] `json:"privacyInformationUrl,omitempty"`

	// The publisher of the app.
	Publisher nullable.Type[string] `json:"publisher,omitempty"`

	// Indicates the publishing state of an app.
	PublishingState *MobileAppPublishingState `json:"publishingState,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseManagedMobileLobAppImpl) ManagedMobileLobApp() BaseManagedMobileLobAppImpl {
	return s
}

func (s BaseManagedMobileLobAppImpl) ManagedApp() BaseManagedAppImpl {
	return BaseManagedAppImpl{
		AppAvailability:       s.AppAvailability,
		Version:               s.Version,
		Assignments:           s.Assignments,
		Categories:            s.Categories,
		CreatedDateTime:       s.CreatedDateTime,
		Description:           s.Description,
		Developer:             s.Developer,
		DisplayName:           s.DisplayName,
		InformationUrl:        s.InformationUrl,
		IsFeatured:            s.IsFeatured,
		LargeIcon:             s.LargeIcon,
		LastModifiedDateTime:  s.LastModifiedDateTime,
		Notes:                 s.Notes,
		Owner:                 s.Owner,
		PrivacyInformationUrl: s.PrivacyInformationUrl,
		Publisher:             s.Publisher,
		PublishingState:       s.PublishingState,
		Id:                    s.Id,
		ODataId:               s.ODataId,
		ODataType:             s.ODataType,
	}
}

func (s BaseManagedMobileLobAppImpl) MobileApp() BaseMobileAppImpl {
	return BaseMobileAppImpl{
		Assignments:           s.Assignments,
		Categories:            s.Categories,
		CreatedDateTime:       s.CreatedDateTime,
		Description:           s.Description,
		Developer:             s.Developer,
		DisplayName:           s.DisplayName,
		InformationUrl:        s.InformationUrl,
		IsFeatured:            s.IsFeatured,
		LargeIcon:             s.LargeIcon,
		LastModifiedDateTime:  s.LastModifiedDateTime,
		Notes:                 s.Notes,
		Owner:                 s.Owner,
		PrivacyInformationUrl: s.PrivacyInformationUrl,
		Publisher:             s.Publisher,
		PublishingState:       s.PublishingState,
		Id:                    s.Id,
		ODataId:               s.ODataId,
		ODataType:             s.ODataType,
	}
}

func (s BaseManagedMobileLobAppImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ ManagedMobileLobApp = RawManagedMobileLobAppImpl{}

// RawManagedMobileLobAppImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawManagedMobileLobAppImpl struct {
	managedMobileLobApp BaseManagedMobileLobAppImpl
	Type                string
	Values              map[string]interface{}
}

func (s RawManagedMobileLobAppImpl) ManagedMobileLobApp() BaseManagedMobileLobAppImpl {
	return s.managedMobileLobApp
}

func (s RawManagedMobileLobAppImpl) ManagedApp() BaseManagedAppImpl {
	return s.managedMobileLobApp.ManagedApp()
}

func (s RawManagedMobileLobAppImpl) MobileApp() BaseMobileAppImpl {
	return s.managedMobileLobApp.MobileApp()
}

func (s RawManagedMobileLobAppImpl) Entity() BaseEntityImpl {
	return s.managedMobileLobApp.Entity()
}

var _ json.Marshaler = BaseManagedMobileLobAppImpl{}

func (s BaseManagedMobileLobAppImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseManagedMobileLobAppImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseManagedMobileLobAppImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseManagedMobileLobAppImpl: %+v", err)
	}

	delete(decoded, "size")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.managedMobileLobApp"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseManagedMobileLobAppImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalManagedMobileLobAppImplementation(input []byte) (ManagedMobileLobApp, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedMobileLobApp into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.managedAndroidLobApp") {
		var out ManagedAndroidLobApp
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedAndroidLobApp: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.managedIOSLobApp") {
		var out ManagedIOSLobApp
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedIOSLobApp: %+v", err)
		}
		return out, nil
	}

	var parent BaseManagedMobileLobAppImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseManagedMobileLobAppImpl: %+v", err)
	}

	return RawManagedMobileLobAppImpl{
		managedMobileLobApp: parent,
		Type:                value,
		Values:              temp,
	}, nil

}
