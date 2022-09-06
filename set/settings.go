package set

import (
	"encoding/json"
	"fmt"
	"os"
)

type setting struct {
	ServerHost string
	ServerPort string
	PgHost     string
	PgPort     string
	PgUser     string
	PgPass     string
	PgBase     string
}

var cfg setting

func init() {
	fmt.Println("Read setting.cfg")

	file, e := os.Open("set/setting.cfg")
	if e != nil {
		fmt.Println(e.Error())
		panic("Конфигурационный файл отсутствует")
	}

	defer file.Close()

	stat, e := file.Stat()
	if e != nil {
		fmt.Println(e.Error())
		panic(" NE удалось прочитать информацию  файле конфигурации")
	}

	readByte := make([]byte, stat.Size())

	_, e = file.Read(readByte)
	if e != nil {
		fmt.Println(e.Error())
		panic("NE удалось прочитать файл конфигурации")
	}

	json.Unmarshal(readByte, &cfg)
	if e != nil {
		fmt.Println(e.Error())
		panic("NE удалось прочитать файл конфигурации")
	}

	///
	e = json.Unmarshal(readByte, &cfg)
	if e != nil {
		fmt.Println(e.Error())
		panic("NE удалось расшифровать файл конфигурации")
	}

	fmt.Println(cfg.ServerHost + ": " + cfg.ServerPort)

}
