# go-practice

This is a repo containing random programs that I've created, simply for the purpose of learning the [Go programming language](https://go.dev/).

Each subdirectory contains a different project.

# Contents

## /base64

A simple command-line base64 encoder/decoder.

### How to

To encode, provide a single command-line argument string to be encoded:

```bash
$ ./base64 "This text will be encoded!"
$ VGhpcyB0ZXh0IHdpbGwgYmUgZW5jb2RlZCE=
```

To decode, use `-d` or `--decode` followed by the string to be decoded:

```bash
$ ./base64 -d VGhpcyB0ZXh0IHdpbGwgYmUgZW5jb2RlZCE=
$ This text will be encoded!
```