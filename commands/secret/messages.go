package secret

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func respondError(err error) error {
	err = fmt.Errorf("error occured: %s", err.Error())
	logrus.Error(err.Error())
	return err
}

func table(kvs map[string]string) (out string) {
	for k, v := range kvs {
		out += fmt.Sprintf("%s=%s\n", k, v)
	}

	return
}
