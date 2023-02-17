package utils

import "testing"

type Hello struct {
	Name string `validate:"required"`
	Age  int    `validate:"gte=0,lte=10,required"`
}

func TestValidate(t *testing.T) {

	h := &Hello{
		Name: "Tom",
	}

	err := GetFirestErr(h)
	t.Logf("validate err:%+v", err)
}
