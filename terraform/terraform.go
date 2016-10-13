package terraform

import (
	"github.com/hashicorp/terraform/terraform"
	"os"
)

func ReadState(path string) (*terraform.State, error) {
	if f, err := os.Open(path); err != nil {
		return nil, err
	} else {
		defer f.Close()
		state, _ := terraform.ReadState(f)
		return state, nil
	}
}
