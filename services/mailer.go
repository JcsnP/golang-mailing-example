package services

import (
  "log"

  "gopkg.in/gomail.v2"

  "github.com/JcsnP/go-mail/config"
)

type Mailer struct {}

func (m *Mailer) Send(message *gomail.Message) {
  message.SetHeader("From", "janedoe@mail.com")

  if err := config.Mailer.DialAndSend(message); err != nil {
    log.Panicln("[Mailer] ", err)
  }
}
