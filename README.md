# Ed
The standard Linux text editor implementation in Go

# Implementation details
Current implementation is based on the source code from<br>
"The software tools in Pascal" by Brian Kernighan.<br>

# Commands
      Navigation:
    
    : 1              go to the first line
    : $              go to the last line
    : 12             go to line number 12
    : .              go to the current line (does nothing, used in expressions)
    : .+10           scroll down 10 lines
    : .-10           scroll up 10 lines
    : $-5            scroll to 5 lines before the last line
    
       Find/Replace:
    
    : /pat/          scroll to first "pat" occurrence ("pat" can be regexp )
    : /pat/;//       scroll to the second "pat" occurrence
    : //             scroll to the next "pat" occurrence
    : \\             scroll to the previous "pat" occurence
    : s/pat/sub/     substitute "pat" with "sub" on the current line, "pat" can be regexp
    : 2,5s/pat/sub/  substitute "pat" with "sub" within lines 2,5 inclusive
    
       I/O:
    
    : e file.txt     load "file.txt" to the buffer
    : r file.txt     insert content of the "file.txt" to the current line in buffer
    : f file.txt     set current file name to "file.txt"
    : w file.txt     save file as "file.txt"
    : w              save current file
    : q              exit from editor
    
# Usage
    $ ed                # opens editor with 'out.txt' source file name
    $ ed my_file.txt    # opens editor with 'my_file.txt' if it exists,
    # otherwise sets source filename to 'my_file.txt'
