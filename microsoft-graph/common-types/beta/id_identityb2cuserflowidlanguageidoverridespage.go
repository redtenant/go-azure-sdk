package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityB2cUserFlowIdLanguageIdOverridesPageId{}

// IdentityB2cUserFlowIdLanguageIdOverridesPageId is a struct representing the Resource ID for a Identity B 2c User Flow Id Language Id Overrides Page
type IdentityB2cUserFlowIdLanguageIdOverridesPageId struct {
	B2cIdentityUserFlowId           string
	UserFlowLanguageConfigurationId string
	UserFlowLanguagePageId          string
}

// NewIdentityB2cUserFlowIdLanguageIdOverridesPageID returns a new IdentityB2cUserFlowIdLanguageIdOverridesPageId struct
func NewIdentityB2cUserFlowIdLanguageIdOverridesPageID(b2cIdentityUserFlowId string, userFlowLanguageConfigurationId string, userFlowLanguagePageId string) IdentityB2cUserFlowIdLanguageIdOverridesPageId {
	return IdentityB2cUserFlowIdLanguageIdOverridesPageId{
		B2cIdentityUserFlowId:           b2cIdentityUserFlowId,
		UserFlowLanguageConfigurationId: userFlowLanguageConfigurationId,
		UserFlowLanguagePageId:          userFlowLanguagePageId,
	}
}

// ParseIdentityB2cUserFlowIdLanguageIdOverridesPageID parses 'input' into a IdentityB2cUserFlowIdLanguageIdOverridesPageId
func ParseIdentityB2cUserFlowIdLanguageIdOverridesPageID(input string) (*IdentityB2cUserFlowIdLanguageIdOverridesPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityB2cUserFlowIdLanguageIdOverridesPageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityB2cUserFlowIdLanguageIdOverridesPageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityB2cUserFlowIdLanguageIdOverridesPageIDInsensitively parses 'input' case-insensitively into a IdentityB2cUserFlowIdLanguageIdOverridesPageId
// note: this method should only be used for API response data and not user input
func ParseIdentityB2cUserFlowIdLanguageIdOverridesPageIDInsensitively(input string) (*IdentityB2cUserFlowIdLanguageIdOverridesPageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityB2cUserFlowIdLanguageIdOverridesPageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityB2cUserFlowIdLanguageIdOverridesPageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityB2cUserFlowIdLanguageIdOverridesPageId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateIdentityB2cUserFlowIdLanguageIdOverridesPageID checks that 'input' can be parsed as a Identity B 2c User Flow Id Language Id Overrides Page ID
func ValidateIdentityB2cUserFlowIdLanguageIdOverridesPageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityB2cUserFlowIdLanguageIdOverridesPageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity B 2c User Flow Id Language Id Overrides Page ID
func (id IdentityB2cUserFlowIdLanguageIdOverridesPageId) ID() string {
	fmtString := "/identity/b2cUserFlows/%s/languages/%s/overridesPages/%s"
	return fmt.Sprintf(fmtString, id.B2cIdentityUserFlowId, id.UserFlowLanguageConfigurationId, id.UserFlowLanguagePageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity B 2c User Flow Id Language Id Overrides Page ID
func (id IdentityB2cUserFlowIdLanguageIdOverridesPageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identity", "identity", "identity"),
		resourceids.StaticSegment("b2cUserFlows", "b2cUserFlows", "b2cUserFlows"),
		resourceids.UserSpecifiedSegment("b2cIdentityUserFlowId", "b2cIdentityUserFlowId"),
		resourceids.StaticSegment("languages", "languages", "languages"),
		resourceids.UserSpecifiedSegment("userFlowLanguageConfigurationId", "userFlowLanguageConfigurationId"),
		resourceids.StaticSegment("overridesPages", "overridesPages", "overridesPages"),
		resourceids.UserSpecifiedSegment("userFlowLanguagePageId", "userFlowLanguagePageId"),
	}
}

// String returns a human-readable description of this Identity B 2c User Flow Id Language Id Overrides Page ID
func (id IdentityB2cUserFlowIdLanguageIdOverridesPageId) String() string {
	components := []string{
		fmt.Sprintf("B 2c Identity User Flow: %q", id.B2cIdentityUserFlowId),
		fmt.Sprintf("User Flow Language Configuration: %q", id.UserFlowLanguageConfigurationId),
		fmt.Sprintf("User Flow Language Page: %q", id.UserFlowLanguagePageId),
	}
	return fmt.Sprintf("Identity B 2c User Flow Id Language Id Overrides Page (%s)", strings.Join(components, "\n"))
}
