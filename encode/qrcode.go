package encode

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/mdp/qrterminal/v3"
)

func EncodToQrCode() {
	N := 100
	buf := &bytes.Buffer{}
	for i := 0; i < N; i++ {
		text := "hello " + strconv.FormatInt(int64(i), 10)
		qrterminal.Generate(text, qrterminal.L, buf)
		fmt.Println(len(buf.Bytes()))
		_, _ = os.Stdout.Write(buf.Bytes())
		fmt.Printf("\033[0;0H")
		buf.Reset()
		fmt.Print()
		time.Sleep(1 * time.Second)
	}
}
