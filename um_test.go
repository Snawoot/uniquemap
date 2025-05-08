package uniquemap

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestSimple(t *testing.T) {
	r := strings.NewReader(`
[
	{
		"x": 1,
		"y": 1,
		"name": "A"
	},
	{
		"x": 1,
		"y": 1,
		"name": "A"
	},
	{
		"x": 1,
		"y": 1,
		"name": "A"
	},
	{
		"x": 1,
		"y": 1,
		"name": "B"
	},
	{
		"x": 1,
		"y": 2,
		"name": "B"
	},
	{
		"x": 1,
		"y": 2,
		"name": "B"
	},
	{
		"x": 2,
		"y": 2,
		"name": "B"
	},
	{
		"x": 2,
		"y": 2,
		"name": "B"
	},
	{
		"x": 2,
		"y": 2,
		"name": "B"
	},
	{
		"comment": "dummy"
	}
]
`)
	var doc []any
	dec := json.NewDecoder(r)
	if err := dec.Decode(&doc); err != nil {
		t.Errorf("unable to decode JSON stream: %v", err)
		return
	}
	var handles []Handle[map[string]any, string, any]
	for _, elem := range doc {
		handles = append(handles, Make[map[string]any, string, any](elem.(map[string]any)))
	}
	if handles[0] != handles[1] {
		t.Errorf("handles[0] != handles[1]; %v != %v", handles[0].Value(), handles[1].Value())
	}
	if handles[0] != handles[2] {
		t.Errorf("handles[0] != handles[2]; %v != %v", handles[0].Value(), handles[2].Value())
	}
	if handles[0] == handles[3] {
		t.Errorf("handles[0] == handles[3]; %v != %v", handles[0].Value(), handles[3].Value())
	}
}
