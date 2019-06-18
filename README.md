# Gsync

Gsync is a simple tool to work more efficiently with multiple git repository.

The goal of gsync is to offer a better developer experience when usage of a single git mono-repo is impossible.

Gsync let you reference a project containing multiple git repository and then :
* Clean up all git repo from unstaged changed and checkout to master


## Install

Gsync is a golang project using go modules. The only step to install gsync is to build it using go >= 1.11  

```bash
go build
```

## Usage

For usage example, feel free to have a look on [the example folder](examples/getting-started/)