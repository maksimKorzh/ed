# ed
The standard Linux text editor implementation in Go

# Implementation details
Current implementation is based on the source code from<br>
"The software tools in Pascal" by Brian Kernighan. On top<br>
of the original version it adds command "t" - copy block<br>
of text to a given line and also commands "y" - yank and<br>
"x" - put. The limitation is that global commands are not<br>
implemented however the behaviour of "s" (substitute) command<br>
is indeed global, so you can say "1,$s/pat/sub/" and have the<br>
pattern "pat" being replaced with "sub" string. Also "!" command is<br>
not available since I simply don't need it<br>
