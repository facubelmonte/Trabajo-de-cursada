# Trabajo-de-cursada Gestor de Trabajos Practicos - Exponiendo la logica como API (TPe)
Trabajo de cursada de Belmonte y Hobiague
El proyecto es un sistema de gestion y seguimiento de entregas academicas para registra, consultar y administar los estados de los trabajos practicos de la carrera universitaria

La tabla se compone:
- id: Identificador unico
- titulo: Nombre descriptivo del trabajo practico
- materia: Asignatura corresponiente
- fecha_entrega: Fecha limite de la entrega
- estado: Estado actual del TP
- created-at: marca temporal de creacion con zona horaria

El sistema implementa una arquitectura de persistencia basada en postgreSQL y generacion de codigo con sqlc
SQL-Fist y Type-Safety: A diferencia de los ORMs tradicionales basados en reflexión, sqlc toma las consultas declaradas en db/queries/queries.sql y el DDL en db/schema/schema.sql, compilando interfaces de Go fuertemente tipadas y seguras en tiempo de compilación.
Driver de Conectividad: Utiliza la librería estándar database/sql junto al driver github.com/lib/pq.
Aislamiento en Testing: Las pruebas se ejecutan contra un contenedor efímero de PostgreSQL orquestado mediante Docker Compose, garantizando bases de datos limpias e independientes en cada corrida.

Ejecucion de test de tp2:
Para verificar el proyecto se debe realizar el siguiente comando dentro de la carpeta principal:
make test

Ejecucion de test de tp3:
En el directorio Trabajo-de-cursada, se debe bajar la branch tp3, y acceder a ella
Luego de acceder, se debe iniciar el servidor mediante el comando: go run main.go
Ademas se debe levantar en otra terminal, o antes de levantar el servidor, el docker con la base de datos, mediante el comando: docker compose up -d
Y para finalizar, se ejecuta las pruebas mmediante el comando ./request.bash
