package services

import "strings"

func ConfigureMessage(message string) []string {
	// Dividir el mensaje en líneas
	lines := strings.Split(message, "\n")

	// Crear un slice para almacenar las opciones
	options := make([]string, 0)

	// Recorrer cada línea y extraer las opciones numeradas
	for _, line := range lines {
		// Obtener el número de opción
		if len(line) > 2 && line[1] == '.' {
			option := strings.TrimSpace(line[2:])
			options = append(options, option)
		}
	}

	// Retornar el slice de opciones
	return options
}
