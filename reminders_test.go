package activetask

import "testing"

func TestParseId(t *testing.T) {
	id := parseIdFromTask("#1234 Task task task")
	if id != 1234 {
		t.Fatalf("ID did not match: got %d", id)
	}
}
