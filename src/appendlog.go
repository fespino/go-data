package main

import (
  "os"
  "bufio"
  "strings"
  "fmt"
)

const tomb = "~"

func Set(db string, key string, value string) error {

  f, err := os.OpenFile(db, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModeAppend | 0644)
  if err != nil {
    return err
  }

  defer f.Close()

  buffer := bufio.NewWriter(f)

  buffer.WriteString(key)
  buffer.WriteString(" ")
  buffer.WriteString(value)
  buffer.WriteString("\n")

  err = buffer.Flush()
  if err != nil {
    return err
  }
  return nil
}

func notFound(key string) (string, error){
  return "", fmt.Errorf("key not found: %s", key)
}

func Get(db string, key string) (string, error) {
  f, err := os.Open(db)
  if err != nil {
    return "", err
  }
  defer f.Close()

  scanner := bufio.NewScanner(f)
  res, err := notFound(key)
  for scanner.Scan() {
    record := scanner.Text()
    s := strings.Split(record, " ")
    if s[0] == key {
      res, err = s[1], nil
    }
  }
  if err := scanner.Err(); err != nil {
	  return "", err
  }
  if res == tomb {
    return notFound(key)
  }
  return res, err
}

func Del(db string, key string) error {
  return Set(db, key, tomb)
}

func main() {
  switch args := os.Args[1:]; args[0] {
    case "set":
      Set(args[1], args[2], args[3])
    case "get":
      s, _ := Get(args[1], args[2])
      fmt.Println(s)
    case "del":
      Del(args[1], args[2])
    default:
      panic(fmt.Sprintf("unknown command: %s", args[0]))
  }
}
