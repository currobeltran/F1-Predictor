# F1-Predictor
Proyecto para la asignatura de Infraestructura Virtual de 4º curso de Ingeniería Informática

## ¿Por qué F1 Predictor? Motivación del proyecto

En Fórmula 1, como en cualquier otra modalidad deportiva, las estadísticas juegan un papel fundamental. Todos los equipos están calculando probabilidades sobre que suceso puede o no ocurrir a cada momento, para anticiparse a los rivales y sacar el mejor resultado posible. Por ejemplo, en una situación de bandera amarilla o de Safety Car, si el equipo está preparado para ello con la estrategia más adecuada, puede hacerle conseguir un resultado que en condiciones normales no podría ser posible. También sería útil para dichos equipos deducir como enfocar un Gran Premio dependiendo de estadísticas relacionadas con el circuito, como pueden ser los adelantamientos, ya que no será igual correr en un circuito donde históricamente se adelanta muy poco que en otro donde se producen muchos eventos de este tipo.

Además, los más aficionados suelen buscar estadísticas antes de cada Gran Premio, para ver datos históricos de dicho circuito, records marcados, los ganadores de los grandes premios anteriores, entre otros muchos datos con el objetivo de aumentar sus conocimientos sobre la categoría.

Nuestra aplicación tratará información de manera básica, pero intentando cumplir los objetivos mencionados aquí arriba sentando una base que en un futuro podría mejorar a un servicio más especializado; con métodos para el cálculo de probabilidades de eventos más precisos.

## Platform as a Service (PaaS)

Para esta última parte del proyecto, se ha planteado usar un Platform as a Service para realizar el despliegue del microservicio construido. Para ello, se ha utilizado Heroku, una plataforma freemium que nos permitirá poner en marcha nuestra aplicación de manera básica sin ningún coste. Se ha decidido escoger esta opción debido principalmente a que soporta aplicaciones escritas en Go que usan además el microframework gin-gonic; y nos permite inclusive realizar el despliegue de manera automática cada vez que actualicemos el repositorio remoto de GitHub. Con características similares también se encuentran Platform.sh o CloudFoundry; pero tienen algunos inconvenientes como que solo dan una prueba gratuita en la cual debemos introducir una tarjeta de crédito para disponer de la misma (Platform.sh) o que no se puede desplegar automáticamente desde GitHub (CloudFoundry, no se encuentra esta característica en la documentación oficial). Otras opciones, como Openshift, no permiten desplegar utilizando este lenguaje de programación.

En primer lugar, para configurar la aplicación en Heroku, tendremos que crear una nueva aplicación desde el CLI con la opción --region eu y el nombre que queremos darle a la app. Posteriormente desde el panel de control disponible en la web podremos conectar nuestra cuenta de Heroku con la que tenemos en GitHub para, de esta manera, conectar la aplicación recién creada con el repositorio que dispone de los fuentes del microservicio. De esta manera, podremos realizar el despliegue automático cada vez que realicemos un push de este repositorio (y se hayan pasado los test de integración continua).

El siguiente paso que tendremos que hacer para poder desplegar nuestra aplicación es crear el archivo Procfile.

## Código fuente y test del código

Puede ver el código de la aplicación y los test de comprobación [pulsando aquí](https://github.com/currobeltran/F1-Predictor/tree/master/src/f1predictor)

Para ejecutar los test tendremos disponible un contenedor, cuyo repositorio se encuentra en este [enlace](https://hub.docker.com/r/currobeltran/f1-predictor). Para iniciar dicho contenedor, se puede ejecutar el siguiente comando desde el directorio raíz del proyecto:

```shell
docker run -t -v `pwd`:/test currobeltran/f1-predictor
```

## ¿Qué pasos se han seguido para realizar el proyecto?

Durante este proyecto se han ido cumpliendo distintos objetivos para llegar a la finalización del mismo, que se encontrarán en [este documento](./docs/pasosrealizados.md) en constante actualización.

## Documentación anterior

En cuanto a otra información de interés, tenemos los siguientes puntos:

- [¿Por qué he usado estas herramientas?](./docs/herramientas.md)
- [¿Cómo he configurado git?](./docs/configuracion.md)
- [¿Qué historias de usuario se están siguiendo?](./docs/hu.md)
- [Información sobre contenedor de pruebas](./docs/docker.md)

## Autor

[Francisco Beltrán](https://github.com/currobeltran)