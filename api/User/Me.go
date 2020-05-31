package User

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetMe(w http.ResponseWriter, r *http.Request) {
	response := ""
	decoder := json.NewDecoder(r.Body)

	decodeErr := decoder.Decode(&response)

	if decodeErr != nil {
		fmt.Println("decode error ", decodeErr)
		return
	}
	fmt.Println(response)

}
