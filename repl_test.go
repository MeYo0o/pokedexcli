package main

import "testing"

func TestCleanInput(t *testing.T) {

	//* Put the cases
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		}, {
			input:    "MeYo 7allow!!",
			expected: []string{"MeYo", "7allow!!"},
		},
	}

	//* loop through them and check the conditions inside
	for _, c := range cases {
		actual := cleanInput(c.input)

		for i := range actual {
			word := actual[i]
			expected := c.expected[i]

			if word != expected {
				t.Errorf("Actual: %s != Expected: %s", word, expected)
			}
		}

	}
}
