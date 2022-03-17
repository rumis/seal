package utils

import "testing"

func TestStruct2Map(t *testing.T) {

	in := struct {
		Name string `seal:"name"`
		Stat int    `seal:"stat"`
	}{
		Name: "murong",
		Stat: 1,
	}

	out, err := Struct2Map(in)
	if err != nil {
		t.Fatal(err)
	}

	iname, ok := out["name"]
	if !ok {
		t.Fatal("name error")
	}
	if name, ok := iname.(string); !ok || name != "murong" {
		t.Fatal("name type error")
	}

	istat, ok := out["stat"]
	if !ok {
		t.Fatal("stat error")
	}
	if stat, ok := istat.(int); !ok || stat != 1 {
		t.Fatal("name type error")
	}
}

func TestStruct2MapSlice(t *testing.T) {

	in := []struct {
		Name string `seal:"name"`
		Stat int    `seal:"stat"`
	}{{
		Name: "murong",
		Stat: 1,
	}}

	out, err := Struct2MapSlice(in)
	if err != nil {
		t.Fatal(err)
	}

	iname, ok := out[0]["name"]
	if !ok {
		t.Fatal("name error")
	}
	if name, ok := iname.(string); !ok || name != "murong" {
		t.Fatal("name type error")
	}

	istat, ok := out[0]["stat"]
	if !ok {
		t.Fatal("stat error")
	}
	if stat, ok := istat.(int); !ok || stat != 1 {
		t.Fatal("name type error")
	}
}
