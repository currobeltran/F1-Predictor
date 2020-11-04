# Documentación sobre Travis

Para añadir integración continua en nuestro proyecto nos hemos registrado en Travis, un sistema que nos permitirá ejecutar los test de nuestro código de manera automática cada vez que actualicemos el repositorio.

## ¿Qué pasos hemos seguido?

Primero, debemos crear una cuenta en la página de [Travis](https://travis-ci.com/), donde nos podemos registrar con nuestra cuenta de GitHub para más comodidad.

Cuando ya esté creada, tenemos que activar la GitHub App que nos permitirá ejecutar Travis cada vez que se realice una actualización sobre el repositorio. Para ello simplemente tendremos que ir a la configuración de nuestro perfil de Travis y en la sección de repositorios tendremos un aviso recordándonos que no tenemos la GitHub App instalada, con un botón para iniciar el proceso. Esto nos redireccionará a GitHub, donde nos permitirá confirmar su instalación. Al volver a la configuración de Travis veremos como aparecen todos nuestros repositorios.

El siguiente paso será crear un archivo .yml (que será el ejecutado por Travis) en la raíz de nuestro proyecto, llamado .travis.yml. En nuestro caso es muy simple, ya que como vemos tenemos 2 líneas; una que define el lenguaje como `minimal` y otra que llama a la regla travis de nuestro makefile, que construye el contenedor y lo ejecuta. Se escoge la opción minimal ya que, al ejecutarse un contenedor no será necesario instalar una versión específica para un lenguaje o la genérica; además de ser una versión que da mejores datos en cuanto a velocidad de ejecución se refiere (unos 10s de diferencia entre los tests realizados con la opción de lenguaje go y minimal).

## Otros sistemas de integración continua

A parte de Travis, también existen otros sistemas que nos permiten realizar las mismas tareas, como es el caso de Shippable, Jenkins, Circle CI, Bitrise...

En primer lugar, buscamos un servicio gratuito, por tanto algunas alternativas como TeamCity, Bitrise o Bamboo quedarían descartadas; ya que solo dejan un tiempo de prueba o tienen limitaciones a la hora de realizar la integración continua. De entre algunos de los servicios gratuitos, veamos sus características:

- Shippable: A diferencia de Travis este servicio se ejecuta sobre docker y no sobre una máquina virtual, lo que es más ventajoso ya que el estado del contenedor es persistente y no hará falta reinstalar elementos que ya estaban presentes o que no se han modificado. Además, es muy simple de configurar ya que con el script usado para Travis se puede ejecutar también la integración continua de Shippable. Como muchos otros servicios de este estilo, se pueden enlazar de manera sencilla a un repositorio de GitHub para que, cuando este se actualice, se ejecute de manera automática el proceso de CI. El uso de contenedores le permite tener información de otras ejecuciones, agilizando más el proceso ya que no será necesario instalar todos los elementos y dependencias en cada ejecución.

- Jenkins: Esta alternativa, además de ser gratuita, es open source. Su funcionamiento se puede extender mediante plugins (por lo que en nuestro servicio de integración continua siempre tendremos lo justamente necesario, haciéndolo más eficiente); pero su configuración es más dificil ya que tenemos que iniciar nosotros mismos el servicio de manera local por ejemplo a través de un contenedor, a diferencia de otras alternativas que se ejecutan automáticamente y se configuran desde una interfaz web.

- Circle CI: Muy similar a Travis, nos permite enlazar nuestra cuenta de GitHub con la de Circle CI; con las ventajas que ello conlleva como es poder ejecutar de manera automática el servicio cada vez que se actualice el repositorio. Al igual que Travis, funciona en base a un archivo de configuración .yaml. Como punto en contra podemos comentar que funciona sobre una máquina virtual y no sobre un contenedor a diferencia de Shippable, lo que reduciría su rendimiento respecto a esta última.

Teniendo en cuenta lo dicho anteriormente, vamos a utilizar Shippable, ya que en comparación con Travis es más rápido (diferencia más que notable cuando se realizan test sobre múltiples versiones), donde la opción escogida tarda 2 minutos en ejecutar 7 versiones distintas del lenguaje en el mismo test, mientras que Travis tarda casi 5 minutos en hacer la misma operación (dicha diferencia existe también en test de una única versión, donde Shippable tarda 18-23 segundos y Travis unos 30-40).

Esta elección viene también dada por la simplicidad en la solución en aspectos como pueden ser la configuración y ejecución del proceso de integración continua, a diferencia de Jenkins; ya que como dijimos anteriormente Shippable se ejecuta automáticamente mientras que Jenkins realiza su funcionalidad de manera local mediante, por ejemplo, un contenedor (además que la configuración de dicho contenedor será más tediosa que el script que, por ejemplo, hemos tenido que realizar nosotros). Este problema se enfatizaría aún más si estuviesemos hablando de un gran entorno colaborativo, ya que nosotros mismos tendríamos que mantener el servicio de integración continua activo 24/7; mientras que con Shippable no nos tenemos que preocupar de posibles caidas del servicio.

Para ello tendremos que seguir pasos muy similares a los seguidos para Travis, es decir, registrarnos en la página del [servicio](https://shippable.com) con nuestra cuenta de GitHub. Posteriormente, en la configuración seleccionaremos qué repositorios queremos que se comprueben cada vez que se actualicen.

Para acabar con la configuración de Shippable, se ha creado un script llamado shippable.yml. Esto se debe a que el script de travis que tenemos no funcionaría correctamente en este servicio, ya que la opción de lenguaje `minimal` daría un error. Este fichero de configuración lo que realizará será definir el lenguaje del programa que vamos a testear (Go), las versiones del lenguaje a comprobar, y ejecutar el comando correspondiente a las pruebas del proyecto.
