package fixtures

var console = `
"cd ~/git"
"git clone https://github.com/chenjiandongx/oscar.git"
"Cloning into 'oscar'..."
"remote: Enumerating objects: 101, done."
"remote: Counting objects: 100% (101/101), done."
"remote: Compressing objects: 100% (71/71), done."
"remote: Total 101 (delta 53), reused 77 (delta 29), pack-reused 0"
"Receiving objects: 100% (101/101), 244.37 KiB | 23.00 KiB/s, done."
"Resolving deltas: 100% (53/53), done."
`

var GitConsole = readLines(console)
