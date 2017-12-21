package utils

import (
  "encoding/json"
  "bytes"
)

func StructToJson (data interface{}) ([]byte, error) {
  buf := new(bytes.Buffer)

  if err := json.NewEncoder(buf).Encode(data); err != nil {
    return nil, err
  }

  return buf.Bytes(), nil
}
