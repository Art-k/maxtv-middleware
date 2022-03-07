package log_processing

import (
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"io"
	. "maxtv_middleware/pkg/common"
	"os"
)

func InitLog(f *os.File) {

	Log = logrus.New()

	w := io.MultiWriter(os.Stdout, f)

	Log.SetOutput(w)
	Log.SetFormatter(&easy.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "[%lvl%]:\t%time%\t%msg%\n",
	})

	Log.SetLevel(logrus.InfoLevel)
	Log.Trace("Application Started")

}
