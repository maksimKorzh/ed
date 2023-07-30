package main

/* skipbl -- skip blanks and tabs at s[i]... */
func skipbl(s string, i *int) {
  for s[*i] == BLANK || s[*i] == TAB { *i++ }
}
