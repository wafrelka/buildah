package parsemount

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		mount       string
		defaultType string
		expected    Mount
	}{
		{
			name:        "implicit mount type",
			mount:       "src=foo,target=bar,ro",
			defaultType: "nullfs",
			expected: Mount{
				Tokens: []string{"src=foo", "target=bar", "ro"},
				Type:   "nullfs",
				Source: "foo",
			},
		},
		{
			name:        "explicit bind type",
			mount:       "type=bind,source=foo,destination=bar,Z",
			defaultType: "nullfs",
			expected: Mount{
				Tokens: []string{"type=bind", "source=foo", "destination=bar", "Z"},
				Type:   "bind",
				Source: "foo",
			},
		},
		{
			name:        "explicit cache type",
			mount:       "type=cache,from=stage,target=/cache",
			defaultType: "bind",
			expected: Mount{
				Tokens: []string{"type=cache", "from=stage", "target=/cache"},
				Type:   "cache",
				From:   "stage",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, tt.expected, Parse(tt.mount, tt.defaultType))
		})
	}
}

func TestParseFlag(t *testing.T) {
	t.Parallel()

	expected := Mount{
		Tokens: []string{"type=bind", "source=foo", "destination=bar", "from=stage"},
		Type:   "bind",
		Source: "foo",
		From:   "stage",
	}

	actual := ParseFlag("--mount=type=bind,source=foo,destination=bar,from=stage", "nullfs")

	require.Equal(t, expected, actual)
}
