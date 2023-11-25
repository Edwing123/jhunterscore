# Repository for JHuntersCore and friends

This repository contains the code for the following services:

- JHuntersCore.
- JHuntersCMSView.

## Inicio el servidor

### Crear archivo de configuracion

Tienes que crear un archivo de configuracion
en formato JSON.

Dicho archivo tiene que contener un objeto JSON
con las siguientes propiedades:

- `server_port` - Numero de puerto en donde el servidor escuchara peticiones.
- `data_dir` - Ruta del directorio en donde la aplicacion guardara archivos/datos.

Guarda este archivo en algun lugar en el sistema de archivos, recuerda guardar
la ruta del archivo, ya que se utilizara en el siguiente paso.

Por otro lado, si estas leyendo este documento en el repo, podes ver
un ejemplo de archivo de configuracion [aqui](./config.example.json).

### Ejecuta el servidor

Abre un sesion de terminal y navega en donde tengas almacenado el ejecutable del servidor,
luego ejecutar el siguiente comando:

```sh
./servidor -config=<ruta config>
```

Recuerda reemplazar `<ruta config>` con la ruta en donde guardaste
el archivo de configuracion.

El resultado de esto seria el programa creando el "data dir", y comenzando a escuchar
por peticiones en el puerto proporcionado en el archivo de configuracion.
