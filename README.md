# go_basics

## Comandos básicos
- *go run*: Ejecuta un __script__ de go.
- *go build* [fichero_main]: Genera la solucion en un unico ejecutable.
- *go mod init* [ruta a tu paquete]: 
```go
go mod github.com/gmateos/accesodb
```
> **NOTA:**: go mod genera un fichero 'go.mod' el cual contiene todas las dependencias de tercero de tu proyecto.

- *go get [dependencia]*: La dependencia es la ruta de la paquetería, por ej:
```
go get github.com/gorm
```
> Instala la depencia de tu proyecto y modifica el fichero _go.mod_. 



