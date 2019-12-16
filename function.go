package demo

import (
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/tekkamanendless/gcfhook"
)

// CloudFunction is an HTTP Cloud Function with a request parameter.
func CloudFunction(w http.ResponseWriter, r *http.Request) {
	log := logrus.New()

	if value := os.Getenv("GCP_PROJECT"); value == "" {
		log.Infof("GCP_PROJECT is not set; falling back to normal logging.")
	} else {
		hook, err := gcfhook.NewForRequest(r)
		if err != nil {
			log.Errorf("Could not set up gcfhook: %v", err)
		}

		if hook != nil {
			// Flush the logging entries when we're done.
			defer hook.Flush()

			// Add the hook.
			log.AddHook(hook)
			// Nullify the console output; we don't want to duplicate it.
			log.SetFormatter(&gcfhook.NullFormatter{})
		}
	}

	log.Infof("This is an info message.")
	log.Warnf("This is a warning message.")
	log.Errorf("This is an error message.")

	w.Write([]byte("Okay"))
}
