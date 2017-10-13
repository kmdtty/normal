package normal 

import (
  "testing"
)

func Test_diff(t *testing.T) {
  d := diff(`{"a":"1", "b":"2"}`, `{"a":"3", "c":"4"}`)
  t.Logf("%v", d)
}

func _Test_merge(t *testing.T) {
  d := diff(`{"a":1, "b":2}`, `{"a":3, "c":4}`)
  result := merge(`{"a":1, "b":2}`, d)
  expected := `{"a":3, "c":4}`
  if result != expected {
    t.Errorf("The merge result: %q is not expecting word: %q", result, expected) 
  }
}
