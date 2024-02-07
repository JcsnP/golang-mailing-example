package main

import (
  "os"
  "log"

  "gopkg.in/gomail.v2"
  "github.com/joho/godotenv"

  "github.com/JcsnP/go-mail/config"
  "github.com/JcsnP/go-mail/services"
)

func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }
  config.ConnectMailer(
    os.Getenv("MAILER_HOST"),
    os.Getenv("MAILER_USERNAME"),
    os.Getenv("MAILER_PASSWORD"),
  )

  m := services.Mailer{}
  message := gomail.NewMessage()
  message.SetHeader("To", "johndoe@mail.com")
  message.SetHeader("Subject", "Hello John!!!")
  message.SetBody("text/html", "Test sending email from Golang")
  m.Send(message)
  log.Println("Mail sent.")
}
