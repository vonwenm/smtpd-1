//
//
package main

import (
	"testing"
)

//
// Test address and host matches, since so much depends on them.
var aMatches = []struct {
	a, pat string
}{
	{"abc@def", "abc@def"},
	{"", "<>"},
	{"abc@def", "abc@"},
	{"abc@def", "@def"},
	{"abc", "abc@"},
	{"abc@def.ghi", "@.ghi"},
	{"abc@def.ghi", "@.def.ghi"},
	{"anything@anything", "@"},
}
var nMatches = []struct {
	a, pat string
}{
	{"abc@def", "not"},
	{"abc@def", "@ghi"},
	{"abc@def", "def@"},
	{"anything", "<>"},
	{"noat", "@"},
	{"abc@zamdef.ghi", "@.def.ghi"},
}

func TestAddrMatches(t *testing.T) {
	for _, inp := range aMatches {
		if !matchAddress(inp.a, inp.pat) {
			t.Errorf("did not match '%s' with pattern '%s'", inp.a, inp.pat)
		}
	}
	for _, inp := range nMatches {
		if matchAddress(inp.a, inp.pat) {
			t.Errorf("did match '%s' with pattern '%s'", inp.a, inp.pat)
		}
	}
}

var ahMatches = []struct {
	h, pat string
}{
	{"abc", "abc"},
	{"abc", ".abc"},
	{"abc.def", ".def"},
	{"abc.def.ghi", ".ghi"},
}
var nhMatches = []struct {
	h, pat string
}{
	{"abc", "not"},
	{"abc", ".not"},
	{"prefabc", ".abc"},
}

func TestHostMatches(t *testing.T) {
	for _, inp := range ahMatches {
		if !matchHost(inp.h, inp.pat) {
			t.Errorf("did not match '%s' with pattern '%s'", inp.h, inp.pat)
		}
	}
	for _, inp := range nhMatches {
		if matchHost(inp.h, inp.pat) {
			t.Errorf("did match '%s' with pattern '%s'", inp.h, inp.pat)
		}
	}
}
