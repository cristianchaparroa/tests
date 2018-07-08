package main

import(
  "testing"
)

func TestFirstNotRepeatingCharacter(t *testing.T) {
  var test = []struct {
    expetected string
    s  string
  } {

    {"c","abacabad"},
    {"_","abacabaabacaba"},
    {"z","z"},
    {"c", "bcb"},
    {"_","bcccccccb"},
    {"d","abcdefghijklmnopqrstuvwxyziflskecznslkjfabe"},
    {"_","zzz"},
    {"y","bcccccccccccccyb"},
    {"d","xdnxxlvupzuwgigeqjggosgljuhliybkjpibyatofcjbfxwtalc"},
    {"g", "ngrhhqbhnsipkcoqjyviikvxbxyphsnjpdxkhtadltsuxbfbrkof"},
  }

  for _, ts := range test {
    r := firstNotRepeatingCharacter(ts.s)
    if ts.expetected != r {
      t.Errorf("Expected:%v, but get:%v", ts.expetected, r)
    }
  }
}
