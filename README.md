# Trabajo-de-cursada Gestor de Trabajos Practicos - Persistencia (TP2)
Trabajo de cursada de Belmonte y Hobiague
El poryecto es un sistema de gestion y seguimiento de entregas academicas para registra, consultar y administar los estados de los trabajos practicos de la carrera universitaria

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

Ejecucion de test:
Para verificar el proyecto se debe realizar el siguiente comando dentro de la carpeta principal:
make test
