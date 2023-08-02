package main

import "strings"

/* getword -- get word from s[i] into out */
func getword(s string, i *int, out *string) int {
  skipbl(s, i)
  name := strings.Split(s[*i:], " ")[0]
  name = name[:len(name)-1]
  *out = name
  *i += len(*out)-1
  if len(s) == 0 {
    return 0
  } else {
    return *i
  }
}
