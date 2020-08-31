package file

//testing

import (
	"fmt"
	"runtime"
	"strings"
	"testing"

	"github.com/monopolly/console"
	"github.com/stretchr/testify/assert"
)

func TestDownload(t *testing.T) {
	function, _, _, _ := runtime.Caller(0)
	fn := runtime.FuncForPC(function).Name()
	var log = console.New()
	log.OK(fmt.Sprintf("%s\n", fn[strings.LastIndex(fn, ".Test")+5:]))
	a := assert.New(t)
	_ = a

	f, err := Download("https://upload.wikimedia.org/wikipedia/commons/e/e1/Gaoliang_Bridge.JPG")
	a.Nil(err)

	Save("download_small.jpg", f)

}

func BenchmarkNew(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		//db.Search(30, 0, []uint32{1, 2, 3})
	}
}

func BenchmarkGetFreeParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {

		}
	})
}
