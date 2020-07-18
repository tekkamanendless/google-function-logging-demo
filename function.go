package demo

import (
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/tekkamanendless/gcfstructuredlogformatter"
)

// CloudFunction is an HTTP Cloud Function with a request parameter.
func CloudFunction(w http.ResponseWriter, r *http.Request) {
	log := logrus.New()

	if value := os.Getenv("FUNCTION_TARGET"); value == "" {
		log.Infof("FUNCTION_TARGET is not set; falling back to normal logging.")
	} else {
		formatter := gcfstructuredlogformatter.New()

		log.SetFormatter(formatter)
	}

	log.Infof("This is an info message.")
	log.Warnf("This is a warning message.")
	log.Errorf("This is an error message.")

	w.Write([]byte("Okay"))
}

