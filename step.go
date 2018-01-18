package cwl

// Steps represents "steps" field in CWL.
type Steps []Step

// New constructs "Steps" from interface.
func (_ Steps) New(i interface{}) Steps {
	dest := Steps{}
	switch x := i.(type) {
	case []interface{}:
		for _, v := range x {
			s := Step{}.New(v)
			dest = append(dest, s)
		}
	case map[string]interface{}:
		for key, v := range x {
			s := Step{}.New(v)
			s.ID = key
			dest = append(dest, s)
		}
	}
	return dest
}

// Step represents WorkflowStep.
// @see http://www.commonwl.org/v1.0/Workflow.html#WorkflowStep
type Step struct {
	ID            string
	In            StepInputs
	Out           []StepOutput
	Run           Run
	Requirements  []Requirement
	Scatter       []string
	ScatterMethod string
}

// Run `run` accept string | CommandLineTool | ExpressionTool | Workflow
type Run struct {
	Value    string
	Workflow *Root
}

// New constructs "Step" from interface.
func (_ Step) New(i interface{}) Step {
	dest := Step{}
	switch x := i.(type) {
	case map[string]interface{}:
		for key, v := range x {
			switch key {
			case "id":
				dest.ID = v.(string)
			case "run":
				switch x2 := v.(type) {
				case string:
					dest.Run.Value = x2
				case map[string]interface{}:
					dest.Run.Workflow = dest.Run.Workflow.AsStep(v)
				}
			case "in":
				dest.In = StepInput{}.NewList(v)
			case "out":
				dest.Out = StepOutput{}.NewList(v)
			case "requirements":
				dest.Requirements = Requirements{}.New(v)
			case "scatter":
				dest.Scatter = StringArrayable(v)
			case "scatterMethod":
				dest.ScatterMethod = v.(string)
			}
		}
	}
	return dest
}

// Len for sorting
func (steps Steps) Len() int {
	return len(steps)
}

// Less for sorting
func (steps Steps) Less(i, j int) bool {
	return steps[i].ID < steps[j].ID
}

// Swap for sorting
func (steps Steps) Swap(i, j int) {
	steps[i], steps[j] = steps[j], steps[i]
}
