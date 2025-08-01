# OVERVIEW

Welcome to the **SKILSTAK Beginner Boost**!

Since 2020, our community has come together every year to do a series of extremely casual live streams and videos with one mission in mind: to help beginners get good tech jobs starting from absolute zero.

## SKILSTAK app

Planning a PWA with only the following on a simple front page:

- Current announcement
- Last posted video
- Search box
- Three-level, collapsible tags index

## Tags index

I use as system of tagging that works well with YouTube's "three main tags" algorithm. I have decided to use it as a method of categorization with some tags being level one, two, and three. This works out well when using the tags in software to look something up.

Beyond the three tags are additional qualifying tags that distinguish between doing and learning about something, for example, `#coding #golang #bonzai #tutorial` would cover how to create something with the Bonzai framework versus me making something. Here are some qualifiers:

- `#protips` - short discoveries
- `#tutorial` - contains step-by-step way to do something
- `#educational` - contains something to learn
- `#short` - YouTube short under three minutes
- `#rants` - angry stuff
- `#reactions` - honest, immediate, authentic reactions to other stuff
- `#tech` - just 'cuz

### Top level

- `#techlife` - getting, keeping, and enjoying a career in tech
- `#homelab` - building and using a home tech lab for fun and profit
- `#coding` - learning to code and doing it for fun and profit
- `#hacking` - breaking into things for fun and profit
- `#education` - conversations and discoveries about true education
- `#contentcreation` - how to create and maintain content
- `#irl` - me doing unrehearsed, unscripted things IRL

### `#techlife`

- `#learning`
- `#notetaking`
- `#productivity`
- `#chatgpt`
- `#hacking`
- `#getajob`
- `#networking`

### `#homelab`

- `#networking`
- `#linux` - (includes `#unix`)
    - `#terminal`
    - `#scripting`
    - `#bash`
- `#containers`
    - `#podman`
- `#kubernetes`
    - `#containerd`
- `#terraform`
- `#ansible`
- `#proxmox`
- `#minecraft`
- `#hardware`


- `#hacking`
    - `#discovery`
- `#coding`
    - `#beginners`
    - `#gamedev`
    - `#webdev`
    - `#enterprise`
    - `#leetcode`
    - `#bonzai`

- `#education`

BASIC    | [Master basic computer skills](basic.md)
WEB      | [Create a basic web site](web.md)
TERMINAL | [Learn the terminal](term.md)
CODE     | [Learn to code](code.md)
LXC      | [Leverage Linux containers](lxc.md)

## Boost is not a course

Remember that this is just a boost to get you going. We'll shoot for breadth over depth so that you go away at least knowing what you don't know and have the confidence to do stuff to fill in the gaps in your knowledge and mastery. The biggest challenge most learners have is knowing what to learn and why.

The Beginner Boost is not a course. There is no syllabus, no certificate, no credential, no proof you did anything except your own work, notes, and knowledge. There's barely a schedule that changes organically as we go through it.

!!! warning "Start thinking like a hacker"

    Most people need to forget the old, broken ideas about learning, and start learning and thinking like a hacker. People who require the on-a-plate, tell-me-what-to-do form of education generally do very poorly in tech careers. We'll talk more about that during the Boost itself when we discuss the nature of *true* learning.

## Targeted careers

When deciding what to include we like to keep a list of typical job titles you might see out there that require or strongly benefit from the skills, knowledge, and abilities (SKAs for short) covered by the Boost:

* Security Analyst (Hacker)
* Site Reliability Engineer (SRE)
* Cloud Native Engineer
* Machine Learning Engineer
* Systems Engineer
* Platform Engineer
* Infrastructure Engineer
* DevOps Engineer
* Software Developer
* Web Developer
* Computer Scientist
* Computer Engineer
* Robotics Engineer
* Rocket Scientist
* Physicist
* Astronomer
* Any other career involving science and technology

!!! tip "Check out the Bureau of Labor and Statistics technology careers"

    The US BLS updates an annual collection of statistics related to career growth, demand, and pay. It is definitely worth checking out.

    <https://www.bls.gov/ooh/computer-and-information-technology>

    Take particular note of the **Job Outlook** indicating how fast the demand for each career is growing. For example, demand for **Information Security Analysts** (hackers) is growing by 35%, one of the fastest growing professions *of all* outpacing the next fastest growing tech career by 10% (**Software Developers**, 25%)! Beginner Boost skills are absolutely essential for *any* career in the Information Security industry and will give any Software Developer a solid advantage over all the rest.

