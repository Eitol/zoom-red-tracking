![Información de seguimiento](docs/Logo-Zoom-Registrado.png)

# Cliente de Zoom (Empresa de envios)

Este paquete de Go proporciona una interfaz de cliente simple para interactuar con el servicio de rastreo de la empresa de envíos Zoom. 

Permite obtener información de seguimiento de un paquete usando su número de guía.

La información obtenida es similar a la que se consigue en la página:



### Requisitos
Para usar este paquete, necesitará:

- Go versión 1.16 o posterior


### Instalación

```sh
go get  github.com/Eitol/zoom-red-tracking
```

### Uso básico

```go
package main

import (
    "fmt"
    "github.com/Eitol/zoom-red-tracking/zoom"
)

func main() {
	
	guiaDeEjemplo := 1553486107
	c := zoom.NewDefaultClient() 
	// reemplace guiaDeEjemplo con el número de guía de su paquete 
	info, err := c.GetTrackingInfo(guiaDeEjemplo) 
	if err != nil {
		fmt.Println("Error obteniendo información de seguimiento:", err)
		return
	}
	fmt.Printf("Información de seguimiento: %+v\n", info)
}
```

### Documentación del API

- https://documenter.getpostman.com/view/6789630/S1Zz6V2v#2bb89975-10f8-40cc-88d0-530e499320c5

### Aclaratoria

Este paquete no es oficial de Zoom Red, es un proyecto personal que hice para practicar Go y que puede ser útil para otras personas.