package mdtoconfluence

import "testing"

func TestReplaceHeading(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"h1. Heading", "# Heading"},
		{"h2. Heading", "## Heading"},
		{"h3. Heading", "### Heading"},
		{"h4. Heading", "#### Heading"},
		{"h5. Heading", "##### Heading"},
		{"h6. Heading", "###### Heading"},
		{"h7. Heading", "h7. Heading"},
		{"Heading", "Heading"},
	}
	for _, c := range cases {
		got := ReplaceHeading(c.in)
		if got != c.want {
			t.Errorf("ReplaceHeading(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
