package entity

import (
	"fmt"
	"time"
)

type User struct {
	ID       uint      `json:"id"`
	Name     string    `json:"name"`
	Age      uint      `json:"age"`
	Birthday time.Time `json:"birthday"`
}

func (u User) String() string {
	return fmt.Sprintf("\nid: %d\nname: %s \nage: %d \nbirthday: %s", u.ID, u.Name, u.Age, u.Birthday)
}
