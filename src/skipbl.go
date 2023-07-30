package main

/* skipbl -- skip blanks and tabs at s[i]... */
func skipbl(s string, i *int) {
  if *i >= len(s) {return }
  for s[*i] == BLANK || s[*i] == TAB {
    *i ++
    if *i >= len(s) {return }
  }
}
