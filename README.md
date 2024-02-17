# 😀 emoji - a simple TUI emoji picker

https://github.com/eriks47/emoji/assets/90338990/93ebb03e-0a4e-42d5-97a4-d71d2769810a

## ⭐ Features
- Almost 2000 emojis
- Simple and easy to use

## 📋 Usage
There are no settings, no flags -- just run the binary and pick an emoji:
```shell
$ emoji
```
After pressing enter on a selected emoji it will be copied to the system
clipboard so that you can paste it wherever you want.

The search matches all emojies that contain the query *as a substring* and
the search is case insensitive.

## ⚙️  Build
```shell
go build main.go emoji.go
sudo cp main /usr/bin/emoji
```



