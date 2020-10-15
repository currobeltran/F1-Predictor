# Justificación del uso de herramientas

- Lenguaje: Se ha utilizado Go básicamente por su buen rendimiento (muy parecido al de C y C++ debido a que son muy similares en el proceso de compilación) y por su facilidad a la hora de aprender el lenguaje desde 0 como es el caso; ya que la sintaxis es bastante simple.

- Base de datos: La decisión de utilizar phpMyAdmin viene relacionada sobre todo por su gran funcionalidad ya que nos permite, a través de su interfaz web, crear o eliminar bases de datos, tablas y campos, incluyendo el poder ejecutar sentencias SQL desde la interfaz. Además es una alternativa gratuita y open source. La base de datos se usará en nuestro proyecto para almacenar los distintos datos de pilotos, circuitos, grandes premios...

- Sistema de log: Hemos decidido utilizar LogRocket porque es un sistema gratuito y que tiene una instalación muy simple, además de tener una interfaz web muy fácil de manejar. Entre sus funciones podemos destacar que se puede ir a cualquier momento pasado durante la monitorización del sistema y ver que eventos han ido ocurriendo en cada momento, pudiendo detectar de manera más ágil los fallos que hayan podido ocurrir.

- Herramienta de construcción: Debido a que el lenguaje escogido utiliza un gestor de tareas implícito en el compilador, no existe ningún task runner específico que podamos instalar. Es por ello que, a la hora de realizar tareas de manera automática, se recurre al uso de make (task runner genérico), aunque no es muy útil cuando se pasan tests o se compila el proyecto, ya que las reglas correspondientes solo tendrán un comando; el correspondiente para realizar estas acciones a través del compilador de go. En cambio, sí nos será más útil cuando necesitemos instalar dependencias en nuestro proyecto.
