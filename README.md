## Sussay

Cowsay but Amogus

The code is horrible. So is making ASCII art.

Q: Do I regret making this?
A: Not a single bit.

Q: Why?
A: I was sitting in an exam sctibbling an amogus onto the last page and wanted to give it a speech bubble. That's when i recalled there is `cowsay`, `whalesay` but no `sussay` yet. 

## Build Instructions

You will need to have go installed.

Run the following snipped to compile and install it.

```sh
go build -o sussay ./src
chmod ugo+x sussay && chown root:root sussay
cp sussay /usr/bin/sussay
```
