package main

/* move -- move line1 through line2 after line3 */
func move(line3 *int) stcode {
  if line1 <= 0 || (*line3 >= line1 && *line3 < line2) {
    return ERR
  } else {
    blkmove(line1, line2, *line3);
    if (*line3 > line1) {
      curln = *line3
    } else {
      curln = *line3 + (line2 - line1 + 1)
    }
    return OK
  }
}
