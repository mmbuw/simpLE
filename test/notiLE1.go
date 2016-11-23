package main

import (
  "github.com/0xAX/notificator"
)

var notify *notificator.Notificator

func main() {

  notify = notificator.New(notificator.Options{
    DefaultIcon: "icon/default.png",
    AppName:     "My test App",
  })

  notify.Push("title", "text", "/home/user/icon.png", notificator.UR_CRITICAL)
}