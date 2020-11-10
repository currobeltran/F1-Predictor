# F1-Predictor
Proyecto para la asignatura de Infraestructura Virtual de 4º curso de Ingeniería Informática

## ¿Por qué F1 Predictor? Motivación del proyecto

En Fórmula 1, como en cualquier otra modalidad deportiva, las estadísticas juegan un papel fundamental. Todos los equipos están calculando probabilidades sobre que suceso puede o no ocurrir a cada momento, para anticiparse a los rivales y sacar el mejor resultado posible. Por ejemplo, en una situación de bandera amarilla o de Safety Car, si el equipo está preparado para ello con la estrategia más adecuada, puede hacerle conseguir un resultado que en condiciones normales no podría ser posible. También sería útil para dichos equipos deducir como enfocar un Gran Premio dependiendo de estadísticas relacionadas con el circuito, como pueden ser los adelantamientos, ya que no será igual correr en un circuito donde históricamente se adelanta muy poco que en otro donde se producen muchos eventos de este tipo.

Además, los más aficionados suelen buscar estadísticas antes de cada Gran Premio, para ver datos históricos de dicho circuito, records marcados, los ganadores de los grandes premios anteriores, entre otros muchos datos con el objetivo de aumentar sus conocimientos sobre la categoría.

Nuestra aplicación tratará información de manera básica, pero intentando cumplir los objetivos mencionados aquí arriba sentando una base que en un futuro podría mejorar a un servicio más especializado; con métodos para el cálculo de probabilidades de eventos más precisos.

## Integración continua

Para añadir integración continua en nuestro proyecto nos hemos registrado en Travis, un sistema que nos permitirá ejecutar los test de nuestro código de manera automática cada vez que actualicemos el repositorio.

### ¿Cómo hemos configurado Travis?

Antes de nada, si es la primera vez que usamos este servicio; debemos crear una cuenta en la página de [Travis](https://travis-ci.com/), donde nos podemos registrar con nuestra cuenta de GitHub para más comodidad; y posteriormente activar la GitHub App de Travis, para tener la capacidad de ejecutar el proceso en nuestros repositorios.

El siguiente paso será crear un archivo .yml (que será el ejecutado por Travis) en la raíz de nuestro proyecto, llamado [.travis.yml](./.travis.yml). En nuestro caso es muy simple, ya que como vemos tenemos 2 líneas que ejecutarán todo el proceso necesario:

- `language: minimal`: Este parámetro se establece con este valor y no como `go`, ya que no vamos a necesitar ningún lenguaje de programación en concreto en la máquina virtual que ejecutará el proceso de prueba del código. Esto se debe a que dichas pruebas se ejecutarán en un contenedor que tendrá todo el entorno configurado para ello en el archivo [Dockerfile](./Dockerfile). Añadir que, además, esta versión tiene una ejecución más rápida que las asociadas a un lenguaje o la genérica, ya que es más liviana que las demás (con la ventaja de que una de las pocas herramientas que incluye la imagen es Docker).

- `script: make travis`: Como vemos, llama a la regla travis de nuestro [makefile](./makefile), que tendrá la correspondiente orden para ejecutar el contenedor de tests almacenado en el repositorio de Docker Hub.

### ¿Cómo hemos configurado Shippable?

Los primeros pasos para iniciar este servicio en nuestro repositorio son iguales a los seguidos en Travis, es decir, registrarnos en su [página web](https://shippable.com) y activar sobre que repositorios queremos que se ejecute Shippable.

En cuanto a la ejecución del proceso sobre este repositorio, se ha creado un archivo llamado [shippable.yml](./shippable.yml) (se podría usar .travis.yml, pero Shippable no reconoce `minimal` como una opción de lenguaje). Las órdenes dadas son las siguientes:

- `language: go`: Definimos el párametro del lenguaje como go, ya que en este caso no se ejecutará el contenedor y sí será necesario instalarlo.

- `go: - 1.4 ... - 1.15`: Definimos para que versiones queremos ejecutar las pruebas. Hemos elegido las versiones del lenguaje 1.4 y 1.15 porque son la primera y la última en las que funciona nuestro código (además la versión 1.15 es en la cual se ha desarrollado). Se establecen además versiones intermedias de 2 en 2 ya que, según la [documentación de Go](https://golang.org/doc/devel/release.html#policy), una versión major se mantiene durante la salida de dos versiones del lenguaje nuevas.

- `build: ci: make test`: Ejecutamos la regla test de nuestro makefile, la cual se encargará de ejecutar las pruebas sobre nuestro proyecto.

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