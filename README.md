# F1-Predictor
Proyecto para la asignatura de Infraestructura Virtual de 4º curso de Ingeniería Informática

## ¿Por qué F1 Predictor? Motivación del proyecto

En Fórmula 1, como en cualquier otra modalidad deportiva, las estadísticas juegan un papel fundamental. Todos los equipos están calculando probabilidades sobre que suceso puede o no ocurrir a cada momento, para anticiparse a los rivales y sacar el mejor resultado posible. Por ejemplo, en una situación de bandera amarilla o de Safety Car, si el equipo está preparado para ello con la estrategia más adecuada, puede hacerle conseguir un resultado que en condiciones normales no podría ser posible. También sería útil para dichos equipos deducir como enfocar un Gran Premio dependiendo de estadísticas relacionadas con el circuito, como pueden ser los adelantamientos, ya que no será igual correr en un circuito donde históricamente se adelanta muy poco que en otro donde se producen muchos eventos de este tipo.

Además, los más aficionados suelen buscar estadísticas antes de cada Gran Premio, para ver datos históricos de dicho circuito, records marcados, los ganadores de los grandes premios anteriores, entre otros muchos datos con el objetivo de aumentar sus conocimientos sobre la categoría.

Nuestra aplicación tratará información de manera básica, pero intentando cumplir los objetivos mencionados aquí arriba sentando una base que en un futuro podría mejorar a un servicio más especializado; con métodos para el cálculo de probabilidades de eventos más precisos.

## Microservicios

En esta parte del proyecto vamos a diseñar y testear un microservicio. En primer lugar, tenemos que elegir un microframework para realizar la implementación de dicho microservicio. Estudiando distintas alternativas dentro de las posibilidades que ofrece Go, se ha decidido usar el módulo [gin-gonic](https://github.com/gin-gonic/gin), que tendrá todos los elementos necesarios para la realización de esta fase del proyecto, con diseño de rutas (y la agrupación de las mismas) y posibilidad de crear un middleware para realizar funciones como el registro de eventos personalizados (aún así ya el mismo sistema al recibir una petición lo registra y lo muestra por consola de forma transparente al desarrollador). En cuanto a las ventajas podemos mencionar que existe la capacidad de convertir estructuras de datos cualesquiera a JSON de manera automática a través de un método, lo cual nos permitirá olvidarnos de ese problema en nuestra aplicación. También podríamos añadir a este apartado que técnicamente es muy eficiente debido al uso de [httprouter](https://github.com/julienschmidt/httprouter), un enrutador de peticiones HTTP que tiene características extra (como la capacidad de añadir variables en la definición de la ruta, obviamente presente en nuestro microframework) frente al multiplexador de rutas del módulo `net/http` y que según los benchmarks que tienen referenciados, son la segunda opción más rápida de las comparadas.

Además, es software libre y dispone de una [documentación](https://gin-gonic.com/docs/) muy extensa donde, con ejemplos, explica toda la funcionalidad de la herramienta, y sobre todo, la que necesitamos para esta fase del proyecto: La creación de rutas con su correspondiente testeo.

Unas opciones muy similares a la elegida son [Goji](https://github.com/goji/goji), [Pat](https://github.com/bmizerany/pat) y [pastis](https://github.com/guregodevo/pastis), ya que tiene unas capacidades similares, (pastis además tiene una estructura de datos propia para definir recursos). El problema de estas dos opciones es que no tienen ningún tipo de soporte ya que llevan años sin recibir ni un commit. Un ejemplo de la falta de capacidad de estas dos herramientas es que no se encontraba en ninguna parte de la documentación como realizar correctamente una petición con algún tipo de contenido en el cuerpo (problema visto en el microframework pastis).

Se dejan de lado opciones que no concuerdan con el modelo de desarrollo que estamos siguiendo como podría ser FlowBase o TarsGo (no están diseñados para los paradigmas de programación que se está siguiendo) y Nitro, que actualmente se considera inestable según su documentación debido a un remodelado de la aplicación.

Entrando más en materia de que es lo que se ha desarrollado en esta fase de nuestro proyecto, tenemos la siguiente historia de usuario: [Como usuario, quiero que existan distintas rutas en el microservicio con el objetivo de acceder a los distintos recursos del mismo](https://github.com/currobeltran/F1-Predictor/issues/71) en la que se pueden ver los problemas que se han ido resolviendo para conseguir el resultado final. Para llegar a el, se han creado dos nuevos paquetes, uno que contendrá tanto el código de la función que ejecutará el microservicio como el test del mismo, y otro que contendrá las utilidades relacionadas con la API que utilizará el paquete main para cumplir su funcionalidad; separando de esta manera todo lo relacionado con la API del código de la aplicación ejecutada (f1predictor).

Primeramente, se han diseñado las rutas, correspondiendo al cumplimiento de las historias de usuario realizadas anteriormente para describir la funcionalidad básica que queríamos que cumpliese la aplicación de cara al usuario final, disponibles en los siguientes enlaces: [HU1](https://github.com/currobeltran/F1-Predictor/issues/3), [HU2](https://github.com/currobeltran/F1-Predictor/issues/4), [HU3](https://github.com/currobeltran/F1-Predictor/issues/5).

En el archivo [main](./src/main/main.go) vemos cuales son las que se han creado, almacenadas en una función denominada `DiseñoRutas` (establecidas de esta manera para facilitar el testeo). En la misma se crearán los objetos que tendrán todo el código para procesar los distintos tipos de peticiones dependiendo del recurso al que se quiera acceder y las rutas, que estarán divididas por grupos donde parte del nombre del recurso es común.

En las funciones que ejecutan la peticion (pertenecientes a los distintos objetos recurso), realizamos una serie de pasos que consisten, normalmente, en obtener la información necesaria de los archivos JSON que mantendrán la misma durante la ejecución de la aplicación. Dichos datos se almacenarán en un tipo de dato nativo de Go en el que se permite almacenar directamente todo lo extraido del JSON (un mapa con clave de tipo string y valor de tipo la estructura de datos de f1predictor que vayamos a utilizar en la función). Al realizar todo el procesamiento de la información (o la creación de nuevos objetos), se actualizan los archivos afectados para que los cambios sean efectivos durante la ejecución. Por último, dentro del objeto Context que se le pasa como parámetro a cada función escribimos la respuesta, tanto si ha dado error con su código correspondiente e informando de lo sucedido (con un error 400 si la petición no se ha formulado correctamente o un error 404 si no se ha encontrado el recurso solicitado); como si la operación ha tenido éxito (devolviendo la información del recurso solicitado, creado o borrado en formato JSON).

Para realizar los tests de integración, hemos creado el siguiente [archivo](./src/main/main_test.go) donde probamos que las rutas anteriormente definidas y los métodos que ejecutan al recibir una petición otorgan un resultado esperado. Para ello, en todas las funciones de Test, llamaremos a la función DiseñoRutas que genera toda la estructura de rutas explicada; y crearemos una petición HTTP creando tanto el objeto donde se escribirá el resultado como el objeto que tendrá la información de la petición. Para hacer la petición efectiva, debemos llamar al método ServeHTTP del objeto que gestiona el funcionamiento del microservicio; el cual convertirá la petición realizada con los parámetros convencionales de un handlerHTTP en Go a una petición que el microframework pueda procesar. Por último se compara si la información escrita en el objeto respuesta creado es el esperado (tanto el código de la operación como la información contenida en el JSON).

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