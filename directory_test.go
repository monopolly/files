package file

//testing

import (
	"fmt"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/monopolly/console"
)

func TestDir(t *testing.T) {
	function, _, _, _ := runtime.Caller(0)
	fn := runtime.FuncForPC(function).Name()
	var log = console.New()
	log.OK(fmt.Sprintf("%s\n", fn[strings.LastIndex(fn, ".Test")+5:]))
	a := assert.New(t)
	_ = a

	a.Len(FileList("dir", "txt"), 3)
	a.Len(FileList("dir"), 4)
	a.Len(FileList("dir", "doc"), 1)
	a.Len(FileList("", "txt"), 3)
	log.Info(FileList("dir"))

}
