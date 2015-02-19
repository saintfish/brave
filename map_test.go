package brave

import (
	"testing"
	"reflect"
)

func TestMapToken(t *testing.T) {
	type Node struct {
		I    int
		S    string
		Kids []Node
		p    string
	}
	type Tree Node
	type Forest []Tree

	var srcForest = Forest{
		Tree{I: 1, S: "a", Kids: []Node{
			Node{I: 2, S: "b"},
			Node{I: 3, S: "c"},
		}},
		Tree{I: 1, S: "b", p: "a"},
	}
	var expForest = Forest{
		Tree{I: 10, S: "aa", Kids: []Node{
			Node{I: 20, S: "bb"},
			Node{I: 30, S: "cc"},
		}},
		Tree{I: 10, S: "bb"},
	}
	var m = map_{
		1: 10, 2: 20, 3: 30,
		"a": "aa", "b": "bb", "c": "cc",
	}
	// Try multiple times to test src is not changed
	for i := 0; i < 10; i++ {
		dstForest, ok := mapData(srcForest, m).(Forest)
		if !ok {
			t.Errorf("mapToken doesn't return the same type")
			break
		}
		if !reflect.DeepEqual(dstForest, expForest) {
			t.Errorf("%v != %v", dstForest, expForest)
			break
		}
	}
}
