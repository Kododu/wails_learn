package connection

import (
	"changeme/internal/define"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"os"
)

// List 返回链接列表
func List() ([]*define.Connection, error) {
	file, err := ioutil.ReadFile(confFileName())
	if err != nil {
		return nil, err
	}
	conf := &define.Config{}
	err = json.Unmarshal(file, conf)
	if err != nil {
		return nil, err
	}
	return conf.Connections, nil
}

// Create a new connection
func Create(conn *define.Connection) error {
	log.Printf("%v", conn)
	if conn.Identify == "" {
		conn.Identify = uuid.New().String()
	}
	if conn.Address == "" {
		return errors.New("address must be not empty")
	}
	if conn.Name == "" {
		conn.Name = conn.Address
	}
	if conn.Port == "" {
		conn.Port = "6379"
	}
	conf := new(define.Config)
	pwd, _ := os.Getwd()
	fileName := confFileName()
	log.Println(fileName)
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		err := os.MkdirAll(pwd, 0666)
		if err != nil {
			return err
		}
		conf.Connections = []*define.Connection{conn}
	} else {
		err = json.Unmarshal(file, conf)
		if err != nil {
			return err
		}
	}
	update := false
	for i, c := range conf.Connections {
		if c.Name == conn.Name {
			conf.Connections[i] = conn
			update = true
		}
	}
	if !update {
		conf.Connections = append(conf.Connections, conn)
	}
	marshal, err := json.Marshal(conf)
	log.Println(string(marshal))
	if err != nil {
		return err
	}
	return ioutil.WriteFile(fileName, marshal, 666)
}

func Delete(identity string) error {
	file, err := ioutil.ReadFile(confFileName())
	if err != nil {
		return err
	}
	conf := &define.Config{}
	err = json.Unmarshal(file, conf)
	if err != nil {
		return err
	}
	del := false
	for i := range conf.Connections {
		if identity == conf.Connections[i].Identify {
			conf.Connections = append(conf.Connections[:i], conf.Connections[i+1:]...)
			del = true
			break
		}
	}
	if !del {
		return errors.New("not found connection in conf : " + identity)
	}
	marshal, err := json.Marshal(conf)
	log.Println(string(marshal))
	if err != nil {
		return err
	}
	return ioutil.WriteFile(confFileName(), marshal, 666)
}

func confFileName() string {
	pwd, _ := os.Getwd()
	return pwd + string(os.PathSeparator) + define.ConfigName
}
