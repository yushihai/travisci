package hello
import (
	"testing"
)
func TestHelloGolang(t *testing.T) {
	if helloGolang() == "helloGolang" {
		t.Log("测试通过")
	} else {
		t.Error("测试失败")
	}
}