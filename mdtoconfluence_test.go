package mdtoconfluence

import "testing"

func TestReplaceStringHeading(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"# Heading\n", "h1. Heading\n"},
		{"## Heading\n", "h2. Heading\n"},
		{"### Heading\n", "h3. Heading\n"},
		{"#### Heading\n", "h4. Heading\n"},
		{"##### Heading\n", "h5. Heading\n"},
		{"###### Heading\n", "h6. Heading\n"},
		{"h7. Heading\n", "h7. Heading\n"},
		{"Heading\n", "Heading\n"},
		{"# Heading\n## Heading ### Heading\n", "h1. Heading\nh2. Heading ### Heading\n"},
	}
	for _, c := range cases {
		got := ReplaceStringHeading(c.in)
		if got != c.want {
			t.Errorf("ReplaceHeading(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
