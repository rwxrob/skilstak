# 2023 Boost - Week 3

📺 <https://youtu.be/mUmQk1o6FFs>

(PLEASE SKIP TO WEEK 4.) This week we setup an Ubuntu Linux container image using podman command and apt package manager, learn how to save (commit) an image copy so we have backups during our setup, and how to navigate the command line.

1.  Why Ubuntu?
2.  Pull an Ubuntu Linux image into podman images
3.  Create (run) a named container from the Ubuntu local image
4.  Exit the container
5.  Understand difference between running and exited containers
6.  Use –rm for temporary containers
7.  Use podman ps -a to show all containers (running and exited)
8.  Use podman rm to remove unwanted containers
9.  Use bash tab completion to save on typing
10. Use up and down arrows (for now) to get previous commands
11. Restart and attach to named container
12. Use Advanced Package Manager (APT)
13. Use update command to refresh list of packages
    1.  Do NOT use upgrade command
14. Clear the screen with clear command (not control-L)
15. Install documentation helpers
16. Use man to get help about commands
17. Use help to get help about bash builtins
18. Look for runnable stuff
    1.  Use which command
    2.  Use type builtin
    3.  Use command -v command
    4.  Use file command to see what type of thing it is
19. Everything is a file (inode) even directories
    1.  Use ls command
    2.  Use pwd command
    3.  Use cd to change directories
    4.  Use chmod to change permissions
    5.  Use chown to change owner (and group)
20. Add new user with adduser (or useradd)
21. Delete user with deluser (or userdel)
22. Add new group with addgroup (or groupadd)
23. Delete group with delgroup (or groupdel)
24. Commit (save) container as image
25. Push saved image to GitHub Container Registry
    1.  Create Personal Access Token on GitHub
    2.  Use podman push to push local image to remote ghcr.io
    3.  View saved GitHub packages (containers)
