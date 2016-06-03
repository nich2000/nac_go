package nac_log

import (
  "log"
  "os"
)

var logger = log.New(os.Stdout, "", log.Lmicroseconds)

func Add_to_log(message string) {
  logger.Print(message)
}
