package main
//import "fmt"
/* lnappend -- append lines after "line" */
func lnappend(line int, glob bool) stcode {
  var inline string
  var stat stcode
  var done bool
  if (glob) {
    stat = ERR
  } else {
    curln = line
    stat = OK
    done = false
    for done == false && stat == OK {
      inline = getline()
      //if len(inline) == 1 {
      //  stat = ENDDATA
      //} else 
      if inline[0] == PERIOD && inline[1] == NEWLINE {
        done = true
      } else if puttxt(inline) == ERR {
        stat = ERR
      }
    }
  }
  return stat
}
