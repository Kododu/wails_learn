package connection

import (
	"changeme/internal/constant/exception"
	"changeme/internal/dao"
	"changeme/internal/define"
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
)

func DBConnection(identify string) error {
	file, err := ioutil.ReadFile(confFileName())
	if err != nil {
		return err
	}
	conf := &define.Config{}
	err = json.Unmarshal(file, conf)
	if err != nil {
		return err
	}
	for i := range conf.Connections {
		if identify == conf.Connections[i].Identify {
			return dao.DialDB(conf.Connections[i])
		}
	}

	return errors.Wrap(exception.ErrConnNotFoundInConf, identify)
}
