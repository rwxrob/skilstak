# 2023 Boost - Week 4

📺 <https://youtu.be/SPoqW1CMEhY>

Let's do this again. Weeks 2 and 3 were a little harder than most beginners might be able to handle. Most Linux users will be using a Linux that has been installed already. The Boost is primarily intended for *them* (not admins and hackers, who are also users). So we've prepared a quick-start with an existing Linux container image. Here's how to get it and another take at how to get going without having to setup anything.

1. What concepts do all *users* need to understand?
1. What does it mean to be a Linux "user" (vs "admin" or "hacker")?
1. Create a local Linux container for the Boost
    * `podman run -it --hostname skilstak --name boost -v shared://shared ghcr.io/rwxrob/ws-skilstak`
1. What is the difference between a terminal and a command line?
1. What do the parts of the prompt mean?
1. How do I clear the screen?
    * Do not depend on Control-L
    * `clear` (or `c`)
1. How do I see where I am?
    * `pwd`
1. How do I see what is in this directory?
    * `ls -al`
1. How do get help information about a command?
    * `CMD -h` or `CMD --help`
    * `man CMD`
    * `help CMD`
    * `apropos KEYWORD` (if you want)
1. What is a directory and what is a file?
1. What is a "parent" and "sub" directory?
    * current directory: `.`
    * parent directory: `..`
    * previous directory: `-`
1. How do I change directories?
    * `cd DIR`
    * `cd -`
    * `cd ..`
    * `cd`
1. How do I see what's in a file?
    * `cat FILE [FILE]`
    * `tac FILE [FILE]`
    * `more FILE`
    * `less FILE`
    * `head FILE`
    * `tail FILE`
1. What is a file descriptor and how do I use them?
    * Understand origins of TTY (teletype machines)
    * standard input (`/dev/stdin`, 0)
    * standard output (`/dev/stdout`, 1)
    * standard error (`/dev/stderr`, 2)
1. What is a pipe?
    * `ls -al | more`
1. How do write output of a command to a file?
    * `ls -al > /tmp/foo`
1. How do I append command output to end of a file?
    * `echo some >> /tmp/foo`
    * `echo thing >> /tmp/foo`
1. How do I zero out an existing or new file?
    * `> /tmp/foo`
1. How do I create a new file (or update time changed)?
    * `touch /tmp/foo`
1. How do I remove a file?
    * `rm FILE`
1. How do I find files with a keyword in the name?
    * `find . -name '*vim*' 2> /dev/null`
1. How do I send stderr to stdout as well?
    * `find / -name '*md*' 1> /tmp/foo 2> /dev/stdout`
    * `find / -name '*md*' > /tmp/foo 2> /dev/stdout`
    * `find / -name '*md*' > /tmp/foo 2>&1`
    * `find / -name '*md*' &> /tmp/foo`
