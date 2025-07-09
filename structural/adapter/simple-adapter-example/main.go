package main

import (
	"fmt"
	"strings"
)

// Abstract Product
type VideoMetadata interface {
	GetTitle() string
	GetTags() []string
	FetchData()
}

// Simple Video Metadata
type SimpleVideoMetadata struct {
	title string
	tags  []string
}

func (v *SimpleVideoMetadata) FetchData() {
	v.title = "Example Video Title"
	v.tags = []string{"example", "education"}
}
func (v *SimpleVideoMetadata) GetTitle() string {
	return v.title
}

func (v *SimpleVideoMetadata) GetTags() []string {
	return v.tags
}

// Legacy Video Metadata
type LegacyVideoMetadata struct {
	meta_title string
	meta_tags  string
}

func (v *LegacyVideoMetadata) GetTitle() string {
	return v.meta_title
}
func (v *LegacyVideoMetadata) GetTags() string {
	return v.meta_tags
}
func (v *LegacyVideoMetadata) FetchData() {
	v.meta_title = "Old Video Title"
	v.meta_tags = "old,legacy"
}

type LegacyVideoMetadataAdapter struct {
	legacyVideoMetadata *LegacyVideoMetadata
}

func (a *LegacyVideoMetadataAdapter) GetTitle() string {
	return a.legacyVideoMetadata.meta_title
}
func (a *LegacyVideoMetadataAdapter) GetTags() []string {
	return strings.Split(a.legacyVideoMetadata.meta_tags, ",")
}
func (a *LegacyVideoMetadataAdapter) FetchData() {
	a.legacyVideoMetadata.FetchData()
}

// Functions
func NewSimpleVideoMetadata() *SimpleVideoMetadata {
	return &SimpleVideoMetadata{}
}
func NewLegacyVideoMetadata() *LegacyVideoMetadata {
	return &LegacyVideoMetadata{}
}
func NewLegacyVideoMetadataAdapter(legacyVideoMetadata *LegacyVideoMetadata) *LegacyVideoMetadataAdapter {
	return &LegacyVideoMetadataAdapter{legacyVideoMetadata: legacyVideoMetadata}
}

func PrintMetadata(video VideoMetadata) {
	fmt.Printf("Video Metadata\nTitle: %s\nTags: %s\n", video.GetTitle(), strings.Join(video.GetTags(), ", "))
}

func main() {
	simpleVideoMetadata := NewSimpleVideoMetadata()
	simpleVideoMetadata.FetchData()

	legacyVideoMetadata := NewLegacyVideoMetadata()

	legacyVideoAdapter := NewLegacyVideoMetadataAdapter(legacyVideoMetadata)
	legacyVideoAdapter.FetchData()

	PrintMetadata(simpleVideoMetadata)
	PrintMetadata(legacyVideoAdapter)
}
