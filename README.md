# F1-Predictor
Proyecto para la asignatura de Infraestructura Virtual de 4º curso de Ingeniería Informática

## ¿Por qué F1 Predictor? Motivación del proyecto

En Fórmula 1, como en cualquier otra modalidad deportiva, las estadísticas juegan un papel fundamental. Todos los equipos están calculando probabilidades sobre que suceso puede o no ocurrir a cada momento, para anticiparse a los rivales y sacar el mejor resultado posible. Por ejemplo, en una situación de bandera amarilla o de Safety Car, si el equipo está preparado para ello con la estrategia más adecuada, puede hacerle conseguir un resultado que en condiciones normales no podría ser posible. También sería útil para dichos equipos deducir como enfocar un Gran Premio dependiendo de estadísticas relacionadas con el circuito, como pueden ser los adelantamientos, ya que no será igual correr en un circuito donde históricamente se adelanta muy poco que en otro donde se producen muchos eventos de este tipo.

Además, los más aficionados suelen buscar estadísticas antes de cada Gran Premio, para ver datos históricos de dicho circuito, records marcados, los ganadores de los grandes premios anteriores, entre otros muchos datos con el objetivo de aumentar sus conocimientos sobre la categoría.

Nuestra aplicación tratará información de manera básica, pero intentando cumplir los objetivos mencionados aquí arriba sentando una base que en un futuro podría mejorar a un servicio más especializado; con métodos para el cálculo de probabilidades de eventos más precisos.

## Microservicios

En esta parte del proyecto vamos a diseñar y testear un microservicio. En primer lugar, tenemos que elegir un microframework para realizar la implementación de dicho microservicio. Estudiando distintas alternativas dentro de las posibilidades que ofrece Go, se ha decidido usar el módulo Pastis; el cual provee un uso muy simple y las funcionalidades justas y necesarias para este punto del proyecto. Básicamente, si utilizamos este microframework podremos:

- Diseñar rutas para acceder a distintos recursos a través de una función. El objeto que controle las peticiones tendrá distintos métodos para las acciones HTTP que se pueden realizar. Estas rutas seguirán una expresión regular donde, si coincide con el URI de la petición, se llamará a la función correspondiente. De esta manera podemos obtener las variables del URI de manera sencilla posteriormente y el comportamiento del microservicio estará mucho mejor definido.

- Definir recursos, por tanto podemos simplificar aún más la definición de las acciones en relación a las rutas ya que podemos asociarlas a objetos (en nuestro caso structs) y que las acciones de las rutas sean los correspondientes métodos GET, PUT, POST... de dicho objeto.

- Dispone de la capacidad de crear un middleware, con el cual por ejemplo podremos crear el log de la ejecución del microservicio (aunque ya dispone de una API propia destinada a tal fin, pero utilizaremos el middleware para controlar de manera precisa la funcionalidad del sistema de log).

Algunos puntos más a favor es que disponen de una documentación muy clara y es software libre disponible a través de [GitHub](https://github.com/guregodevo/pastis). Como punto en contra podemos decir es un proyecto que lleva sin actualizarse desde 2017, por tanto puede que en futuras versiones del lenguaje este proyecto esté obsoleto; sin embargo resulta muy útil para implementar nuestro microservicio de la manera más eficiente posible.

Una opción muy similar a la elegida es [Goji](https://github.com/goji/goji), ya que sigue el mismo proceso para definir rutas, pero no nos permite definir recursos, un punto importante ya que con Pastis podremos estructurar nuestro código correctamente. Sería preferencial escogerla si el proyecto estuviese activo, pero desde 2019 no recibe ningún commit y a largo plazo esta solución tendría el mismo problema que el anterior: puede quedar obsoleto.

Se dejan de lado opciones que no concuerdan con el modelo de desarrollo que estamos siguiendo como podría ser FlowBase o TarsGo (no están diseñados para los paradigmas de programación que se está siguiendo) y Nitro, que actualmente se considera inestable según su documentación debido a un remodelado de la aplicación.

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