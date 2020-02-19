package romantoint

import (
  "testing"
  "io"
  "os"
  "bufio"
  "encoding/json"
)

type Test struct {
  Input string `json:"input"`
  Output int `json:"output"`
}

func TestRomanToInt(testUtils *testing.T) {
  f, err := os.Open("tests.json")

  if err != nil {
    testUtils.Error(err)
    return
  }

  defer f.Close()

  reader := bufio.NewReader(f)
  decoder := json.NewDecoder(reader)
  var tests map[string]Test

  for {
    err = decoder.Decode(&tests)
    if err == nil {

    } else if err == io.EOF {
      break
    } else {
      testUtils.Error(err)
      break
    }
  }

}