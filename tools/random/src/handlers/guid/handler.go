package handler_guid

import (
	"fmt"

	"github.com/google/uuid"
)

func Handler() {
	id := uuid.New()
	fmt.Println(id.String())
}
