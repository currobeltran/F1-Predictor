#Se utiliza la imagen de golang alpine en su última versión
FROM golang:1.15-alpine

#Establecemos el directorio de trabajo y copiamos el makefile en el mismo
WORKDIR /test
COPY makefile .

#Se instala make ya que nos hará falta para ejecutar los tests automáticamente
#Incluimos también gcc ya que será necesario a la hora de ejecutar nuestro código
#Instalamos gin-gonic para ejecutar los test del API
RUN apk update && apk upgrade \
&& apk add make \
&& apk add build-base \
&& apk add git \
&& go get -u github.com/gin-gonic/gin

#No podemos ejecutar el test sin ser superusuario, ya que necesita instalar dependencias
#en el proceso de prueba

#Comando que se ejecutará, en este caso para realizar los tests del proyecto
CMD make test 