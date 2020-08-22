package fixtures

var gitConsole = `
# How to contribute to the oscar project?

# 0x00: clone this project
> cd ~/git
> git clone https://github.com/chenjiandongx/oscar.git
Cloning into 'oscar'...
remote: Enumerating objects: 101, done.
remote: Counting objects: 100% (101/101), done.
remote: Compressing objects: 100% (71/71), done.
remote: Total 101 (delta 53), reused 77 (delta 29), pack-reused 0
Receiving objects: 100% (101/101), 244.37 KiB | 23.00 KiB/s, done.
Resolving deltas: 100% (53/53), done.

# 0x01: upgrade something
> echo "what an awesome project like oscar" >> README.md

# 0x02: display what you have changed
> git status
On branch oscar-git
Changes not staged for commit:

  (use "git add <file>..." to update what will be committed)
  (use "git checkout -- <file>..." to discard changes in working directory)

        modified:   README.md

no changes added to commit (use "git add" and/or "git commit -a")

# 0x03: add commit
> git add . && git commit -m "update readme"
warning: LF will be replaced by CRLF in fixtures/git.go.
The file will have its original line endings in your working directory.
[oscar-git da3d630] update git.go
1 file changed, 24 insertions(+), 2 deletions(-)

# 0x04: push it to the remote.
> git push
Enumerating objects: 11, done.
Counting objects: 100% (11/11), done.
Delta compression using up to 12 threads
Compressing objects: 100% (5/5), done.
Writing objects: 100% (6/6), 552 bytes | 552.00 KiB/s, done.
Total 6 (delta 3), reused 0 (delta 0)
To https://github.com/chenjiandongx/oscar.git
   0f96015..7d14615  master -> master

# 0x05: pr welcome
> echo "oscar!"
`

var GitConsole = readLines(gitConsole)
