# F1-Predictor
Proyecto para la asignatura de Infraestructura Virtual de 4º curso de Ingeniería Informática

## ¿Por qué F1 Predictor? Motivación del proyecto

En Fórmula 1, como en cualquier otra modalidad deportiva, las estadísticas juegan un papel fundamental. Todos los equipos están calculando probabilidades sobre que suceso puede o no ocurrir a cada momento, para anticiparse a los rivales y sacar el mejor resultado posible. Por ejemplo, en una situación de bandera amarilla o de Safety Car, si el equipo está preparado para ello con la estrategia más adecuada, puede hacerle conseguir un resultado que en condiciones normales no podría ser posible. También sería útil para dichos equipos deducir como enfocar un Gran Premio dependiendo de estadísticas relacionadas con el circuito, como pueden ser los adelantamientos, ya que no será igual correr en un circuito donde históricamente se adelanta muy poco que en otro donde se producen muchos eventos de este tipo.

Además, los más aficionados suelen buscar estadísticas antes de cada Gran Premio, para ver datos históricos de dicho circuito, records marcados, los ganadores de los grandes premios anteriores, entre otros muchos datos con el objetivo de aumentar sus conocimientos sobre la categoría.

Nuestra aplicación tratará información de manera básica, pero intentando cumplir los objetivos mencionados aquí arriba sentando una base que en un futuro podría mejorar a un servicio más especializado; con métodos para el cálculo de probabilidades de eventos más precisos.

## ¿Qué pasos se han seguido para realizar el proyecto?

Durante este proyecto se han ido cumpliendo distintos objetivos para llegar a la finalización del mismo, que se encontrarán en [este documento](./docs/pasosrealizados.md) en constante actualización.

## Código fuente e instalación

Puede ver el código de la aplicación [pulsando aquí](https://github.com/currobeltran/F1-Predictor/tree/master/src/f1predictor)

A la hora de realizar la instalación de la aplicación, primero debemos generar el binario que se ejecutará. Para ello, se tiene que introducir el siguiente comando:

`go build ruta-hasta-main.go`

Después de esto y para ejecutar el programa, tendríamos que escribir:

`go run ruta-hasta-binario`

## Documentación adicional

En cuanto a otra información de interés, tenemos los siguientes puntos:

- [¿Por qué he usado estas herramientas?](./docs/herramientas.md)
- [¿Cómo he configurado git?](./docs/configuracion.md)
- [¿Qué historias de usuario se están siguiendo?](./docs/hu.md)

## Autor

[Francisco Beltrán](https://github.com/currobeltran)