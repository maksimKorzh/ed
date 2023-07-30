package main

/* ctoi -- convert string at sri] to integer, increment i */
func ctoi(s string, i *int) int {
  var n, sign int
  skipbl(s, i)
  if s[*i] == '-' {
    sign = -1
  } else {
    sign = 1
  }
  if s[*i] == '+' || s[*i] == '-' { *i++ }
  for *i < len(s) && s[*i] >= '0' && s[*i] <= '9' {
    n = n * 10 + int(s[*i]) - '0'
    *i++
  }
  return sign * n
}
