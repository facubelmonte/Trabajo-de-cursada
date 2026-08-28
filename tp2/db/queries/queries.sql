-- name: CrearTP :one
INSERT INTO tps (titulo, materia, fecha_entrega, estado)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: ObtenerTP :one
SELECT * FROM tps WHERE id = $1;

-- name: ListarTP :many
SELECT * FROM tps ORDER BY fecha_entrega ASC;

-- name: ActualizarTP :exec
UPDATE tps 
SET titulo = $1, materia = $2, fecha_entrega = $3, estado = $4 
WHERE id = $5;

-- name: EliminarTP :exec
DELETE FROM tps WHERE id = $1;
