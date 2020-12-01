
<h1 align="center">MicroShell</h1>
<p align="center"><img src="https://i.imgur.com/xe9L9DY.png"/></p>

---
**Microshell** is a lightweight and simple Linux shell written in `golang`.
I wrote it as an exercise to learn `go` and also know a little bit more about the Linux tools.
Special thanks to [Simon Jürgensmeyer](https://github.com/sj14) for writting the article [Writing a simple shell in Go](https://sj14.gitlab.io/post/2018/07-01-go-unix-shell/) and inspiring me to try it out by myself c:

---
## Goals
- :heavy_check_mark: Run commands on user $PATH
- :heavy_check_mark: `cd` builtin command
- :white_check_mark: `test` builtin command
- :white_check_mark: tab autocompletion
- :white_check_mark: wildcard expansion
---
### Install
```
$ go get https://github.com/franciscocid/microshell
```
---
<img src="https://golang.org/lib/godoc/images/footer-gopher.jpg"/>