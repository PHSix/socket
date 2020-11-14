package requests

import (
	"fmt"
	"testing"
)


func TestGetUrl(t *testing.T){
  q := Get("今天成都天气")
  fmt.Println(q)
}
