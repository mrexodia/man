# man

Wrapper for the [Bwana](https://www.bruji.com/bwana) man page viewer. Based on a [blog post](https://jonathansblog.co.uk/open-man-pages-in-safari-or-other-browsers-in-osx) by Jonathan Mitchell.

## Installation & Usage

1. Download and install [Bwana](https://www.bruji.com/bwana) ([direct link to DMG](https://www.bruji.com/bwana/bwana.dmg))
2. [Install go](https://golang.org/doc/install) and make sure `$GOPATH/bin` is in your `PATH` **before** `/usr/bin`
3. `go get github.com/mrexodia/man`
4. `man ls` or `git init --help` should open the relevant manpage in your default browser.