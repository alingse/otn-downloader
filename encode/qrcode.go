package encode

import (
	"encoding/base64"
	"encoding/json"
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

type MetaValue struct {
	Filename string `json:"filename"`
	Total int `json:"total"`
	FileSize int `json:"file_size"`
	ChunkSize int `json:"chunk_size"`
}

func printQRCode(v Value) {
	fmt.Printf("\033[0;0H")
	text := fmt.Sprintf("%s:%s:%s", v.Key, v.Index, v.Value)
	qrterminal.Generate(text, qrterminal.L, os.Stdout)
}

func loadValues(filepath string, cfg Config) ([]Value, []Value, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	datas := make([]Value, 0)
	buf := make([]byte, cfg.ChunkSize)
	s := 0
	i := 0
	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, nil, err
		}
		s += n
		value := base64.StdEncoding.EncodeToString(buf[:n])
		data := Value{
			Key:   KeyData,
			Index: strconv.FormatInt(int64(i), 10),
			Value: value,
		}
		datas = append(datas, data)
		i++
	}

	_, filename := path.Split(filepath)
	meta := MetaValue{
		Filename: filename,
		Total: i,
		FileSize: s,
		ChunkSize: cfg.ChunkSize,

	}
	metaData, _ := json.Marshal(meta)
	metas := []Value{
		{
			Key: KeyMeta,
			Index: "json",
			Value: string(metaData),
		},
	}
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
	for i, v := range datas {
		if len(cfg.Slices) > 0 && !cfg.Slices[i] {
			continue
		}
		printQRCode(v)
		time.Sleep(d)
	}
	return nil
}

type Config struct {
	Fps       int
	ChunkSize int
	Loop      int
	Slices 	  map[int]bool
}

func EncodToQRCode(filename string, cfg Config) {
	for i := 0; i < cfg.Loop; i++ {
		err := encodeToQRCode(filename, cfg)
		if err != nil {
			panic(err)
		}
	}
}
