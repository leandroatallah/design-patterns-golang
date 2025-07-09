# Adapter

> **Adapter** is a structural design pattern that allows objects with incompatible interfaces to collaborate. (Refactoring Guru)

## Structure

1. The **Client** is a class that contains the existing business logic of the program.
2. The **Client Interface** describes a protocol that other classes must follow to be able to collaborate with the client code.
3. The **Service** is some useful class (usually 3rd-party or legacy). The client can’t use this class directly because it has an incompatible interface.
4. The **Adapter** is a class that’s able to work with both the client and the service: it implements the client interface, while wrapping the service object. The adapter receives calls from the client via the client interface and translates them into calls to the wrapped service object in a format it can understand.
5. The client code doesn’t get coupled to the concrete adapter class as long as it works with the adapter via the client interface. Thanks to this, you can introduce new types of adapters into the program without breaking the existing client code. This can be useful when the interface of the service class gets changed or replaced: you can just create a new adapter class without changing the client code.

## Example

```go
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
```
