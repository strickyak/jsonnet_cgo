package jsonnet_test

import jsonnet "github.com/strickyak/jsonnet_cgo"

import (
	"errors"
	"fmt"
	"testing"
)

// Demo for the README.
func Test_Demo(t *testing.T) {
	vm := jsonnet.Make()
	vm.ExtVar("color", "purple")

	x, err := vm.EvaluateSnippet(`Test_Demo`, `"dark " + std.extVar("color")`)
	if err != nil {
		panic(err)
	}
	if x != "\"dark purple\"\n" {
		panic("fail: we got " + x)
	}
	vm.Destroy()
}

// importFunc returns a couple of hardwired responses.
func importFunc(base, rel string) (result string, path string, err error) {
	if rel == "alien.conf" {
		return `{ type: "alien", origin: "Ork", name: "Mork" }`, "alien.conf", nil
	}
	if rel == "human.conf" {
		return `{ type: "human", origin: "Earth", name: "Mendy" }`, "human.conf", nil
	}
	return "", "", errors.New(fmt.Sprintf("Cannot import %q", rel))
}

// check there is no err, and a == b.
func check(t *testing.T, err error, a, b string) {
	if err != nil {
		t.Errorf("got error: %q", err.Error())
	}
	if a != b {
		t.Errorf("got %q but wanted %q", a, b)
	}
}

func Test_Simple(t *testing.T) {

	// Each time there's a new version, this will force an update to this code.
	check(t, nil, jsonnet.Version(), `v0.9.3`)

	vm := jsonnet.Make()
	vm.TlaVar("color", "purple")
	vm.TlaVar("size", "XXL")
	vm.TlaCode("gooselevel", "1234 * 10 + 5")
	vm.ExtVar("color", "purple")
	vm.ExtVar("size", "XXL")
	vm.ExtCode("gooselevel", "1234 * 10 + 5")
	vm.ImportCallback(importFunc)

	x, err := vm.EvaluateSnippet(`test1`, `20 + 22`)
	check(t, err, x, `42`+"\n")
	x, err = vm.EvaluateSnippet(`test2`, `function(color, size, gooselevel) color`)
	check(t, err, x, `"purple"`+"\n")
	x, err = vm.EvaluateSnippet(`test2`, `std.extVar("color")`)
	check(t, err, x, `"purple"`+"\n")
	vm.StringOutput(true)
	x, err = vm.EvaluateSnippet(`test2`, `"whee"`)
	check(t, err, x, `whee`+"\n")
	vm.StringOutput(false)
	x, err = vm.EvaluateSnippet(`test3`, `
    local a = import "alien.conf";
    local b = import "human.conf";
    a.name + b.name
    `)
	check(t, err, x, `"MorkMendy"`+"\n")
	x, err = vm.EvaluateSnippet(`test4`, `
    local a = import "alien.conf";
    local b = a { type: "fictitious" };
    b.type + b.name
    `)
	check(t, err, x, `"fictitiousMork"`+"\n")
}

func Test_FileScript(t *testing.T) {
	vm := jsonnet.Make()
	x, err := vm.EvaluateFile("test2.j")
	check(t, err, x, `{
   "awk": "/usr/bin/awk",
   "shell": "/bin/csh"
}
`)
}

func Test_Misc(t *testing.T) {
	vm := jsonnet.Make()

	vm.MaxStack(10)
	vm.MaxTrace(10)
	vm.GcMinObjects(10)
	vm.GcGrowthTrigger(2.0)

	x, err := vm.EvaluateSnippet("Misc", `
    local a = import "test2.j";
    a.awk + a.shell`)
	check(t, err, x, `"/usr/bin/awk/bin/csh"`+"\n")
}
