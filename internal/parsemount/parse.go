package parsemount

import "strings"

// Mount is the parsed form of a single --mount option.
type Mount struct {
	Tokens []string
	Type   string
	Source string
	From   string
}

// Parse parses one mount option, without a leading "--mount=".
// Callers pass the default type explicitly because the default depends on the
// caller's platform-specific compatibility behavior.
func Parse(mount, defaultType string) Mount {
	tokens := strings.Split(mount, ",")
	parsed := Mount{
		Tokens: tokens,
		Type:   defaultType,
	}
	for _, token := range tokens {
		key, value, hasValue := strings.Cut(token, "=")
		if !hasValue {
			continue
		}
		switch key {
		case "type":
			parsed.Type = value
		case "src", "source":
			parsed.Source = value
		case "from":
			parsed.From = value
		}
	}
	return parsed
}

// ParseFlag parses one mount flag, accepting either "--mount=..." or "...".
func ParseFlag(flag, defaultType string) Mount {
	return Parse(strings.TrimPrefix(flag, "--mount="), defaultType)
}
