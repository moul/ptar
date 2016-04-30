# ptar
:ant: parallel tar (SMP)

## Usage

```console
$ ptar -h
Usage:
  ptar [OPTIONS]

Application Options:
  -C, --chunks=   Amount of chunks.
      --affinity= If set, ptar will only archive the files of a specific affinity.
                  Else, it will create create all the archives. (default: -1)
  -c              Create new archives containing the specified items.
  -f=             Read the archive from or write the archive to the specified file.
                  The filename can be - for standard input or standard output.

Help Options:
  -h, --help      Show this help message
```

## Example

Run 4 instances of ptar

```console
$ ptar -C 4 -cf /tmp/test.tar .
make: Nothing to be done for `all'.
INFO[0000] child(0): running ptar [--affinity=0 -c -f /tmp/test.tar.0 -C 4 .]
INFO[0000] child(1): running ptar [--affinity=1 -c -f /tmp/test.tar.1 -C 4 .]
INFO[0000] child(2): running ptar [--affinity=2 -c -f /tmp/test.tar.2 -C 4 .]
INFO[0000] child(3): running ptar [--affinity=3 -c -f /tmp/test.tar.3 -C 4 .]
INFO[0000] + .git/HEAD
INFO[0000] + .git/hooks/applypatch-msg.sample
INFO[0000] + .git/hooks/pre-commit.sample
INFO[0000] + .git/hooks/update.sample
INFO[0000] + .git/logs/refs/heads/master
INFO[0000] + .git/logs/refs/remotes/origin/master
INFO[0000] + .git/objects/f1/9ca84ffa79c01e195a6150ab841c233989d961
INFO[0000] + .git/objects/pack/pack-3a210da928ef49fde1c519dde1cd6d2c22a9d4a6.pack
INFO[0000] + .git/refs/heads/master
INFO[0000] + .git/refs/remotes/origin/master
INFO[0000] + Makefile
INFO[0000] + cmd/ptar/main.go
INFO[0000] + .git/config
INFO[0000] + .git/hooks/commit-msg.sample
INFO[0000] + .git/hooks/pre-push.sample
INFO[0000] + .git/index
INFO[0000] + .git/logs/HEAD
INFO[0000] + .git/objects/06/4817a0756d36a1d097255f9780c582ff85795d
INFO[0000] + .git/objects/24/5dcc4af163e13ff5e5671bf68dae4e9fe81f96
INFO[0000] + .git/objects/55/e4755b7266a052d2b6aa2baff540f14035c61e
INFO[0000] + .git/objects/6f/ae121fc52c7c2ab7bf9b5029975f7116108528
INFO[0000] + .git/objects/8d/82a4e03cd942f7233835d461b05678782b0a87
INFO[0000] + .git/objects/d5/d0ca5bb14d5d88d3f1d062d6e6817d6ed585c6
INFO[0000] + .git/packed-refs
INFO[0000] + README.md
INFO[0000] + ptar
INFO[0000] + .git/description
INFO[0000] + .git/hooks/post-update.sample
INFO[0000] + .git/hooks/pre-rebase.sample
INFO[0000] + .gitignore
INFO[0000] + .git/COMMIT_EDITMSG
INFO[0000] + .git/hooks/pre-applypatch.sample
INFO[0000] + .git/hooks/prepare-commit-msg.sample
INFO[0000] + .git/info/exclude
INFO[0000] + .git/logs/refs/remotes/origin/HEAD
INFO[0000] + .git/objects/00/1143af1d8ce1dc8612277e368a68d00e05b887
INFO[0000] + .git/objects/14/cf8accc94a56058c727003ee2803625042915e
INFO[0000] + .git/objects/31/e37136b071f6bd6fdec012606287b5a36231b8
INFO[0000] + .git/objects/5d/45ae9cf0e179aaa73cd1d39b2c224db050736e
INFO[0000] + .git/objects/73/a7b524f111ebe9acbe2bbd6ce5b7f061ebab06
INFO[0000] + .git/objects/95/5ab86c5937bd1107cd5c99caeb6a0588d27943
INFO[0000] + .git/objects/f1/7f6d1ee8f09de64d4b126625d33bdc208a3ff6
INFO[0000] + .git/objects/pack/pack-3a210da928ef49fde1c519dde1cd6d2c22a9d4a6.idx
INFO[0000] + .git/refs/remotes/origin/HEAD
INFO[0000] + LICENSE
```

## License

MIT
