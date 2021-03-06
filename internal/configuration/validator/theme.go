package validator

import (
	"github.com/authelia/authelia/internal/configuration/schema"
)

// ValidateTOTP validates and update TOTP configuration.
func ValidateTheme(configuration *schema.ThemeConfiguration, validator *schema.StructValidator) {
	if configuration.Name == "" {
		configuration.Name = schema.DefaultThemeConfiguration.Name
	}

	if configuration.PrimaryColor == "" {
		configuration.PrimaryColor = schema.DefaultThemeConfiguration.PrimaryColor
	}

	if configuration.SecondaryColor == "" {
		configuration.SecondaryColor = schema.DefaultThemeConfiguration.SecondaryColor
	}
}
