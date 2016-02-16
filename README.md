# ssedit

ssedit is a small command line utility built on top of scp that allows you to
open and save files / directories on a remote machine using your favourite
text editor (or any program you want really), with ease.

Disclaimer: This is super new, I've only really started working on it, and while
it does work, there are bugs. You've been warned.

Binary distributions are on their way.

### Install

```bash
$ go get -u github.com/hazbo/ssedit/ssedit
```

Providing GOPATH/bin is in your PATH, running the above is all you should need
to do.

### Usage

```bash
$ ssedit open hazbo@example.com /home/hazbo/config.json
```

Since ssedit is built up on top of scp, the above example assumes that you have
access to the given server via ssh keys. The default text editor (that would
open with the above command) is vim. This is easily changed by use of the `-e`
flag, for example:

```
$ ssedit open -e subl hazbo@example.com /home/hazbo/config.json
```

Due to the way in which ssedit saves your edited files to the remote, a session
is started which essentially watches for changes in your files. To exit out of
the session, it's `Ctrl+C`.

By default, remote files do also get stored locally during the session. Once
you're finished with ssedit they should be automatically deleted. However this
program is super new and there might be bugs, so you can manually clear the
locally stored files by doing the following if something goes wrong:

```bash
$ ssedit clear
```

### Contributing

  - Fork ssedit
  - Create a new branch (`git checkout -b my-feature`)
  - Commit your changes (`git commit`)
  - Push to your new branch (`git push origin my-feature`)
  - Create new pull request
