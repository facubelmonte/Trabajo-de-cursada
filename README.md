# Trabajo-de-cursada Gestor de Trabajos Practicos - Renderiado del lado del sevidor con Templ
Trabajo de cursada de Belmonte y Hobiague
El proyecto es un sistema de gestion y seguimiento de entregas academicas para registra, consultar y administar los estados de los trabajos practicos de la carrera universitaria

La tabla se compone:
- id: Identificador unico
- titulo: Nombre descriptivo del trabajo practico
- materia: Asignatura corresponiente
- fecha_entrega: Fecha limite de la entrega
- estado: Estado actual del TP
- created-at: marca temporal de creacion con zona horaria

Ejecucion de test de tp5:
En la rama tp5, ejecutar el siguiente comando que realiza todas las tareas previas, levanta la aplicacion y al finalizar detiene y elimina contenedores y volumenes:
make run-app
Luego abrir: http://localhost:8080/

