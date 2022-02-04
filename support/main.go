package main

import (
	"errors"
)

type ResourceParser interface {
	Parse(values map[string]interface{}) (string, error)
}

var ResourceParsers = map[string]ResourceParser{
	"google_compute_instance":   NewGCPPacker("google_compute_instance"),
	"google_compute_disk":       NewGCPPacker("google_compute_disk"),
	"google_notebooks_instance": NewGCPPacker("google_notebooks_instance"),
}

type GPCParser struct {
	resource string
}

func NewGCPPacker(resource string) *GPCParser {
	return &GPCParser{resource: resource}
}

func (p *GPCParser) Parse(values map[string]interface{}) (string, error) {
	switch p.resource {
	case "google_compute_instance":
		// Parsing logic
		return "image-id", nil
	case "google_compute_disk":
		// Parsing logic
		return "image-id", nil
	case "google_notebooks_instance":
		// Parsing logic
		return "image-id", nil
	default:
		return "", errors.New("failed parse resource")
	}
}
