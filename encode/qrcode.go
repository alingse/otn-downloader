package encode

import (
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"path"
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

func loadValues(filepath string, cfg Config) ([]Value, []Value, error) {
	_, filename := path.Split(filepath)
	metas := make([]Value, 0, 2)
	metas = append(metas, Value{
		Key:   KeyMeta,
		Index: "filename",
		Value: filename,
	})

	file, err := os.Open(filepath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	datas := make([]Value, 0)
	buf := make([]byte, cfg.ChunkSize)
	i := 0
	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, nil, err
		}
		value := base64.StdEncoding.EncodeToString(buf[:n])
		data := Value{
			Key:   KeyData,
			Index: strconv.FormatInt(int64(i), 10),
			Value: value,
		}
		datas = append(datas, data)
		i++
	}
	metas = append(metas, Value{
		Key:   KeyMeta,
		Index: "total",
		Value: strconv.FormatInt(int64(i), 10),
	})
	return metas, datas, nil
}

var metaSleep = 5 * time.Second

func encodeToQRCode(filename string, cfg Config) error {
	metas, datas, err := loadValues(filename, cfg)
	if err != nil {
		return err
	}
	for _, v := range metas {
		printQRCode(v)
		time.Sleep(metaSleep)
	}

	d := 1 * time.Second / time.Duration(cfg.Fps)
	for _, v := range datas {
		printQRCode(v)
		time.Sleep(d)
	}
	return nil
}

type Config struct {
	Fps       int
	ChunkSize int
	Loop      int
}

func EncodToQRCode(filename string, cfg Config) {
	for i := 0; i < cfg.Loop; i++ {
		err := encodeToQRCode(filename, cfg)
		if err != nil {
			panic(err)
		}
	}
}
