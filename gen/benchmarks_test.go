package gen

import "testing"

func BenchmarkToCamelName(b *testing.B) {
	for n := 0; n < b.N; n++ {
		toCamelName("_hello_world_holl")
	}
}

func TestToCamelName(t *testing.T) {

	if toCamelName("_hello_world_holl_") != "HelloWorldHoll" {
		t.Fail()
	}

}
