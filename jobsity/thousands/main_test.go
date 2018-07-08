package main

import (
  "testing"
)

func TestThousands(t *testing.T) {
  var test = []struct {
    number int
    expected string
  } {
    {1234, "1,234"},
    {123456789,"123,456,789"},
    { 12, "12" },
  }

  for _,ts := range test {
    r := Thousands(ts.number)
    if (ts.expected != r) {
      t.Errorf("Expected:%v, but get:%v\n",ts.expected,r)
    }
  }
}
