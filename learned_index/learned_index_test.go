package learned_index

import (
	"testing"
)

func TestLipp(t *testing.T) {
	ind := InitIndex()

	if ind.Insert(2, 3) != nil {
		t.Error("Wrong result for insert key")
	}

	if _, err := ind.Lookup(2); err != nil {
		t.Error("Wrong result for search existing key")
	}

	if ind.Insert(11, 3) != nil {
		t.Error("Wrong result for insert key")
	}

	if ind.Insert(12, 3) != nil {
		t.Error("Wrong result for insert key")
	}

	if ind.Insert(13, 4) != nil {
		t.Error("Wrong result for insert key")
	}

}
