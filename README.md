# Golang broker jardín
Proyecto para construcción de Broker MQTT Para automatización
de jardín

## Installation

### Golang

```bash
  go mod download
  go build main.go
  main
```

### Docker
En el proyecto hay dos archivos Dockerfile. Dockerfile.dev para ambiente de desarrollo
con hot-reaload. Y un dockerfile para producción.

## Environment Variables

Para ejecutar este proyecto, deberá agregar las siguientes variables de entorno a su archivo .env

| Variable          | Descripción                          | Requerido     |
| :---------------- | :----------------------------------  | :------------ |
| `PORT`            | Puerto TCP Para el MQTT Server       | **Requerido** |
| `* USER_NAME`     | Nombres de usuario para los clientes | **Opcional**  |
| `* PASSWORD_NAME` | Contraseñas para los usarios         | **Opcioanal** |

Para las variables "*", se pueden crear tantas como usuarios sean necesarios.
Estos deben ser añadidos en la clase `settings` al ser instancias y ser configurado
sus permisos en `main`.