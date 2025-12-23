package output

import (
	"encoding/json"
	"fmt"

	"github.com/akmanon/k-ray/pkg/models"
)

func PrintJson(findings []models.Findings) error {
	data, err := json.MarshalIndent(findings, "", " ")
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}
