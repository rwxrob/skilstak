# 2023 Boost - Week 15

📺 <https://youtu.be/LhMj51fpqeQ>

* What's a shell and why do I care?
* Which shell should I use?
* What is the difference between an interactive shell and shell scripting?
* What is the difference between GPLv2 and GPLv3 licenses?
* Everyone should use `shellcheck`.
* Dynamic `shellcheck` with Vim Ale.
* Getting started coding POSIX shell and bash.
* Coding your own `tripwire` program to catch hackers (in 20 lines of shell code).
* The find command, perhaps the least appreciated in all of UNIX.

Here's our example from the session. Play around with it to get good with `find` and regular expressions.

```sh
#!/bin/sh

rootdir='/'

exclude=\
'container|cache|'\
'\/[a-f0-9]{3,}_0|'\
"^$HOME"\
'/.(bash_history|viminfo|.cache/google)'

include='^/(etc|dev)'

# Use find to get all files that have changed since a specific time.
#   0 mean last 24 hours
#   1 means 24-48 hours

find "$rootdir" -type f -mtime 0 2>/dev/null | \
  grep -E "$include" | \
  grep -E -v "$exclude"
```
