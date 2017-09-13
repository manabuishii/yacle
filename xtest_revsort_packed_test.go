package cwl

import (
	"testing"

	. "github.com/otiai10/mint"
)

func TestDecode_revsort_packed(t *testing.T) {
        f := cwl("revsort-packed.cwl")
        root := NewCWL()
        err := root.Decode(f)
        Expect(t, err).ToBe(nil)
        Expect(t, root.Version).ToBe("v1.0")
        Expect(t, len(root.Graphs)).ToBe(3)
        // Graph 0
        Expect(t, root.Graphs[0].Class).ToBe("Workflow")
        Expect(t, root.Graphs[0].ID).ToBe("#main")
        Expect(t, root.Graphs[0].Doc).ToBe("Reverse the lines in a document, then sort those lines.")
        Expect(t, root.Graphs[0].Hints[0].Class).ToBe("DockerRequirement")
        Expect(t, root.Graphs[0].Hints[0].DockerPull).ToBe("debian:8")
        Expect(t, len(root.Graphs[0].Inputs)).ToBe(2)
        Expect(t, root.Graphs[0].Inputs[0].Types[0].Type).ToBe("File")
        Expect(t, root.Graphs[0].Inputs[0].ID).ToBe("#main/input")
        Expect(t, root.Graphs[0].Inputs[0].Doc).ToBe("The input file to be processed.")
        Expect(t, root.Graphs[0].Inputs[1].Types[0].Type).ToBe("boolean")
        Expect(t, root.Graphs[0].Inputs[1].ID).ToBe("#main/reverse_sort")
        Expect(t, root.Graphs[0].Inputs[1].Doc).ToBe("If true, reverse (decending) sort")
        Expect(t, len(root.Graphs[0].Outputs)).ToBe(1)
        Expect(t, root.Graphs[0].Outputs[0].ID).ToBe("#main/output")
        Expect(t, root.Graphs[0].Outputs[0].Types[0].Type).ToBe("File")
        Expect(t, root.Graphs[0].Outputs[0].Doc[0]).ToBe("The output with the lines reversed and sorted.")
        Expect(t, root.Graphs[0].Outputs[0].Source[0]).ToBe("#main/sorted/output")
        Expect(t, root.Graphs[0].Steps[0].ID).ToBe("#main/rev")
        Expect(t, root.Graphs[0].Steps[0].Run.Value).ToBe("#revtool.cwl")
        Expect(t, len(root.Graphs[0].Steps[0].In)).ToBe(1)
        Expect(t, root.Graphs[0].Steps[0].In[0].Source[0]).ToBe("#main/input")
        Expect(t, root.Graphs[0].Steps[0].In[0].ID).ToBe("#main/rev/input")
        Expect(t, root.Graphs[0].Steps[0].Out[0].ID).ToBe("#main/rev/output")
        Expect(t, root.Graphs[0].Steps[1].Out[0].ID).ToBe("#main/sorted/output")
        Expect(t, root.Graphs[0].Steps[1].ID).ToBe("#main/sorted")
        Expect(t, root.Graphs[0].Steps[1].Run.Value).ToBe("#sorttool.cwl")
        Expect(t, len(root.Graphs[0].Steps[1].In)).ToBe(2)
        Expect(t, root.Graphs[0].Steps[1].In[0].Source[0]).ToBe("#main/rev/output")
        Expect(t, root.Graphs[0].Steps[1].In[0].ID).ToBe("#main/sorted/input")
        Expect(t, root.Graphs[0].Steps[1].In[1].Source[0]).ToBe("#main/reverse_sort")
        Expect(t, root.Graphs[0].Steps[1].In[1].ID).ToBe("#main/sorted/reverse")
        Expect(t, root.Graphs[0].Steps[1].Out[0].ID).ToBe("#main/sorted/output")
        // Graph 1
        Expect(t, root.Graphs[1].Class).ToBe("CommandLineTool")
        Expect(t, root.Graphs[1].ID).ToBe("#revtool.cwl")
        Expect(t, root.Graphs[1].Doc).ToBe("Reverse each line using the `rev` command")
        Expect(t, root.Graphs[1].BaseCommands[0]).ToBe("rev")
        Expect(t, root.Graphs[1].Stdout).ToBe("output.txt")
        Expect(t, len(root.Graphs[1].Inputs)).ToBe(1)
        Expect(t, root.Graphs[1].Inputs[0].Types[0].Type).ToBe("File")
        Expect(t, root.Graphs[1].Inputs[0].ID).ToBe("#revtool.cwl/input")
        Expect(t, len(root.Graphs[1].Outputs)).ToBe(1)
        Expect(t, root.Graphs[1].Outputs[0].ID).ToBe("#revtool.cwl/output")
        Expect(t, root.Graphs[1].Outputs[0].Types[0].Type).ToBe("File")
        Expect(t, root.Graphs[1].Outputs[0].Binding.Glob[0]).ToBe("output.txt")
        // Graph 2
        Expect(t, root.Graphs[2].Class).ToBe("CommandLineTool")
        Expect(t, root.Graphs[2].ID).ToBe("#sorttool.cwl")
        Expect(t, root.Graphs[2].Doc).ToBe("Sort lines using the `sort` command")
        Expect(t, root.Graphs[2].BaseCommands[0]).ToBe("sort")
        Expect(t, root.Graphs[2].Stdout).ToBe("output.txt")
        Expect(t, len(root.Graphs[2].Inputs)).ToBe(2)
        Expect(t, root.Graphs[2].Inputs[0].Types[0].Type).ToBe("boolean")
        Expect(t, root.Graphs[2].Inputs[0].ID).ToBe("#sorttool.cwl/reverse")
        Expect(t, root.Graphs[2].Inputs[0].Binding.Position).ToBe(1)
        Expect(t, root.Graphs[2].Inputs[0].Binding.Prefix).ToBe("--reverse")
        Expect(t, root.Graphs[2].Inputs[1].Types[0].Type).ToBe("File")
        Expect(t, root.Graphs[2].Inputs[1].ID).ToBe("#sorttool.cwl/input")
        Expect(t, root.Graphs[2].Inputs[1].Binding.Position).ToBe(2)
        Expect(t, len(root.Graphs[2].Outputs)).ToBe(1)
        Expect(t, root.Graphs[2].Outputs[0].ID).ToBe("#sorttool.cwl/output")
        Expect(t, root.Graphs[2].Outputs[0].Types[0].Type).ToBe("File")
        Expect(t, root.Graphs[2].Outputs[0].Binding.Glob[0]).ToBe("output.txt")
}
