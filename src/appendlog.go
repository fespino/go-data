import (
  "os"
  "buffer"
  "buffio"
  "strings"
  "fmt"
 )
 
func Set(db string, key string, value string) error {
  var buffer bytes.Buffer
  
  f, err := os.OpenFile(db, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModeAppend)
  if err != nil {
    return err
  }
  
  defer f.Close()
  
  buffer.WriteString(key)
  buffer.WriteString(" ")
  buffer.WriteString(value)
  buffer.WriteString("\n")
  
  _, err := f.WriteString(buffer)
  if err != nil {
    return err
  }
  return nil
}

func Get(db string, key string) (string, error) {
  f, err := os.OpenFile(db, os.O_RDONLY)
  if err != nil {
    return "", err
  }
  defer f.Close()
  
  scanner := bufio.Scanner(f)
  res, err := "", fmt.Errorf("key not found: %s", key)
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
  return res, err
}
