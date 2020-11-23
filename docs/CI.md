# ¿Qué otros sistemas de integración continua existen? ¿Por qué Shippable como segundo sistema de CI?

A parte de Travis, también existen otros sistemas que nos permiten realizar las mismas tareas, como es el caso de Shippable, Jenkins, Circle CI, Bitrise...

En primer lugar, buscamos un servicio gratuito, por tanto algunas alternativas como TeamCity, Bitrise o Bamboo quedarían descartadas; ya que solo dejan un tiempo de prueba o tienen limitaciones a la hora de realizar la integración continua. De entre algunos de los servicios gratuitos, veamos sus características:

- Shippable: A diferencia de Travis este servicio se ejecuta sobre docker y no sobre una máquina virtual, lo que es más ventajoso ya que el estado del contenedor es persistente y no hará falta reinstalar elementos que ya estaban presentes o que no se han modificado. Además, es muy simple de configurar ya que con el script usado para Travis se puede ejecutar también la integración continua de Shippable. Como muchos otros servicios de este estilo, se pueden enlazar de manera sencilla a un repositorio de GitHub para que, cuando este se actualice, se ejecute de manera automática el proceso de CI. El uso de contenedores le permite tener información de otras ejecuciones, agilizando más el proceso ya que no será necesario instalar todos los elementos y dependencias en cada ejecución.

- Jenkins: Esta alternativa, además de ser gratuita, es open source. Su funcionamiento se puede extender mediante plugins (por lo que en nuestro servicio de integración continua siempre tendremos lo justamente necesario, haciéndolo más eficiente); pero su configuración es más dificil ya que tenemos que iniciar nosotros mismos el servicio de manera local por ejemplo a través de un contenedor, a diferencia de otras alternativas que se ejecutan automáticamente y se configuran desde una interfaz web.

- Circle CI: Muy similar a Travis, nos permite enlazar nuestra cuenta de GitHub con la de Circle CI; con las ventajas que ello conlleva como es poder ejecutar de manera automática el servicio cada vez que se actualice el repositorio. Al igual que Travis, funciona en base a un archivo de configuración .yaml. Como punto en contra podemos comentar que funciona sobre una máquina virtual y no sobre un contenedor a diferencia de Shippable, lo que reduciría su rendimiento respecto a esta última.

Teniendo en cuenta lo dicho anteriormente, se ha utilizado Shippable, ya que en comparación con Travis es más rápido (diferencia más que notable cuando se realizan test sobre múltiples versiones), donde la opción escogida tarda 2 minutos en ejecutar 7 versiones distintas del lenguaje en el mismo test, mientras que Travis tarda casi 5 minutos en hacer la misma operación (dicha diferencia existe también en test de una única versión, donde Shippable tarda 18-23 segundos y Travis unos 30-40).

Esta elección viene también dada por la simplicidad en la solución en aspectos como pueden ser la configuración y ejecución del proceso de integración continua, a diferencia de Jenkins; ya que como dijimos anteriormente Shippable se ejecuta automáticamente mientras que Jenkins realiza su funcionalidad de manera local mediante, por ejemplo, un contenedor (además que la configuración de dicho contenedor será más tediosa que el script que, por ejemplo, hemos tenido que realizar nosotros). Este problema se enfatizaría aún más si estuviesemos hablando de un gran entorno colaborativo, ya que nosotros mismos tendríamos que mantener el servicio de integración continua activo 24/7; mientras que con Shippable no nos tenemos que preocupar de posibles caidas del servicio.

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