# Documentación sobre Travis

Para añadir integración continua en nuestro proyecto nos hemos registrado en Travis, un sistema que nos permitirá ejecutar los test de nuestro código de manera automática cada vez que actualicemos el repositorio.

## ¿Qué pasos hemos seguido?

Primero, debemos crear una cuenta en la página de [Travis](https://travis-ci.com/), donde nos podemos registrar con nuestra cuenta de GitHub para más comodidad.

Cuando ya esté creada, tenemos que activar la GitHub App que nos permitirá ejecutar Travis cada vez que se realice una actualización sobre el repositorio. Para ello simplemente tendremos que ir a la configuración de nuestro perfil de Travis y en la sección de repositorios tendremos un aviso recordándonos que no tenemos la GitHub App instalada, con un botón para iniciar el proceso. Esto nos redireccionará a GitHub, donde nos permitirá confirmar su instalación. Al volver a la configuración de Travis veremos como aparecen todos nuestros repositorios.

El siguiente paso será crear un archivo .yml (que será el ejecutado por Travis) en la raíz de nuestro proyecto, llamado .travis.yml. En nuestro caso es muy simple, ya que como vemos tenemos que definir el lenguaje de programación usado (en este caso Go), su versión y los comandos que ejecutará Travis. Dicho comando, como se puede ver en el script, es make travis; ya que en nuestro makefile se ha creado una regla específica para esta tarea. En dicha regla vemos el comando correspondiente a la ejecución del contenedor de test almacenado en DockerHub creado en la práctica anterior. El ejecutar dicho contenedor y no la regla make test es conveniente ya que no será necesario especificar ninguna dependencia en nuestro script .travis.yml, ya que el entorno necesario para ejecutar los tests se encontrará en el contenedor construido a tal efecto.

## Otros sistemas de integración continua

A parte de Travis, también existen otros sistemas que nos permiten realizar las mismas tareas, como es el caso de Shippable, Jenkins, Circle CI, Bitrise...

En primer lugar, buscamos un servicio gratuito, por tanto algunas alternativas como TeamCity, Bitrise o Bamboo quedarían descartadas; ya que solo dejan un tiempo de prueba o tienen limitaciones a la hora de realizar la integración continua. De entre algunos de los servicios gratuitos, veamos sus características:

- Shippable: A diferencia de Travis este servicio se ejecuta sobre docker y no sobre una máquina virtual, lo que es más ventajoso ya que el estado del contenedor es persistente y no hará falta reinstalar elementos que ya estaban presentes o que no se han modificado. Además, es muy simple de configurar ya que con el script usado para Travis se puede ejecutar también la integración continua de Shippable. Como muchos otros servicios de este estilo, se pueden enlazar de manera sencilla a un repositorio de GitHub para que, cuando este se actualice, se ejecute de manera automática el proceso de CI.

- Jenkins: Esta alternativa, además de ser gratuita, es open source. Su funcionamiento se puede extender mediante plugins (por lo que en nuestro servicio de integración continua siempre tendremos lo justamente necesario, haciéndolo más eficiente); pero su configuración es más dificil ya que tenemos que iniciar nosotros mismos el servicio de manera local por ejemplo a través de un contenedor docker, a diferencia de otras alternativas que se ejecutan automáticamente y se configuran desde una interfaz web.

- Circle CI: Muy similar a Travis, nos permite enlazar nuestra cuenta de GitHub con la de Circle CI; con las ventajas que ello conlleva como es poder ejecutar de manera automática el servicio cada vez que se actualice el repositorio. Al igual que Travis, funciona en base a un archivo de configuración .yaml. Como punto en contra podemos comentar que funciona sobre una máquina virtual y no sobre un contenedor docker a diferencia de Shippable.

Teniendo en cuenta lo dicho anteriormente, vamos a utilizar Shippable; y para ello tendremos que seguir pasos muy similares a los seguidos para Travis, es decir, registrarnos en la página del [servicio](https://shippable.com) con nuestra cuenta de GitHub. Por último, en la configuración seleccionaremos qué repositorios queremos que se comprueben cada vez que se actualicen, y como nos sirve el mismo archivo para Travis que para Shippable no necesitaremos crear un nuevo archivo de configuración.
