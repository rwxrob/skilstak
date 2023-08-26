# 2023 Boost overview

Here's what we currently plan to cover. We always start on May the 4th and plan to take approximately 34 weeks with an estimated completion date sometime around January 13, 2024 based on approximately two hours per week (the least we have ever done per week). (Estimate is updated regularly and dynamically from the duration estimated below.)

!!! note "Consider waiting until we are done"

    Some people may find it better to simply wait until the entire 2023 Boost has completed so they can work as fast as they want through the material at their own pace---particularly if you are impatient and have a hard time with the informal, often disorganized way we work through the content.

## Manage your learning

* Become an autodidact: learn to manage your own learning. (1h)
* Take efficient, searchable notes with Markdown on GitHub. (1h)
* Use search-centric system navigation on Windows, Mac, or Linux. (20m)
* Grok benefit and origin of terminal command-line interface. (1h)

## Get Linux

* Install and use bash shell on Windows and Mac (besides Linux). (1h)
* Know why Ubuntu, RedHat, and Alpine are most relevant Linux distros. (20m)
* Create and access local Ubuntu Linux Server virtual machine with VMware. (2h)
* Use secure shell (ssh) to safely interact with remote systems. (1h)
    * Grok public key cryptography
* Do things as administrator (root) but only when necessary. (1h)
* Install, and update software packages with apt. (1h)

## Find your way around

* Get help for any Linux command without the Internet. (20m)
* Use tab-completion to speed command line input. (20m)
* Navigate Linux file system and understand general organization. (2h)

## Command power

* Capture, append, and connect command input, output, and errors. (1h)
* Run commands on multiple files using wildcard globbing. (1h)
* Search and filter text input and files with regular expressions. (2h)
* Edit text files in-place from the command line. (1h)

## Edit text like a pro

* Edit text files visually using major terminal editor apps. (4h)
* Customize ed/ex/vi/vim configuration. (1h)

## Move to multiplexer

* Manage terminal windows with screen/tmux. (1h)
* Customize TMUX configuration. (1h)

## Get answers from the Internet without leaving terminal

* Grok Internet networking. (1h)
    * History
        * <https://youtu.be/qvUYPm2nVXM>
    * Ethernet
        * Routing
           * Local Area Network (LAN)
           * Wide Area Network (WAN)
           * `tracepath`, `traceroute`, `pathping`, `tracert`, `mtr`
           * Mac address, Ethernet address (machine address)
           * Dynamic Host Control Protocol (DHCP)
           * Network Interface Controller (NIC)
        * Ports
           * Pub door analogy
           * Reserved ports
           * `netstat`, `nmap`, `lsof`
        * Network Address Translation (NAT)
    * Firewalls
    * IP
    * TCP/IP
    * UDP
    * Domain names
        * Registrars
        * Discovery
            * `whois`
            * `nslookup`, `dig`
            * <https://shodan.io>
        * localhost 127.0.0.1
        * Reserved
        * `/etc/hosts`
    * HTTP, HTTPS (TLS)
    * Secure Shell
        * `~/.ssh/config`
    * Beware of legacy: FTP, Telnet, Gopher, IPP
* Know the common network query commands. (30m)
    * `dig`, `nslookup`
    * `ifconfig`, `ipconfig`, `ip`, `ss`
    * `ping`
    * `nmap`
    * `ncat`/`nc`
    * `iptables`, `ufw`
* Do effective research from command line with lynx or w3m. (1h)
* Use `curl` and `wget` to fetch data and files from command line. (20m)
* Leverage command history using terminal editor actions to save time. (20m)
* Customize lynx configuration. (1h)
* Send AI (ChatGPT) queries from the command line. (1h)

## Find files (and vulnerabilities) like a hacker

* Find specific files anywhere and determine what they are. (1h)
* Understand and manipulate Linux file ownership and permissions. (1h)
* Use symbolic links. (10m)

## Start building your terminal toolkit

* Grok the benefit of creating and maintain your own, portable toolkit. (20m)
* Organize commands and start your own "dot files" GitHub repo. (1h)
* Organize git repositories. (20m)
* Save source to GitHub with `git` and `gh`. (1h)
* Understand implications of different shell scripting languages. (30m)
* Write first POSIX-compliant shell script. (30m)
* Check your scripts are safe with `shellcheck`. (20m)
* Automatically format shell scripts with `shfmt`. (20m)
* Grok bases, binary, octal, decimal, and hexadecimal. (20m)
* Manage processes (programs currently running). (1h)
* Grok UNIX philosophies and use UNIX filters over plugins. (20m)
* Understand ASCII, terminal encoding, and escape sequences. (30m)
* Customize bash profile and environment. (1h)
* Grok and apply everything in the bash manual page (man bash). (4h)
* Do math in the shell safely and reliably with bc and dc. (1h)
* Parse JSON with jq. (1h)
* Parse YAML with yq. (1h)
* Master parsing other delimited formats. (30m)
* Try a little Python. (2h)
* Try a little Go. (2h)

## Learn just enough Web dev

* Know how much Web dev is expected from all techies (20m)
* Learn enough HTML to grok Markdown rendering (1h)
* Learn enough CSS to make rendered Markdown pretty (1h)
* Learn enough vanilla JavaScript to add occasional widget (1h)
* Don't bother learning any "frameworks" (for now) (20m)
* Use only the best Web dev learning resources (20m)
* Understand PWAs and what it means to be "progressive" (20m)
* Use `zet` (or something like it) to publish your live notes (1h)
* Create a static site generator with `pandoc` and 10-lines of bash (1h)
* Publish to Web automatically with GitHub pages (30m)
* Publish to Web automatically with Netlify (30m)
* Learn to leverage Go text and HTML templates (30m)

## Become employable

* Grok that employment is about trust and people, not positions. (20m)
* Network and foster a professional learning network/community. (1h)
    * Grow neck beard (where possible) lol.
    * "Strong opinions, weakly/weekly held."
* Start listening for opportunities now. (20m)
* Forget about LinkedIn. (20m)
* Complete GitHub profile page and grok job discovery from GitHub. (1h)
* Create and maintain a consumable, sustainable resume. (2h)
    * Don't call yourself a "junior"
    * Job titles are mostly bullshit (too bad no role-based employment)
* Target occupations you want and start applying. (1h)
    * How much salary do I actually *need*?

## Have a strategy for what comes next

* Know when and how to certify and why you might want to add it. (1h)
* Set healthy work boundaries and self-care habits. (1h)

## Pay it forward! (1h)

* Remember where you were. (20m)
* Become a mentor. (20m)
* Never stop learning. (20m)

## What *won't* we learn (here)?

There's a lot of stuff to learn and eventually we do hope to cover it in one of the other learning series we have planned. This is a place to capture that stuff when we think of it.

### Fullmetal Linux

* How to install Linux on hardware (laptops, server towers, old and new).
* How to use Linux graphic desktops and window managers

### Polyglot Programming

* How to create and work with containers
* How to create micro-services
* How to interact with Web APIs

### Homelab Init

* How to setup a local VLAN
* How to setup and manage own DNS server
* How to safely tunnel web-facing services through external providers
* How to setup on-prem vanilla Kubernetes cluster
* How to setup local virtual machine hosting with QEMU/libvirt
