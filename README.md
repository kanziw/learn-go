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

## Study 계획

### [Grammar](https://github.com/kanziw/learn-go/tree/master/00_grammar)

### Web application 구현

- [ ] Basic Endpoint
  - [ ] /
    - [ ] 서버 이름 & 버전 반환
  - [ ] /json
    - [ ] { ok: 1 } 반환
  - [ ] /static
    - [ ] static 파일 반환
    - [ ] html & css & jpg
  - [ ] /method/$METHOD 에 get/post/put/delete/patch 만들기
    - [ ] get 에선 querystring & API params  파싱 추가
    - [ ] post 에선 JSON 형태의 body 파싱 추가
- [ ] Basic Test
  - [ ] 각 Endpoint 들에 대한 반환값이 제대로 오는지 테스트케이스 작성
- [ ] Board
  - [ ] 테스트 환경 구축
    - [ ] DB: go
    - [ ] table : users
      - [ ] id: AUTO_INCREASE
      - [ ] name: VARCHAR(20)
      - [ ] created_at: DATETIME(3)
  - [ ] POST /user
    - [ ] name 받아 유저 데이터 생성
      - [ ] { data: User }
  - [ ] GET /users
    - [ ] 전체 유저 리스트 반환
      - [ ] { data: User[] }
  - [ ] GET /user/:id
    - [ ] id 유저 정보 JSON 으로 반환
      - [ ] { data: User }
  - [ ] POST /post
    - [ ] 게시글 작성
      - [ ] { data: Post }
  - [ ] GET /posts
    - [ ] 모든 게시글 조회
      - [ ] { data: Post[] }
  - [ ] GET /post/:id
    - [ ] 게시글 조회
      - [ ] { data: Post }

### Understanding module system in Go

