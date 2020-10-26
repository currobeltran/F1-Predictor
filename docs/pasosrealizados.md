# Pasos realizados para el proyecto

## Hito 0: Comienzo del proyecto

Este es el hito inicial de nuestro proyecto, donde creamos el repositorio del mismo. En él se incluyeron documentos básicos como:

- .gitignore (correspondiente al lenguaje de programación que usaremos, Go) con el que se evitará incluir archivos al seguimiento de nuestro repositorio como pueden ser binarios o archivos de compilación

- Licencia de uso para que pueda ser utilizada como software libre.

- README.md en el que se describía de manera breve en qué iba a consistir nuestro proyecto, además de incluir información sobre como se configuró el repositorio con una llave pública y la justificación de por qué se usan las herramientas que necesita el proyecto.

## Hito 1: Estructura general del proyecto

En este hito comenzamos a ver algo de código, ya que se incluye una primera versión de como se organizará la jerarqía de clases de nuestro proyecto, incluido en [esta carpeta](https://github.com/currobeltran/F1-Predictor/tree/master/src). Se dispone de un paquete con todos los fuentes relacionados con la aplicación.

Se especifica aún más en la documentación, donde se incluyen las historias de usuario (que definirán las funcionalidades que tendrá nuestro proyecto), se modifica la descripción adquiriendo un caracter más motivacional y se incluyen algunas herramientas nuevas a la justificación (que no tienen porque ser las que se usen finalmente, esos puntos tienen un caracter informativo de cara a saber que necesita nuestro proyecto para ser desplegado correctamente en una primera instancia). Por último, se añade también este documento que funcionará como un "cuaderno de bitácora", donde se contará que pasos se han seguido para llegar a la consecución de los objetivos de la asignatura, demostrando de la misma manera el conocimiento de los mismos.

## Hito 2: Integración continua

En este hito aprendemos a desarrollar un servicio para la nube aplicando el desarrollo basado en pruebas. Esta modalidad de desarrollo básicamente se basa en que todo el código que se incorpore al proyecto debe estar previamente testeado, de manera que cumpla con las especificaciones definidas previamente para dicha función (en las historias de usuario). Para ello, se deben realizar test unitarios, que introducirán valores en el método y comprobarán que la salida a esos valores es el resultado esperado, es decir, que cumple con el objetivo que queremos conseguir.

En nuestro caso específico, hemos creado un archivo por cada clase donde se contienen las distintas pruebas que se realizarán a los métodos de una clase. Estos archivos de prueba se encuentran en el mismo directorio que el fichero que contiene la clase a comprobar.

Estos test deberán ejecutarse de manera automática, por lo que necesitaremos utilizar una herramienta de construcción, las cuales nos permiten realizar tareas que requieran de uno o varios pasos determinados ejecutando un único comando. En nuestro caso particular, dentro de un makefile tenemos 2 reglas; una para realizar los test y otra para comprobar la correcta compilación del proyecto (test y build respectivamente).

Por tanto y como resumen, en este hito se han añadido tantos ficheros de test como clases hay en nuestro proyecto (donde se contienen las pruebas para los métodos de cada clase), y un archivo makefile para ejecutar de manera automática (y por el momento) el testeo del código y la compilación del mismo.

## Hito 3: Creación de un contenedor para pruebas

En este hito el principal punto que se ha tratado es la creación de un contenedor para realizar de manera automática las pruebas del código del proyecto. Un contenedor permite ejecutar una aplicación de manera aislada dentro de un host, que puede tener otros contenedores en funcionamiento. Este concepto hoy en día es vital en las infraestructuras virtuales, ya que nos permite que dentro de una máquina se estén ejecutando varias aplicaciones de manera aislada entre sí, como si de máquinas virtuales se tratase pero siendo estas últimas más pesadas y menos portables.

Para esta tarea se ha utilizado Docker, por tanto en nuestro proyecto incluiremos un documento llamado dockerfile con el que se creará el contenedor añadiendo todo lo necesario para la realización de su tarea (los pasos que se siguen, además de la justificación de las especificaciones técnicas, están en este [enlace](docker.md)).

Este contenedor, además, se ha subido a un repositorio público (tanto en Docker Hub como en GitHub Container Registry) lo que nos permitirá, por ejemplo, que cualquier persona pueda acceder al mismo ya construido y que además, puedan ejecutarlo de manera sencilla con un simple comando desde el directorio raíz de nuestro proyecto. Para esto se ha tenido que crear un repositorio en Docker Hub y enlazarlo con el de GitHub; ya que se obtendrá el dockerfile directamente desde este último, además de actualizarse nuestro repo de Docker Hub cada vez que se realiza una actualización en el proyecto. Para configurar GitHub Container Registry se ha creado una GitHub Action (teniendo como base un template llamado Publish Docker Container) que automáticamente cada vez que se actualiza el repositorio en GitHub construye el contenedor y lo almacena en GitHub Container Registry.

En cuanto al avance sobre las funcionalidades del código, se ha comenzado a realizar los métodos correspondientes a las predicciones sobre los eventos que pueden ocurrir en un Gran Premio.
