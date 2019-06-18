#Getting Started

This is a good place to test and learn how to use gsync.

## Usage

The `gsync.yml` file describe a project called `getting-started` with 2 git repositories.

By running the `clean` command, Gsync will detect missing reporitories, clone theme, and checkout all repositories to master.

By doing so, gsync will discard all changes that are not staged.

```sh
gsync clean
```

