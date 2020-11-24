# F1-Predictor
Proyecto para la asignatura de Infraestructura Virtual de 4º curso de Ingeniería Informática

## ¿Por qué F1 Predictor? Motivación del proyecto

En Fórmula 1, como en cualquier otra modalidad deportiva, las estadísticas juegan un papel fundamental. Todos los equipos están calculando probabilidades sobre que suceso puede o no ocurrir a cada momento, para anticiparse a los rivales y sacar el mejor resultado posible. Por ejemplo, en una situación de bandera amarilla o de Safety Car, si el equipo está preparado para ello con la estrategia más adecuada, puede hacerle conseguir un resultado que en condiciones normales no podría ser posible. También sería útil para dichos equipos deducir como enfocar un Gran Premio dependiendo de estadísticas relacionadas con el circuito, como pueden ser los adelantamientos, ya que no será igual correr en un circuito donde históricamente se adelanta muy poco que en otro donde se producen muchos eventos de este tipo.

Además, los más aficionados suelen buscar estadísticas antes de cada Gran Premio, para ver datos históricos de dicho circuito, records marcados, los ganadores de los grandes premios anteriores, entre otros muchos datos con el objetivo de aumentar sus conocimientos sobre la categoría.

Nuestra aplicación tratará información de manera básica, pero intentando cumplir los objetivos mencionados aquí arriba sentando una base que en un futuro podría mejorar a un servicio más especializado; con métodos para el cálculo de probabilidades de eventos más precisos.

## Serverless

En este apartado de nuestro proyecto vamos a desplegar dos funciones relacionadas con nuestro proyecto. Una de ellas nos permitirá seleccionar entre el archivo histórico de los resultados de las sesiones de clasificación del GP de Australia; mientras que la otra nos permitirá mediante el número de coche del piloto obtener su información. El código fuente de la primera función se encuentra en el siguiente [enlace](./api/clasificacion.go), mientras la segunda está [aquí](./functions/index.js).

Para el despliegue de las siguientes funciones hemos utilizado las plataformas Vercel y Firebase respectivamente, en las cuales crearemos una cuenta (con GitHub en el primer caso y con nuestra cuenta de Google en el segundo se puede hacer este proceso de manera automática) y crearemos el proyecto en el cual se alojará nuestra función.

Para desplegar una función desde nuestro repositorio usando alguno de los mencionados servicios, usaremos el CLI que tenga cada plataforma (previa instalación a través del gestor de paquetes npm), que dispondrá de un comando con el que podremos inicializar dichos servicios en nuestro proyecto. Durante el proceso, el CLI nos guiará para configurarlo correctamente y crear todos los archivos necesarios dentro de nuestro repositorio. En el caso de Vercel, se puede enlazar con nuestro repositorio de GitHub para que, en cada actualización del repositorio, se realice un nuevo despliegue a producción; mientras que con Firebase tendremos que realizar este proceso de manera manual (que con ayuda del CLI se realiza con un simple comando, y para automatizarlo hemos hecho un hook que lo ejecutará cada vez que se haga push).

En las dos funciones hemos utilizado la misma forma de tratar una petición HTTP: a través de la URI, se accede al recurso (nuestra función) y se establecen los parámetros de búsqueda que queramos para filtrar la salida del recurso (este proceso se facilitará ya que en cada función se dispone de un formulario para introducir esta información). Esta información de entrada se procesa (ya que los parámetros al obtenerlo aparecerán como un único string) y se llama a las distintas casuísticas de nuestra función.

Por ejemplo, en la función realizada para Vercel (escrita en Go), hemos establecido inicialmente dos casos, que son cuando existe una entrada de parámetros y cuando no. En el primer caso, tratamos la cadena de entrada para decidir los resultados de qué temporada debemos mostrar llamando a los métodos dispuestos para ello, comprobando además si estos son correctos o si hay algún error en la entrada de los parámetros. En el segundo caso, mostramos directamente la página donde estará el formulario que servirá para filtrar la información.

A la hora de dar una respuesta a la petición, siempre debemos establecer que tipo de respuesta estamos dando. En casi todas las situaciones la respuesta será de tipo text/html, ya que mostraremos una página; pero en la función de Vercel establecemos un parámetro para poder obtener la información como application/json.

La historia de usuario que se ha seguido para este hito se puede ver desde este [enlace](https://github.com/currobeltran/F1-Predictor/issues/65)

- [Ir a función clasificacion (vercel)](https://f1-predictor.vercel.app/api/clasificacion)
- [Ir a función pilotos (firebase)](https://us-central1-f1-predictor.cloudfunctions.net/piloto)

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