package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityB2cUserFlowIdLanguageId{}

// IdentityB2cUserFlowIdLanguageId is a struct representing the Resource ID for a Identity B 2c User Flow Id Language
type IdentityB2cUserFlowIdLanguageId struct {
	B2cIdentityUserFlowId           string
	UserFlowLanguageConfigurationId string
}

// NewIdentityB2cUserFlowIdLanguageID returns a new IdentityB2cUserFlowIdLanguageId struct
func NewIdentityB2cUserFlowIdLanguageID(b2cIdentityUserFlowId string, userFlowLanguageConfigurationId string) IdentityB2cUserFlowIdLanguageId {
	return IdentityB2cUserFlowIdLanguageId{
		B2cIdentityUserFlowId:           b2cIdentityUserFlowId,
		UserFlowLanguageConfigurationId: userFlowLanguageConfigurationId,
	}
}

// ParseIdentityB2cUserFlowIdLanguageID parses 'input' into a IdentityB2cUserFlowIdLanguageId
func ParseIdentityB2cUserFlowIdLanguageID(input string) (*IdentityB2cUserFlowIdLanguageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityB2cUserFlowIdLanguageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityB2cUserFlowIdLanguageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityB2cUserFlowIdLanguageIDInsensitively parses 'input' case-insensitively into a IdentityB2cUserFlowIdLanguageId
// note: this method should only be used for API response data and not user input
func ParseIdentityB2cUserFlowIdLanguageIDInsensitively(input string) (*IdentityB2cUserFlowIdLanguageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityB2cUserFlowIdLanguageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityB2cUserFlowIdLanguageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityB2cUserFlowIdLanguageId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.B2cIdentityUserFlowId, ok = input.Parsed["b2cIdentityUserFlowId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "b2cIdentityUserFlowId", input)
	}

	if id.UserFlowLanguageConfigurationId, ok = input.Parsed["userFlowLanguageConfigurationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userFlowLanguageConfigurationId", input)
	}

	return nil
}

// ValidateIdentityB2cUserFlowIdLanguageID checks that 'input' can be parsed as a Identity B 2c User Flow Id Language ID
func ValidateIdentityB2cUserFlowIdLanguageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityB2cUserFlowIdLanguageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity B 2c User Flow Id Language ID
func (id IdentityB2cUserFlowIdLanguageId) ID() string {
	fmtString := "/identity/b2cUserFlows/%s/languages/%s"
	return fmt.Sprintf(fmtString, id.B2cIdentityUserFlowId, id.UserFlowLanguageConfigurationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity B 2c User Flow Id Language ID
func (id IdentityB2cUserFlowIdLanguageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identity", "identity", "identity"),
		resourceids.StaticSegment("b2cUserFlows", "b2cUserFlows", "b2cUserFlows"),
		resourceids.UserSpecifiedSegment("b2cIdentityUserFlowId", "b2cIdentityUserFlowId"),
		resourceids.StaticSegment("languages", "languages", "languages"),
		resourceids.UserSpecifiedSegment("userFlowLanguageConfigurationId", "userFlowLanguageConfigurationId"),
	}
}

// String returns a human-readable description of this Identity B 2c User Flow Id Language ID
func (id IdentityB2cUserFlowIdLanguageId) String() string {
	components := []string{
		fmt.Sprintf("B 2c Identity User Flow: %q", id.B2cIdentityUserFlowId),
		fmt.Sprintf("User Flow Language Configuration: %q", id.UserFlowLanguageConfigurationId),
	}
	return fmt.Sprintf("Identity B 2c User Flow Id Language (%s)", strings.Join(components, "\n"))
}
