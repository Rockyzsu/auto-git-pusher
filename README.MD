# What's it ?

----
## Run git add/commit/push daily automativily
* you can configure crontab that it can run daily to sync your code daily.

## support platform: 
linux, macos, windows,arm based 【since it is golang progam that it should support all platform, you can build it on different platfrom】

## How to use it ?
1. set git repo path to config.json and commit msg

```json
[
  {
    "path": "/root/github/git_test",
    "msg": "fixed with bug"
  },  {
    "path": "/root/github/git_test_tree",
    "msg": "updated everyday"
  }

]
```

2. run command with Golang

```bash
go run main.go
```

## build it to binary file
```bash
go build main.go
```