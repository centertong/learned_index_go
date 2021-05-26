package learnedindex_test

import (
	"learned_index"
	"testing"
)

func TestLipp(t *testing.T) {
	ind := learned_index.InitIndex()
	
	if _, chk := ind.Lookup(2); chk {
		t.Error("Wrong result for search empty key")
	}
	
	if !ind.Insert(2, 3) {
		t.Error("Wrong result for insert key")
	}

	if 	_, chk := ind.Lookup(2); !chk {
		t.Error("Wrong result for search existing key")
	}
	
	if !ind.Insert(11, 3) {
		t.Error("Wrong result for insert key")
	}

	if !ind.Insert(12, 3) {
		t.Error("Wrong result for insert key")
	}

	if ind.Insert(12, 4) {
		t.Error("Wrong result for insert key")
	}

}
