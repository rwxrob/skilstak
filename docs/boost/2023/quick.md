# 2023 Boost - Quick Start

📺 <https://youtu.be/Nfq26YHssj8>

Start learning Linux and containers now. Whether you are coming into the Boost late and want to catch up to keep up with live streams or you just want to get Linux on your local computer as fast as possible here is the fastest way to get started. (If you want to see all three hours of this ws-skilstak container image being created see the related video in the same playlist.)

1.  Install terminal software
    -   Mac: iTerm2
    -   Windows: Terminal Preview
2.  Install bash
    -   Mac: install brew (then) brew install bash
    -   Windows: install git-bash
3.  Install Podman Desktop
4.  Initialize and start the podman virtual machine
    -   Open bash terminal command line
    -   podman machine list
    -   podman machine init
    -   podman machine start
5.  Pull and run a container from SKILSTAK workspace image
    -   podman pull ghcr.io/rwxrob/ws-skilstak
    -   podman run -it –hostname skilstak –name skilstak -v shared://shared ws-skilstak
    -   Use exit to exit container
    -   podman start -a skilstak
