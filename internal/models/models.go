// Package models holds the application model definitions.
package models

import "fmt"

// Course holds all relevant course data.
type Course struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Instructor  string `json:"instructor"`
	CentPrice   int64  `json:"centPrice"`
}

// IsValid returns an error if the course model is invalid.
func (c *Course) IsValid() error {
	if c.Name == "" {
		return fmt.Errorf("course, name is required")
	}
	if c.Description == "" {
		return fmt.Errorf("course, description is required")
	}
	if c.Instructor == "" {
		return fmt.Errorf("course, instructor is required")
	}
	if c.CentPrice < 1 {
		return fmt.Errorf("course, cent price must be grater then zero")
	}
	return nil
}

func (c *Course) Reset() {
	c.ID = 0
	c.Name = ""
	c.Description = ""
	c.Instructor = ""
	c.CentPrice = 0
}
