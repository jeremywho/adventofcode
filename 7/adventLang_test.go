package adventLang

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestAllocProgram(t *testing.T) {
	lines := []string{""}
	p := NewAdventProgram(lines)

	if p.memory == nil {
		t.Errorf("memory was not allocated correctly, memory == nil")
	}

	if len(p.memory) != 0 {
		t.Errorf("memory was not allocated correctly, len != 0")
	}
}

func TestGetMemory(t *testing.T) {
	lines := []string{""}
	p := NewAdventProgram(lines)

	if p.GetMemory() == nil {
		t.Errorf("GetMemory() failed, memory == nil")
	}

	if len(p.GetMemory()) != 0 {
		t.Errorf("GetMemory failed, len != 0")
	}
}

func TestAssignment(t *testing.T) {
	lines := []string{"123 -> x", "x -> y"}
	p := NewAdventProgram(lines)
	p.Run()

	m := p.GetMemory()

	memVarX, ok := m["x"]
	if !ok {
		t.Errorf("Could not retrieve var x from memory")
	}

	if v := <-memVarX; v != 123 {
		t.Errorf("Incorrect value for var x, expected: %d, actual: %d", 123, v)
	}

	if len(p.GetMemory()) != 2 {
		t.Errorf("Assignment to y var failed, len != 2")
	}

	memVarY, ok := m["y"]
	if !ok {
		t.Errorf("Could not retrieve var y from memory")
	}

	if v := <-memVarY; v != 123 {
		t.Errorf("Incorrect value for var y, expected: %d, actual: %d", 123, v)
	}
}

func TestAndOperation(t *testing.T) {
	//note the order, we are trying to get the d value and we don't have the
	//x and y yet
	lines := []string{
		"123 -> x",
		"456 -> y",
		"x AND y -> d",
	}

	p := NewAdventProgram(lines)
	p.Run()

	m := p.GetMemory()

	d, ok := m["d"]
	if !ok {
		t.Errorf("Could not retrieve var d from memory")
	}

	if v := <-d; v != 72 {
		t.Errorf("Incorrect value for var d, expected: %d, actual: %d", 72, v)
	}
}

func TestOrOperation(t *testing.T) {
	lines := []string{
		"123 -> x",
		"456 -> y",
		"x OR y -> e",
	}
	p := NewAdventProgram(lines)
	p.Run()

	m := p.GetMemory()

	e, ok := m["e"]
	if !ok {
		t.Errorf("Could not retrieve var e from memory")
	}

	if v := <-e; v != 507 {
		t.Errorf("Incorrect value for var e, expected: %d, actual: %d", 507, v)
	}
}

func TestLeftShiftOperation(t *testing.T) {
	lines := []string{
		"123 -> x",
		"456 -> y",
		"x LSHIFT 2 -> f",
	}
	p := NewAdventProgram(lines)
	p.Run()

	m := p.GetMemory()

	f, ok := m["f"]
	if !ok {
		t.Errorf("Could not retrieve var f from memory")
	}

	if v := <-f; v != 492 {
		t.Errorf("Incorrect value for var f, expected: %d, actual: %d", 492, v)
	}
}

func TestRightShiftOperation(t *testing.T) {
	lines := []string{
		"123 -> x",
		"456 -> y",
		"y RSHIFT 2 -> g",
	}
	p := NewAdventProgram(lines)
	p.Run()

	m := p.GetMemory()

	g, ok := m["g"]
	if !ok {
		t.Errorf("Could not retrieve var g from memory")
	}

	if v := <-g; v != 114 {
		t.Errorf("Incorrect value for var g, expected: %d, actual: %d", 114, v)
	}
}

func TestNotOperation(t *testing.T) {
	lines := []string{
		"123 -> x",
		"456 -> y",
		"NOT x -> h",
		"NOT y -> i",
	}
	p := NewAdventProgram(lines)

	p.Run()

	m := p.GetMemory()

	h, ok := m["h"]
	if !ok {
		t.Errorf("Could not retrieve var h from memory")
	}

	if v := <-h; v != 65412 {
		t.Errorf("Incorrect value for var h, expected: %d, actual: %d", 65412, v)
	}

	i, ok := m["i"]
	if !ok {
		t.Errorf("Could not retrieve var i from memory")
	}

	if v := <-i; v != 65079 {
		t.Errorf("Incorrect value for var i, expected: %d, actual: %d", 65079, v)
	}
}

func TestInputFromFile(t *testing.T) {
	b, _ := ioutil.ReadFile("input.txt")
	input := string(b[:])
	lines := strings.Split(input, "\n")

	p := NewAdventProgram(lines)
	p.Run()

	mem := p.GetMemory()

	if aChan, ok := mem["a"]; ok {
		a := <-aChan
		t.Log("Got value for 'a': ", a)
	}
}
