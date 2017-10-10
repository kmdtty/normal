package normal 

import (
  "testing"
)

func TestNormal(t *testing.T) {
  result := merge_s("a", "b")
  expected := "abc"
  if result != expected {
    t.Errorf("The merge result: %q is not expecting word: %q", result, expected) 
  }
}
