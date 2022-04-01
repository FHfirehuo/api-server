package logger

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

type Logger struct {
}

func (l *Logger) config() error {

	err := os.MkdirAll(viper.GetString("logger.file.path"), 0744)
	if err != nil {
		fmt.Println("make log file path failed, err:", err)
		return err
	}

	logFile, err := os.OpenFile(viper.GetString("logger.file.path")+viper.GetString("logger.file.name"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return err
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)

	return nil
}

func Info(format string, v ...any) {
	l := Logger{}
	l.config()
	log.SetPrefix("[INFO] ")
	log.Printf(format, v)
	fmt.Println(format, v)
}

func Debug(format string, v ...any) {
	l := Logger{}
	l.config()
	log.SetPrefix("[DEBUG] ")
	log.Printf(format, v)
	fmt.Println(format, v)
}

func Error(format string, v ...any) {
	l := Logger{}
	l.config()
	log.SetPrefix("[ERROR] ")
	log.Printf(format, v)
	fmt.Println(format, v)
}

func Exception(message string, err error) {
	l := Logger{}
	l.config()
	log.SetPrefix("[ERROR] ")
	log.Fatal(message, err)
	fmt.Println(message, err)
}
