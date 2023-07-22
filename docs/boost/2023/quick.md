# 2023 Boost - Quick Start

UPDATE: WSL blew up destroying `podman` so we will never be using it again. The Boost returns to using a virtual machine (as in 2022). See [Week 9](../week9) instead of the quick-start below for how to set that up.

----

📺 <https://youtu.be/Nfq26YHssj8>

Start learning Linux and containers now. Whether you are coming into the Boost late and want to catch up to keep up with live streams or you just want to get Linux on your local computer as fast as possible here is the fastest way to get started. (If you want to see all three hours of this ws-skilstak container image being created see the related video in the same playlist.)

1. Install terminal software
    -   Mac: iTerm2
    -   Windows: Terminal Preview
2. Install bash
    -   Mac: install brew (then) `brew install bash`
    -   Windows: install git-bash
3. Install Podman Desktop
4. Initialize and start the podman virtual machine
    -   Open bash terminal command line
    -   `podman machine list`
    -   `podman machine init`
    -   `podman machine start`
5. Pull and run a container from SKILSTAK workspace image

    ```sh
    podman run -it --hostname skilstak --name boost -v shared://shared ghcr.io/rwxrob/ws-skilstak
    ```

6. Use `exit` to exit container
7. Reattach to container later
    -  `podman start -a boost`
