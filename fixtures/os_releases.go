package fixtures

var osRelease = `
5.0.7-arch1-1-ARCH
5.2
2.6.32-431.el6.i686
9.0.2.2
1.5.19(0.150/4/2)
2.2.1(0.289/5/3)
3.1.9+
3.6.11+
8.0.1-amd64
2.13-DEVELOPMENT
6.1-RELEASE-p15
9.0-RELEASE
2.6.34-gentoo-r12
8.11.11
2.6.35-22-generic
2.6.38.8-g2593b11
12.3.0
3.17.3-1-MANJARO
2.6.22.5-31-default
2.6.38-10-generic
2.6.35.7-unity1
`

var OsRelease = readLines(osRelease)
