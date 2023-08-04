package main

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
      if inline[0] == PERIOD && inline[1] == NEWLINE {
        done = true
      } else if puttxt(inline[:len(inline)-1]) == ERR {
        stat = ERR
      }
    }
  }
  return stat
}
