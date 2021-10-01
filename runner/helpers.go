package runner

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func errorf(r *Runner, format string, n ...interface{}) error {
	r.Status = ERROR

	return fmt.Errorf(format, n...)
}

func uploadLogs() {
	logrus.Error("not implemented")
	// file.io
	/*
		curl -X POST "https://file.io/"
		-H "accept: application/json"
		-H "Content-Type: multipart/form-data"
		-F "expires=1d"
		-F "maxDownloads=1"
		-F "autoDelete=true"
		-F "file=@Eula.txt;type=text/plain"
	*/
}
