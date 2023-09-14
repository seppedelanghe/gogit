# GoGit

GoGit is a cli tool to quickly switch between github _profiles/accounts_. \
Currently only works with very specific SSH `config` setup (see below) and `github`.

## Problem

I found it annoying having 3 GitHub account at once and having to use special `Host` values in my SSH config like:
```
Host github.com-work
    Hostname github.com
    Port 22
    User git
    IdentityFile ~/.ssh/id_a

Host github.com-personal
    Hostname github.com
    Port 22
    User git
    IdentityFile ~/.ssh/id_b

Host github.com-other
    Hostname github.com
    Port 22
    User git
    IdentityFile ~/.ssh/id_b
```

And then I had to alter remote url like this to have the correct permission to the repo: \
`git remote add origin git@github.com-work:username/repo.git`

## Solution

To fix this issue I am building a cli too that can manage my profiles/accounts for me. When I arive at work I run 1 single command to switch to my work git account and everything is setup as it should. When I get home, I just switch accounts again with 1 command.

## Building

### Requirements
- git
- Go 1.20 or up

After cloning the repo, you can build and move the binary `bin/gogit` to a place in your `PATH`. \
I prefer to use `/usr/local/bin/`

__Build binary__
```
make build
```

__Install on MacOS or Linux__ \
Builds binary, moves it to `/usr/local/bin/` and runs `gogit init`.
```bash
make install
sudo make install # if you require root access for /usr/local/bin/
```

## Commands

### `gogit add`

Add a new git account to the GoGit config

### `gogit set`

Set the current active account

### `gogit init`

Create a blank GoGit config file

### `gogit active`

See the current active account

## Features to add

- [ ] Option to add new block to ssh config on `gogit add`
- [ ] Add `gogit remove` command to remove a profile/account
- [ ] Add support for other git remotes
