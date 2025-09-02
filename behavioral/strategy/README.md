# Strategy

> Strategy is a behavioral design pattern that lets you define a family of algorithms, put each of them into a separate class, and make their objects interchangeable. (Refactoring Guru)

## Example

### Basic example

```go
package main

import "fmt"

type RouteContext struct {
 strategy RouteStrategy
}

func (c *RouteContext) SetStrategy(s RouteStrategy) {
 c.strategy = s
}

func (c *RouteContext) GetDurationInMinutes() int {
 return c.strategy.GetDurationInMinutes()
}

// RouteStrategy
type RouteStrategy interface {
 GetDurationInMinutes() int
}

type CarRouteStrategy struct{}

func (s *CarRouteStrategy) GetDurationInMinutes() int {
 return 10
}

type WalkRouteStrategy struct{}

func (s *WalkRouteStrategy) GetDurationInMinutes() int {
 return 25
}

func main() {
 carStrategy := &CarRouteStrategy{}
 context := &RouteContext{}
 context.SetStrategy(carStrategy)
 fmt.Println("Card route time (in minutes): ", context.strategy.GetDurationInMinutes())

 walkStrategy := &WalkRouteStrategy{}
 context.SetStrategy(walkStrategy)
 fmt.Println("Walking route time (in minutes): ", context.strategy.GetDurationInMinutes())
}
```

### Output format example

```go
package main

import (
 "fmt"
 "strings"
)

type OutputFormat int

const (
 Markdown OutputFormat = iota
 Html
)

type ListStrategy interface {
 Start(builder *strings.Builder)
 End(builder *strings.Builder)
 AddListItem(builder *strings.Builder, item string)
}

type MarkdownListStrategy struct{}

func (s *MarkdownListStrategy) Start(builder *strings.Builder) {
}
func (s *MarkdownListStrategy) End(builder *strings.Builder) {
}
func (s *MarkdownListStrategy) AddListItem(builder *strings.Builder, item string) {
 builder.WriteString(" * " + item + "\n")
}

type HtmlListStrategy struct{}

func (s *HtmlListStrategy) Start(builder *strings.Builder) {
 builder.WriteString("<ul>\n")
}
func (s *HtmlListStrategy) End(builder *strings.Builder) {
 builder.WriteString("</ul>\n")
}
func (s *HtmlListStrategy) AddListItem(builder *strings.Builder, item string) {
 builder.WriteString("  <li>" + item + "</li>\n")
}

type TextProcessor struct {
 builder      strings.Builder
 ListStrategy ListStrategy
}

func NewTextProcessor(listStrategy ListStrategy) *TextProcessor {
 return &TextProcessor{strings.Builder{}, listStrategy}
}

func (t *TextProcessor) SetOutputFormat(fmt OutputFormat) {
 switch fmt {
 case Markdown:
  t.ListStrategy = &MarkdownListStrategy{}
 case Html:
  t.ListStrategy = &HtmlListStrategy{}
 }
}

func (t *TextProcessor) AppendList(items []string) {
 s := t.ListStrategy
 s.Start(&t.builder)
 for _, item := range items {
  s.AddListItem(&t.builder, item)
 }
 s.End(&t.builder)
}

func (t *TextProcessor) Reset() {
 t.builder.Reset()
}

func (t *TextProcessor) String() string {
 return t.builder.String()
}

func main() {
 tp := NewTextProcessor(&MarkdownListStrategy{})
 tp.AppendList([]string{"foo", "bar", "baz"})
 fmt.Println(tp)

 tp.Reset()
 tp.SetOutputFormat(Html)

 tp.AppendList([]string{"foo", "bar", "baz"})
 fmt.Println(tp)
}
```
