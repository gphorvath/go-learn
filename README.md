# My Project for Learning Go

## Setup goenv

Goenv simplifies managing the current active version of Go at the local and global levels.

Follow instructions for specific OS here: <https://github.com/go-nv/goenv/blob/master/INSTALL.md>

Once goenv is installed, you should be able to configure the version of go from the terminal. For purposes of learning, let's start with `latest`.

```bash
goenv install latest    # installs the latest version of go
goenv global latest     # sets the global version of go to the latest
goenv global            # to confirm it's set

# you may need to refresh your terminal (aka `source ~/.zshrc` or `source ~/.bashrc`)
go version              # to confirm go is on the path
```

From this point forward, if you need to set a different version of go, consider using `goenv local <version>` to avoid changing the go version for other projects.
