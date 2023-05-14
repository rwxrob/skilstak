# homelab overview

> "I've hired many engineers and would easily hire someone who has built and maintains their own home lab over any possible certifications." (executive hiring manager for large tech corporation)

Advantages:

* Solid, hands-on experience
* Includes hardware, networking, operations, and development skills

## Stages to homelab development

1.  Start with old machine and install Linux/UNIX to breath new life.
    1.  Consider defunct, refurbished school lab computers.
    2.  Consider single-board computers (SBC) for sensor-integration and AI.
    3.  Take advantage of the “hype effect” and buy refurbished over new.
2.  When enough computers invest in a router and build home network.
3.  Save money and buy a refurbished server to host virtual machines.
4.  Save more money and buy multiple, small computers to build HW/HA cluster.

## What about power?

* Invest in a good power meter.
* Secret to managing power is having metered LED PDU (don't skimp on PDU).

## How much do I have to spend?

## Virtualisation considerations

* KVM is master of the universe
* QEMU emulates everything, not just Linux, but bloated a bit
* Firecracker replaced QEMU when small VMs wanted
* Firecracker is master of MicroVMs, but doesn't emulate everything
