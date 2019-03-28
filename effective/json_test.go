package effective

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJson(t *testing.T) {
	user := User{Email:"mo", Password:"", Tips:nil}
	str, _ := json.Marshal(user)

	fmt.Println(string(str))
}