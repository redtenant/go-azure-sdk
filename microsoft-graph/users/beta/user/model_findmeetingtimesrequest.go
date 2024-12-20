package user

import (
	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FindMeetingTimesRequest struct {
	Attendees               *[]beta.AttendeeBase     `json:"attendees,omitempty"`
	IsOrganizerOptional     nullable.Type[bool]      `json:"isOrganizerOptional,omitempty"`
	LocationConstraint      *beta.LocationConstraint `json:"locationConstraint,omitempty"`
	MaxCandidates           nullable.Type[int64]     `json:"maxCandidates,omitempty"`
	MeetingDuration         nullable.Type[string]    `json:"meetingDuration,omitempty"`
	ReturnSuggestionReasons nullable.Type[bool]      `json:"returnSuggestionReasons,omitempty"`
	TimeConstraint          *beta.TimeConstraint     `json:"timeConstraint,omitempty"`
}
