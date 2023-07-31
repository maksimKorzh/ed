package main

import "fmt"

/* getline -- read user input from STDIN */
func getline() string {
  var line string
  var ch rune
  for ch != '\n' {
    _, err := fmt.Scanf("%c", &ch)
    if err != nil { panic(err) }
    line += string(ch)
  }
  return line
}