## Skills format

Content is organized into procedural "how-to"s for specific skills as would be listed in a job description or resume. This provides better searchability, review, and is consistent with popular "recipe" books used by O'Reilly and others.

Explaining procedures in this way is also the basis of learning like a hacker, or a magician for that matter. Each "trick" (or task) is captured with its specific steps. Tricks can then be combined in different ways. Learning this way is much like writing code for your own brain, or programming humans as if each task had a subroutine that can be written and stored in neurons like RAM in a computer.

The Boost learning model, therefore, is designed to help you program your own brain to be able to do things much like an AI or software application. In fact, the only difference is that in order to write the code into your brain you have to repeat the steps of any task over and over to create those neural pathways, which is *exactly* what modern AI training does.

Any skill requires certain knowledge and abilities as well:

* Skills
* Knowledge
* Abilities

Here's how we create a learning segment about a specific skill:

1. Write skill as if a bulleted line in a job description
1. Break down the skill into procedure of tasks
1. Write each task like a skill as well
1. List any ability dependencies for task (typing, fine motor skills, etc.)
1. Give knowledge required for task (terms, context, history, justification)
1. Break down larger skills into smaller ones

## Linux focused

All tech jobs targeted by the Boost significantly benefit from Linux skills. Linux is the operating system (like Mac or Windows) that runs most all of the tech infrastructure in our world today.

### Learn to be a Linux user first (not admin)

Most of the resources about Linux available on the Internet, free and paid, all suffer from the same frustrating flaw: they cover too much assuming the person learning Linux wants to be able to administer a Linux system rather than just benefit from using it.

The Boost content is constantly being simplified and revised with one specific question in mind: Does a Linux user *need* to learn this? The result is simplified content that doesn't go into unnecessary rabbit holes. For example, a machine learning engineer or student using an assigned Linux user account never needs to understand how to create, modify, and delete users and groups on the system. Instead, learning how to co-exist safely with other users on a multi-user system is a more valuable use of time for an absolute beginner.

### Linux learning never ends

One of the best (and some would say worst) things about Linux is that no one ever learns it all. There is always some tip or tweak or new technique. Learning to use Linux is just the first step.
Most of the careers listed above will require further learning skills that build on your new Linux user skills. This is why during the Boost you not only learn the basics of the Linux terminal command line but are also encouraged to learn to organize and design your own autodidactic learning plans. This specific skill is more important than all the others---especially for those seeking careers as knowledge workers of any kind.

How do we learn this in the Boost? By example. When we need to do a bit of research and learning we do it together using the best methods and tools available to any knowledge worker. The Linux terminal command line is objectively the fastest human-computer interface possible for knowledge work. As you master just-in-time learning (or learning at the speed of knowledge workers) then during and after the Boost you can take better control of your own learning plan tailored to your specific needs and goals.

## Please HAVE FUN! 🎉

* [Chat, stickers, clips are all very welcome and appreciated](https://youtu.be/chux1oAhiis)

## Legal usage

Copyright (c) 2013-2024 Robert S. Muhlestein <rob@rwx.gg> Content released under the Creative Commons (BY-NC-ND), code released under the Apache 2.0.

Contributors and project participants implicitly accept the Developer Certificate of Authenticity (DCO) giving over all intellectual property rights to the copyright owner and asserting that they have legal permission to do so.

"SKILSTAK", "SKILSTAK Beginner Boost", "SKILSTAK Boost", “Beginner Boost” and “Boost” are legal trademarks of Robert S. Muhlestein but can be used freely to refer to the this project without limitation. To avoid potential developer confusion, intentionally using these trademarks to refer to other projects---free or proprietary---is prohibited.

The reason for “no derivatives” CC requirement is to preserve the consistency of opinions throughout the content since attribution is required. Without it, forks with changed opinions and resource listings could be purposefully or accidentally taken as the opinions of the original author. This is simply too dangerous to allow. The “no derivatives” clause protects against the inevitable “consensus” problem that plagues community-created content. That said, please reach out by email if you have questions about contributing and collaborating.
