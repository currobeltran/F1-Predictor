#Se utiliza la imagen de golang alpine en su última versión
FROM golang:1.15-alpine

#Establecemos el directorio de trabajo y copiamos el makefile en el mismo
WORKDIR /go
COPY makefile ./

#Se instala make ya que nos hará falta para ejecutar los tests automáticamente 
#Se crea el directorio del proyecto
RUN mkdir ./src/f1predictor/ \
&& apk update && apk upgrade \
&& apk add make
COPY ./src/f1predictor/* ./src/f1predictor/

#Establecemos un usuario para realizar la ejecucion sin privilegios
RUN adduser --disabled-password currobeltran
USER currobeltran

#Comando que se ejecutará, en este caso para realizar los tests del proyecto
CMD make test 