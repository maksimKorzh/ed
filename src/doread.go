package main

import "os"
import "bufio"
import "fmt"

/* doread -- read "fil" after line n */
func doread(n int, fil string) stcode {
  file, err := os.Open(fil)
  if err != nil { return ERR }
  defer file.Close()
  curln = n
  scanner := bufio.NewScanner(file)
  count := 0
  for scanner.Scan() {
    line := scanner.Text()
    puttxt(line)
    count += len(line) + 1
  }
  fmt.Println(count)
  return OK
}

