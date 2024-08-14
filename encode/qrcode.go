package encode

import (
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/mdp/qrterminal/v3"
)

type Key string

const (
	KeyMeta Key = "m"
	KeyData Key = "d"
)

type Value struct {
	Key   Key
	Index string
	Value string
}

func printQRCode(v Value) {
	fmt.Printf("\033[0;0H")
	text := fmt.Sprintf("%s:%s:%s", v.Key, v.Index, v.Value)
	qrterminal.Generate(text, qrterminal.L, os.Stdout)
}

func loadValues(filename string) []Value {
	values := make([]Value, 0)
	meta := Value{
		Key:   KeyMeta,
		Index: "filename",
		Value: filename,
	}
	values = append(values, meta)

	N := 5
	values = append(values, Value{
		Key:   KeyMeta,
		Index: "total",
		Value: strconv.FormatInt(int64(N), 10),
	})

	for i := 0; i < N; i++ {
		text := "hello " + strconv.FormatInt(int64(i), 10) + "\n"
		value := base64.StdEncoding.EncodeToString([]byte(text))
		values = append(values, Value{
			Key:   KeyData,
			Index: strconv.FormatInt(int64(i), 10),
			Value: value,
		})
	}
	return values
}

var metaSleep = 10 * time.Second

func encodeToQRCode(filename string, fps int) {
	d := 1 * time.Second / time.Duration(fps)
	values := loadValues(filename)
	for _, v := range values {
		printQRCode(v)
		if v.Key == KeyMeta {
			time.Sleep(metaSleep)
		}
		time.Sleep(d)
	}
}

func EncodToQRCode(filename string, fps int, loop int) {
	for i := 0; i < loop; i++ {
		encodeToQRCode(filename, fps)
	}
}
