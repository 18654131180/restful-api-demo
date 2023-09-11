package version_test

import (
	"fmt"
	"testing"

	"gitee.com/go-course/restful-api-demo-g7/version"
)

func TestVersion(t *testing.T) {
	fmt.Println(version.FullVersion())
}
