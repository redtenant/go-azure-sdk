package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityB2cUserFlowIdLanguageIdDefaultPageId{}

// IdentityB2cUserFlowIdLanguageIdDefaultPageId is a struct representing the Resource ID for a Identity B 2c User Flow Id Language Id Default Page
type IdentityB2cUserFlowIdLanguageIdDefaultPageId struct {
	B2cIdentityUserFlowId           string
	UserFlowLanguageConfigurationId string
	UserFlowLanguagePageId          string
}

// NewIdentityB2cUserFlowIdLanguageIdDefaultPageID returns a new IdentityB2cUserFlowIdLanguageIdDefaultPageId struct
func NewIdentityB2cUserFlowIdLanguageIdDefaultPageID(b2cIdentityUserFlowId string, userFlowLanguageConfigurationId string, userFlowLanguagePageId string) IdentityB2cUserFlowIdLanguageIdDefaultPageId {
	return IdentityB2cUserFlowIdLanguageIdDefaultPageId{
		B2cIdentityUserFlowId:           b2cIdentityUserFlowId,
		UserFlowLanguageConfigurationId: userFlowLanguageConfigurationId,
		UserFlowLanguagePageId:          userFlowLanguagePageId,
	}
}

// ParseIdentityB2cUserFlowIdLanguageIdDefaultPageID parses 'input' into a IdentityB2cUserFlowIdLanguageIdDefaultPageId
func ParseIdentityB2cUserFlowIdLanguageIdDefaultPageID(input string) (*IdentityB2cUserFlowIdLanguageIdDefaultPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityB2cUserFlowIdLanguageIdDefaultPageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityB2cUserFlowIdLanguageIdDefaultPageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityB2cUserFlowIdLanguageIdDefaultPageIDInsensitively parses 'input' case-insensitively into a IdentityB2cUserFlowIdLanguageIdDefaultPageId
// note: this method should only be used for API response data and not user input
func ParseIdentityB2cUserFlowIdLanguageIdDefaultPageIDInsensitively(input string) (*IdentityB2cUserFlowIdLanguageIdDefaultPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityB2cUserFlowIdLanguageIdDefaultPageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityB2cUserFlowIdLanguageIdDefaultPageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityB2cUserFlowIdLanguageIdDefaultPageId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.B2cIdentityUserFlowId, ok = input.Parsed["b2cIdentityUserFlowId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "b2cIdentityUserFlowId", input)
	}

	if id.UserFlowLanguageConfigurationId, ok = input.Parsed["userFlowLanguageConfigurationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userFlowLanguageConfigurationId", input)
	}

	if id.UserFlowLanguagePageId, ok = input.Parsed["userFlowLanguagePageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userFlowLanguagePageId", input)
	}

	return nil
}

// ValidateIdentityB2cUserFlowIdLanguageIdDefaultPageID checks that 'input' can be parsed as a Identity B 2c User Flow Id Language Id Default Page ID
func ValidateIdentityB2cUserFlowIdLanguageIdDefaultPageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityB2cUserFlowIdLanguageIdDefaultPageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity B 2c User Flow Id Language Id Default Page ID
func (id IdentityB2cUserFlowIdLanguageIdDefaultPageId) ID() string {
	fmtString := "/identity/b2cUserFlows/%s/languages/%s/defaultPages/%s"
	return fmt.Sprintf(fmtString, id.B2cIdentityUserFlowId, id.UserFlowLanguageConfigurationId, id.UserFlowLanguagePageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity B 2c User Flow Id Language Id Default Page ID
func (id IdentityB2cUserFlowIdLanguageIdDefaultPageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identity", "identity", "identity"),
		resourceids.StaticSegment("b2cUserFlows", "b2cUserFlows", "b2cUserFlows"),
		resourceids.UserSpecifiedSegment("b2cIdentityUserFlowId", "b2cIdentityUserFlowId"),
		resourceids.StaticSegment("languages", "languages", "languages"),
		resourceids.UserSpecifiedSegment("userFlowLanguageConfigurationId", "userFlowLanguageConfigurationId"),
		resourceids.StaticSegment("defaultPages", "defaultPages", "defaultPages"),
		resourceids.UserSpecifiedSegment("userFlowLanguagePageId", "userFlowLanguagePageId"),
	}
}

// String returns a human-readable description of this Identity B 2c User Flow Id Language Id Default Page ID
func (id IdentityB2cUserFlowIdLanguageIdDefaultPageId) String() string {
	components := []string{
		fmt.Sprintf("B 2c Identity User Flow: %q", id.B2cIdentityUserFlowId),
		fmt.Sprintf("User Flow Language Configuration: %q", id.UserFlowLanguageConfigurationId),
		fmt.Sprintf("User Flow Language Page: %q", id.UserFlowLanguagePageId),
	}
	return fmt.Sprintf("Identity B 2c User Flow Id Language Id Default Page (%s)", strings.Join(components, "\n"))
}
