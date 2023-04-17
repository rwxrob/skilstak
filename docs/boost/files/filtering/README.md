# How can I change the contents of a file using nothing but commands?

Before learning how to use one of the standard terminal editors you really should first learn how to cope with a terminal that cannot be run in **visual** mode at all. In fact, in the 70s there was no visual mode. Everything was accomplished with commands chained together in what would become known as **pipelines**.

!!! danger

     We all saw you yawn reading that "boomer" pipeline reference. 😉 But hold up a minute before your scroll past it.

     Consider that most hackers use commands for changing file content all the time because their backdoor malware only allows commands to be run without a "full visual interface" that requires additional device access, which can trigger monitoring software watching for use of those devices. So, if you are a potential hacker who doesn't want to get caught it's best you learn how to change what's in a file without an posh editor. That's right, vi *is* posh in such cases.

     If you didn't understand any of that, you will soon enough.

One significant advantage of changing file content using a command pipeline (instead of an editor) is that those commands can be combined with other commands in order to setup automations triggered by events or scheduling. So even though authoring book-like content (such as this) is better from an editor, there are many very real reasons to master file manipulation with command pipelines and filters, *first*. Plus, you'll be able to use those skills directly from within certain editors to supercharge your terminal visual editing sessions as well. People will watch you in agonizing envy as you perform terminal editing magic that their stupid, bloated graphic text editors can't even attempt.
