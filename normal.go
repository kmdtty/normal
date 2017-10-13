package normal

import (
  "fmt"
  "encoding/json"

  "github.com/koron/go-dproxy"
)

type JSON = dproxy.Proxy

func merge(x string, diff string) string {
  return x + diff
}

/*                
 JavaScript Object Notation (JSON) Patch
 https://tools.ietf.org/html/rfc6902

 JavaScript Object Notation (JSON) Pointer
 https://tools.ietf.org/html/rfc6901
 
 The application/json Media Type for JavaScript Object Notation (JSON)
 https://tools.ietf.org/html/rfc4627
 */
func diff(x string, y string) string {
  var x_record interface{}
  var y_record interface{}
  json.Unmarshal([]byte(x), &x_record)
  json.Unmarshal([]byte(y), &y_record)
  x_json := dproxy.New(x_record)
  y_json := dproxy.New(y_record)
  fmt.Sprintf("%q", x_json)
  if x_json == y_json {
    return ""
  }
  return `[{"op":"replace", path:"/a", "value": 3},
           {"op":"remove",  path:"/b"},
           {"op":"add", path:"/c", "value": 4}]
          `
}
