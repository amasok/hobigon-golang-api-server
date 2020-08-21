
# Hobigon Server

![Go](https://img.shields.io/badge/Go-1.15.0-blue.svg)
![GitHub Actions](https://img.shields.io/github/workflow/status/yyh-gl/hobigon-golang-api-server/Workflow%20for%20Golang)

**Hobby + [KANEGON](https://m-78.jp/ultraman-archives/kanegons-cocoon/) = Hobigon**

Hobigon is server which make rich my live.
<br>
It's just a hobby.

I remade that was made in Ruby(Rails).
-> [Hobigon ver.Ruby repository](https://github.com/yyh-gl/hobigon-rails-api-server) (sorry, this is private repo) 

<br>

Hobigon has two interfaces, Web API and CLI.

## Web API
Hobigon don't use WAF and use only [gorilla/mux](https://github.com/gorilla/mux).
<br>
(Use [julienschmidt/httprouter](https://github.com/julienschmidt/httprouter) 
until [This commit](https://github.com/yyh-gl/hobigon-golang-api-server/tree/b0c0fb3e52df7714593386840e64a9bf7f32f1a4))

### Features

- Slack Notification
  - Access ranking of [My Blog](https://yyh-gl.github.io/tech-blog/)'s posts.
  - Today's task list
  - Today's birthday people
- [My Blog](https://yyh-gl.github.io/tech-blog/)
  - Create post data
  - Get post data
  - Like👍 post
- Birthday
  - Create birthday data

## CLI
Hobigon use [urfave/cli](https://github.com/urfave/cli).

### Features

- Slack Notification
  - Access ranking of [My Blog](https://yyh-gl.github.io/tech-blog/)'s posts.
  - Today's task list
  - Today's birthday people


# TODO
- Apply the tactical patterns of DDD
  - Improve anemic domain model
    - Slack
    - Task
