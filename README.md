# GoGit

GoGit is a cli tool to quickly switch between github _profiles/accounts_. \
Currently only works with very specific SSH `config` setup (see below).

## Disclamer

This tool is under development and can cause corrupt files. __Please use with caution!__ \
Before overwriting a file (e.g. `~/.ssh/config` or `~/.gitconfig`), GoGit will make a copy of the previous file. \
This file will have the same name as the original but with `.bak` appended to the end. \
For example `.gitconfig` => `.gitconfig.bak`, these files can come in handy to restore the old file when a corrupt file is produced.

## Problem

I found it annoying having 3 GitHub accounts at once and having to use special `Host` values in my SSH config like this:
```ssh-config
# Work account (default)
Host github.com
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

And then I had to alter the remote url for each local repo to have the correct permissions to the remote repo: \
`git remote add origin git@github.com-work:username/repo.git`

After not finding a nice solution for this problem, I figured I could build one as a side project.

## Solution

To fix this issue I am building a cli too that can manage my git profiles/accounts for me. \
When I arrive at work, I run 1 single command to switch to my work git account and everything is setup as it should. \
When I get home, I just switch accounts again with 1 command.

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
Builds binary and moves it to `/usr/local/bin/`
```bash
make install
sudo make install # if you require root access for /usr/local/bin/
```

## Commands

### `gogit add <name> <email?> <username?> <remote?>`

Add a new git account to the GoGit config with the following arguments:
- `name`: 
    - custom name for profile
- `email`: 
    - (optional) email of your git account
    - default = current email in `~/.gitconfig`
- `username`:
    - (optional) name of your git user
    - default = current name in `~/.gitconfig`
- `remote`: 
    - (optional) custom remote name
    - default = `github.com`

<hr>

### `gogit set <name>`

Set the current active account. 

GoGit will only change the `Host` values in your SSH config. \
For example: `github.com-<name>` to `github.com` and back. \
All other settings will be left unchanged.

<hr>

### `gogit init`

Create a blank GoGit config file

<hr>

### `gogit active`

See the current active account

<hr>

### `gogit remove <name>`

Remove a profile/account from GoGit, leaves your SSH config untouched.

<hr>

### `gogit list`

Get a list of the current configured profiles/accounts with the arrow `->` indication the active one.
```bash
profiles:
-> home
-  work
-  other
```

<br>

## SSH config

GoGit currently requires you to have your SSH config file setup up in a specific way. This needs to be done manually for now. \
See the [problem](#problem) section for an example structure of a working `~/.ssh/config` file.

## GoGit config file

__Located at__: `~/gogit.ini`

```ini
[profiles]
first   = true
second  = false

[preferences]
# enable if you want to auto update your global git config user and email
set-git-user = false

[profile.a]
remote   = github.com
username = User a
email    = user.a@email.com

[profile.b]
remote   = github.com
username = User b
email    = user.b@email.com
```

- Setting `set-git-user` to `true` will make GoGit alter your `~/.gitconfig` file
- Each profile has unique settings

## Features to add

- [ ] Option to atuo add new block to ssh config on `gogit add`
- [ ] Add `gogit remove` command to remove a profile/account
- [x] Add support for other git remotes
- [x] Setting `git config --global` when activating account
