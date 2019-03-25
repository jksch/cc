package models

import (
	"fmt"
	"testing"
)

func TestIsCourseValid(t *testing.T) {
	var tests = []struct {
		name  string
		given Course
		exp   string
	}{
		{
			name: "Invalid name",
			given: Course{
				Name: "",
			},
			exp: "course, name is required",
		},
		{
			name: "Invalid description",
			given: Course{
				Name:        "Go for Gophers",
				Description: "",
			},
			exp: "course, description is required",
		},
		{
			name: "Invalid instructor",
			given: Course{
				Name:        "Go for Gophers",
				Description: "Some interesting go.",
				Instructor:  "",
			},
			exp: "course, instructor is required",
		},
		{
			name: "Invalid price, zero",
			given: Course{
				Name:        "Go for Gophers",
				Description: "Some interesting go.",
				Instructor:  "John Doe",
				CentPrice:   0,
			},
			exp: "course, cent price must be grater then zero",
		},
		{
			name: "Invalid price, negative",
			given: Course{
				Name:        "Go for Gophers",
				Description: "Some interesting go.",
				Instructor:  "John Doe",
				CentPrice:   -1,
			},
			exp: "course, cent price must be grater then zero",
		},
		{
			name: "Valid",
			given: Course{
				Name:        "Go for Gophers",
				Description: "Some interesting go.",
				Instructor:  "John Doe",
				CentPrice:   1000,
			},
			exp: "",
		},
	}

	for ti, tt := range tests {
		ti, tt := ti, tt
		t.Run(fmt.Sprintf("%d. %s", ti, tt.name), func(t *testing.T) {
			t.Parallel()
			got := errStr(tt.given.IsValid())
			if tt.exp != got {
				t.Errorf("%d. %s, exp err: '%s' got: '%s'", ti, tt.name, tt.exp, got)
			}
		})
	}
}

func errStr(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}
