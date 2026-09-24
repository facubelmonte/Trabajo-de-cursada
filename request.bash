#!/usr/bin/env bash

echo "1 crear un nuevo tp"
curl -i -X POST http://localhost:8080/tps \
     -H "Content-Type: application/json" \
     -d '{"titulo": "tp3", "materia": "Programacion Web", "fecha_entrega": "2026-09-30", "estado": "Pendiente"}'
echo -e "\n"

echo "2 obtener todos los tps"
curl -i http://localhost:8080/tps
echo -e "\n"

echo "3 obtener el tp con id 1"
curl -i http://localhost:8080/tps/1
echo -e "\n"

echo "4 actualizar el tp con id 1"
curl -i -X PUT http://localhost:8080/tps/1 \
     -H "Content-Type: application/json" \
     -d '{"titulo": "tp3 terminado", "materia": "Programacion Web", "fecha_entrega": "2026-09-30", "estado": "Entregado"}'
echo -e "\n"

echo "5 verificar actualizacion"
curl -i http://localhost:8080/tps/1
echo -e "\n"

echo "6 eliminar el tp con id 1"
curl -i -X DELETE http://localhost:8080/tps/1
echo -e "\n"

echo "7 verificar eliminacion"
curl -i http://localhost:8080/tps/1
echo -e "\n"
