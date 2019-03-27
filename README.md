# learn-go

## 개발환경 셋팅

### Go 설치 [gvm](<https://github.com/moovweb/gvm>)

#### Pre required if mac

```zsh
❯ xcode-select --install
❯ brew update
❯ brew install mercurial
```

#### Install gvm

```zsh
❯ zsh < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
```

#### install binary

```zsh
# Pre required
❯ gvm install go1.4 -B
❯ gvm use go1.4
❯ export GOROOT_BOOTSTRAP=$GOROOT

❯ gvm listall
gvm gos (available)

   ...
   go1.9.7
   ...
❯ gvm install go1.9.7
❯ gvm use go1.9.7 --default
```

