package main

import "fmt"

/*
* interface Builder
* class Director
* ConcreteBuilderProduct1
* ConcreteBuilderProduct2
* class product1 (FeatureTask)
* class product2 (DefectTask)
 */

// Builder interface
type TaskBuilder interface{}

// Director class
type TaskDirector struct{}

func (d *TaskDirector) makeFeatureTask(builder *FeatureTaskBuilder, title string, priority int8) {
	builder.Reset()
	builder.SetTitle(title)
	builder.SetPriority(priority)
}
func (d *TaskDirector) makeDefectTask(builder *DefectTaskBuilder, title string, priority int8) {
	builder.Reset()
	builder.SetTitle(title)
	builder.SetPriority(priority)
}

// Feature Task
type FeatureTask struct {
	title    string
	priority int8
}

func (t *FeatureTask) SetTitle(title string) {
	t.title = title
}
func (t *FeatureTask) SetPriority(priority int8) {
	t.priority = priority
}

// Defect Task
type DefectTask struct {
	title    string
	priority int8
}

func (t *DefectTask) SetTitle(title string) {
	t.title = title
}
func (t *DefectTask) SetPriority(priority int8) {
	t.priority = priority
}

// Concrete builder for Feature Task
type FeatureTaskBuilder struct {
	result FeatureTask
}

func (b *FeatureTaskBuilder) Reset() {
	b.result.SetTitle("")
	b.result.SetPriority(0)
}
func (b *FeatureTaskBuilder) SetTitle(title string) {
	b.result.SetTitle(title)
}
func (b *FeatureTaskBuilder) SetPriority(priority int8) {
	b.result.SetPriority(priority)
}
func (b *FeatureTaskBuilder) GetResult() FeatureTask {
	return b.result
}

// Concrete builder for Defect Task
type DefectTaskBuilder struct {
	result DefectTask
}

func (b *DefectTaskBuilder) Reset() {
	b.result.SetTitle("")
	b.result.SetPriority(0)
}
func (b *DefectTaskBuilder) SetTitle(title string) {
	b.result.SetTitle(title)
}
func (b *DefectTaskBuilder) SetPriority(priority int8) {
	b.result.SetPriority(priority)
}
func (b *DefectTaskBuilder) GetResult() DefectTask {
	return b.result
}

func main() {
	director := TaskDirector{}

	featureBuilder := FeatureTaskBuilder{}
	director.makeFeatureTask(&featureBuilder, "Create a new button variation", 3)
	feature := featureBuilder.GetResult()
	fmt.Println(feature)

	defectBuilder := DefectTaskBuilder{}
	director.makeDefectTask(&defectBuilder, "Fix request to product list", 5)
	defect := defectBuilder.GetResult()
	fmt.Println(defect)

}
