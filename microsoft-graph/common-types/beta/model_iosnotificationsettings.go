package beta

import (
	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosNotificationSettings struct {
	// Notification Settings Alert Type.
	AlertType *IosNotificationAlertType `json:"alertType,omitempty"`

	// Application name to be associated with the bundleID.
	AppName nullable.Type[string] `json:"appName,omitempty"`

	// Indicates whether badges are allowed for this app.
	BadgesEnabled nullable.Type[bool] `json:"badgesEnabled,omitempty"`

	// Bundle id of app to which to apply these notification settings.
	BundleId *string `json:"bundleID,omitempty"`

	// Indicates whether notifications are allowed for this app.
	Enabled nullable.Type[bool] `json:"enabled,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Determines when notification previews are visible on an iOS device. Previews can include things like text (from
	// Messages and Mail) and invitation details (from Calendar). When configured, it will override the user's defined
	// preview settings.
	PreviewVisibility *IosNotificationPreviewVisibility `json:"previewVisibility,omitempty"`

	// Publisher to be associated with the bundleID.
	Publisher nullable.Type[string] `json:"publisher,omitempty"`

	// Indicates whether notifications can be shown in notification center.
	ShowInNotificationCenter nullable.Type[bool] `json:"showInNotificationCenter,omitempty"`

	// Indicates whether notifications can be shown on the lock screen.
	ShowOnLockScreen nullable.Type[bool] `json:"showOnLockScreen,omitempty"`

	// Indicates whether sounds are allowed for this app.
	SoundsEnabled nullable.Type[bool] `json:"soundsEnabled,omitempty"`
}
