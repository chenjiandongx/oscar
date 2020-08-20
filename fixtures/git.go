package fixtures

var pullConsole = `
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

var statusConsole = `
git status
On branch oscar-git
Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git checkout -- <file>..." to discard changes in working directory)

        modified:   git.go

no changes added to commit (use "git add" and/or "git commit -a")

`

var addConsole = `
git add git.go
warning: LF will be replaced by CRLF in fixtures/git.go.
The file will have its original line endings in your working directory.
`



var GitPullConsole = readLines(pullConsole)
var GitStatusConsole = readLines(statusConsole)
var GitAddConsole = readLines(addConsole)
