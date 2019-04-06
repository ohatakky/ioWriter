package f

import (
	"encoding/json"
	"os"
)

func JsonEncoderWrite() {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "	")
	encoder.Encode(map[string]string{
		"sample": "encoding/json",
	})
}
