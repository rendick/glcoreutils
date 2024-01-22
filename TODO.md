| GNU      | Description     | GLCU                                                                                         |
| -------- | --------------- | -------------------------------------------------------------------------------------------- |
| chgrp    | File utilities  | Changes file group ownership                                                                 |
| chown    | File utilities  | Changes file ownership                                                                       |
| chmod    | File utilities  | Changes the permissions of a file or directory                                               |
| cp       | File utilities  | Copy file or directory                                                                       |
| dd       | File utilities  | Copies and converts a file                                                                   |
| df       | File utilities  | Shows disk free space on file systems                                                        |
| install  | File utilities  | Copies files and set attributes                                                              |
| ln       | File utilities  | Creates a link to a file                                                                     |
| ls       | File utilities  | [ls](https://github.com/rendick/glcoreutils/blob/main/src/ls.go)                             |
| mkdir    | File utilities  | [mkdir](https://github.com/rendick/glcoreutils/blob/main/src/mkdir.go)                       |
| mkfifo   | File utilities  | Makes named pipesNamed_pipe "Named pipe" FIFOs                                               |
| mknod    | File utilities  | Makes block or character special filesDevice_node "Device node"                              |
| mktemp   | File utilities  | Creates a temporary fileTemporary_file "Temporary file" or directory                         |
| mv       | File utilities  | [mv](https://github.com/rendick/glcoreutils/blob/main/src/mv.go)                             |
| realpath | File utilities  | Returns the resolved absolute or relative path for a file                                    |
| rm       | File utilities  | [rm](https://github.com/rendick/glcoreutils/blob/main/src/rm.go)                             |
| rmdir    | File utilities  | Removes empty directories                                                                    |
| shred    | File utilities  | Overwrites a file to hide its contents, and optionally deletes it                            |
| sync     | File utilities  | Flushes file system buffers                                                                  |
| touch    | create file     | [touch](https://github.com/rendick/glcoreutils/blob/main/src/touch.go)                       |
| truncate | File utilities  | Shrink or extend the size of a file to the specified size                                    |
| vdir     | File utilities  | Is exactly like "`ls -l -b`". Files are by default listed in long format.                    |
| b2sum    | Text utilities  | Computes and checks BLAKE2bBLAKE2b "BLAKE2b" message digest                                  |
| base32   | Text utilities  | Encodes or decodes Base32Base32 "Base32", and prints result to standard output               |
| base64   | Text utilities  | Encodes or decodes Base64Base64 "Base64", and prints result to standard output               |
| cat      | Text utilities  | [cat](https://github.com/rendick/glcoreutils/blob/main/src/cat.go)                           |
| cksum    | Text utilities  | Checksums IEEE Ethernet CRC-32 and count the bytes in a file.                                |
| comm     | Text utilities  | Compares two sorted files line by line                                                       |
| csplit   | Text utilities  | Splits a file into sections determined by context lines                                      |
| cut      | Text utilities  | Removes sections from each line of files                                                     |
| expand   | Text utilities  | Converts tabs to spaces                                                                      |
| fmt      | Text utilities  | Simple optimal text formatter                                                                |
| fold     | Text utilities  | Wraps each input line to fit in specified width                                              |
| head     | Text utilities  | Outputs the first part of files                                                              |
| join     | Text utilities  | Joins lines of two files on a common field                                                   |
| md5sum   | Text utilities  | Computes and checks MD5MD5 "MD5" message digest                                              |
| nl       | Text utilities  | Numbers lines of files                                                                       |
| numfmt   | Text utilities  | Reformat numbers                                                                             |
| od       | Text utilities  | Dumps files in octal and other formats                                                       |
| paste    | Text utilities  | Merges lines of files                                                                        |
| ptx      | Text utilities  | Produces a permuted indexPermuted_index "Permuted index" of file contents                    |
| pr       | Text utilities  | Converts text files for printing                                                             |
| sha1sum  | Text utilities  | Computes and checks SHA-1SHA-1 "SHA-1"/SHA-2SHA-2 "SHA-2" message digests                    |
| shuf     | Text utilities  | generate random permutations                                                                 |
| sort     | Text utilities  | sort lines of text files                                                                     |
| split    | Text utilities  | Splits a file into pieces                                                                    |
| sum      | Text utilities  | Checksums and counts the blocks in a file                                                    |
| tac      | Text utilities  | Concatenates and prints files in reverse order line by line                                  |
| tail     | Text utilities  | Outputs the last part of files                                                               |
| tr       | Text utilities  | Translates or deletes characters                                                             |
| tsort    | Text utilities  | Performs a topological sortTopological_sort "Topological sort"                               |
| unexpand | Text utilities  | Converts spaces to tabs                                                                      |
| uniq     | Text utilities  | Removes duplicate lines from a sorted file                                                   |
| wc       | Text utilities  | Prints the number of bytes, words, and lines in files                                        |
| arch     | Shell utilities | Prints machine hardware name same as uname -m                                                |
| basename | Shell utilities | Removes the path prefix from a given pathname                                                |
| chroot   | Shell utilities | Changes the root directory                                                                   |
| date     | Shell utilities | [date](https://github.com/rendick/glcoreutils/blob/main/src/date.go)                         |
| dirname  | Shell utilities | Strips non-directory suffix from file name                                                   |
| du       | Shell utilities | Shows disk usage on file systems                                                             |
| echo     | Shell utilities | Displays a specified line of text                                                            |
| env      | Shell utilities | Displays and modifies environment variablesEnvironment_variable "Environment variable"       |
| expr     | Shell utilities | Evaluates expressions                                                                        |
| factor   | Shell utilities | Factors numbersInteger_factorization "Integer factorization"                                 |
| false    | Shell utilities | Does nothing, but exits unsuccessfully                                                       |
| groups   | Shell utilities | Prints the groupsGroup_identifier_Unix "Group identifier Unix" of which the user is a member |
| hostid   | Shell utilities | Prints the numeric identifier for the current host                                           |
| id       | Shell utilities | [id](https://github.com/rendick/glcoreutils/blob/main/src/id.go)                             |
| link     | Shell utilities | Creates a linkHard_link "Hard link" to a file                                                |
| logname  | Shell utilities | Print the user's login name                                                                  |
| nice     | Shell utilities | Modifies schedulingScheduling_computing "Scheduling computing" priority                      |
| nohup    | Shell utilities | Allows a command to continue running after logging out                                       |
| nproc    | Shell utilities | Queries the number of active processors                                                      |
| pathchk  | Shell utilities | Checks whether file names are valid or portable                                              |
| pinky    | Shell utilities | A lightweight version of fingerFinger_protocol "Finger protocol"                             |
| printenv | Shell utilities | Prints environment variablesEnvironment_variable "Environment variable"                      |
| printf   | Shell utilities | Formats and prints data                                                                      |
| pwd      | Shell utilities | [pwd](https://github.com/rendick/glcoreutils/blob/main/src/pwd.go)                           |
| readlink | Shell utilities | Displays value of a symbolic linkSymbolic_link "Symbolic link"                               |
| runcon   | Shell utilities | Run command with specified security context                                                  |
| seq      | Shell utilities | Prints a sequence of numbers                                                                 |
| sleep    | Shell utilities | Delays for a specified amount of time                                                        |
| stat     | Shell utilities | Returns data about an inodeInode "Inode"                                                     |
| stdbuf   | Shell utilities | Controls buffering for commands that use stdio                                               |
| stty     | Shell utilities | Changes and prints terminal line settings                                                    |
| tee      | Shell utilities | Sends output to multiple files                                                               |
| test     | Shell utilities | Evaluates an expression                                                                      |
| timeout  | Shell utilities | Run a command with a time limit                                                              |
| true     | Shell utilities | Does nothing, but exits successfully                                                         |
| tty      | Shell utilities | Prints terminalComputer_terminal "Computer terminal" name                                    |
| uname    | Shell utilities | [uname](https://github.com/rendick/glcoreutils/blob/main/src/uname.go)                       |
| unlink   | Shell utilities | Removes the specified file using the `unlink` function                                       |
| uptime   | Shell utilities | Tells how long the system has been running                                                   |
| users    | Shell utilities | Prints the user names of users currently logged into the current host                        |
| who      | Shell utilities | [who](https://github.com/rendick/glcoreutils/blob/main/src/who.go)                           |  
| whoami   | Shell utilities | [whoami](https://github.com/rendick/glcoreutils/blob/main/src/whoami.go)                     |
| yes      | Shell utilities | [yes](https://github.com/rendick/glcoreutils/blob/main/src/yes.go)                           |
