# GoGit

GoGit is a cli tool to quickly switch between github _profiles/accounts_. \
Currently only works with very specific SSH `config` setup (see below) and `github`.

## Problem

I found it annoying having 3 GitHub accounts at once and having to use special `Host` values in my SSH config like:
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

And then I had to alter the remote url for each repo to have the correct permissions to the remote repo: \
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

<br>

GoGit will only change the `Host` values in your SSH config. \
For example: `github.com-<name>` to `github.com` and back. \
All other settings will be left unchanged.

<hr>

### `gogit set <name>`

Set the current active account

<hr>

### `gogit init`

Create a blank GoGit config file

<hr>

### `gogit active`

See the current active account

<hr>
<br>

## SSH config

GoGit currently requires you to have your SSH config file setup up in a specific way. The needs to be done manually for now. \
See the [problem](#problem) section for an example structure of a working `~/.ssh/config` file.

## The config file `~/gogit.ini`


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

- The `set-git-user` alters your `~/.gitconfig` file
- Each profile has unique settings

## Features to add

- [ ] Option to add new block to ssh config on `gogit add`
- [ ] Add `gogit remove` command to remove a profile/account
- [x] Add support for other git remotes
- [x] Setting `git config --global` when activating account
