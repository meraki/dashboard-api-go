package meraki

import (
	"fmt"
	"log"
	"regexp"
)

type CustomLogger struct{}

// Debugf implements the resty.Logger interface
func (l *CustomLogger) Debugf(format string, v ...interface{}) {
	// Convertir la salida de depuración a una cadena
	// Convertir el formato y los argumentos a una cadena de depuración
	message := fmt.Sprintf(format, v...)

	// Censurar el encabezado de autorización
	// re := regexp.MustCompile(`Authorization: .*`)
	// censoredMessage := re.ReplaceAllString(message, "Authorization: [censored]")
	re := regexp.MustCompile(`Authorization: Bearer [a-zA-Z0-9]+([a-zA-Z0-9]{5})`)
	censoredMessage := re.ReplaceAllString(message, "Authorization: Bearer ****$1")

	// Imprimir la salida de depuración censurada
	log.Print(censoredMessage)

}

// Errorf implements the resty.Logger interface
func (l *CustomLogger) Errorf(format string, v ...interface{}) {
	// Log errors
	log.Printf(format, v...)
}

func (l *CustomLogger) Warnf(format string, v ...interface{}) {
	// Log warnings
	log.Printf(format, v...)
}
