#!/usr/bin/env bash

#Funcion para verificar que el codigo HTTP coincida con el esperado
verificar_codigo() {
    paso="$1"
    esperado="$2"
    obtenido="$3"
    if [ "$obtenido" -eq "$esperado" ]; then
        echo "[PASS] $paso (HTTP $obtenido)"
    else
        echo "[FAIL] $paso: se esperaba HTTP $esperado pero devolvio $obtenido"
        rm -f respuesta.json
        exit 1
    fi
}

# 1. Crear un nuevo TP (debe devolver 201 y el titulo creado)
codigo=$(curl -s -o respuesta.json -w "%{http_code}" -X POST http://localhost:8080/tps \
     -H "Content-Type: application/json" \
     -d '{"titulo": "tp3", "materia": "Programacion Web", "fecha_entrega": "2026-09-30", "estado": "Pendiente"}')
verificar_codigo "1. Crear un nuevo TP (POST /tps)" 201 "$codigo"
grep -q '"titulo":"tp3"' respuesta.json || { echo "[FAIL] El JSON no contiene el titulo creado"; exit 1; }

# 2.Obtener todos los TPs (debe devolver 200)
codigo=$(curl -s -o respuesta.json -w "%{http_code}" http://localhost:8080/tps)
verificar_codigo "2. Obtener todos los TPs (GET /tps)" 200 "$codigo"
grep -q '"id":1' respuesta.json || { echo "[FAIL] La lista no contiene el TP 1"; exit 1; }

# 3.Obtener el TP con ID 1 (debe devolver 200)
codigo=$(curl -s -o respuesta.json -w "%{http_code}" http://localhost:8080/tps/1)
verificar_codigo "3. Obtener TP con ID 1 (GET /tps/1)" 200 "$codigo"
grep -q '"estado":"Pendiente"' respuesta.json || { echo "[FAIL] El estado inicial no es Pendiente"; exit 1; }

# 4.Actualizar el TP con ID 1 (debe devolver 200 y estado Entregado)
codigo=$(curl -s -o respuesta.json -w "%{http_code}" -X PUT http://localhost:8080/tps/1 \
     -H "Content-Type: application/json" \
     -d '{"titulo": "tp3 terminado", "materia": "Programacion Web", "fecha_entrega": "2026-09-30", "estado": "Entregado"}')
verificar_codigo "4. Actualizar TP con ID 1 (PUT /tps/1)" 200 "$codigo"
grep -q '"estado":"Entregado"' respuesta.json || { echo "[FAIL] El PUT no actualizo el estado a Entregado"; exit 1; }

# 5.Verificar actualizacion con GET (debe devolver 200 y mantener los cambios)
codigo=$(curl -s -o respuesta.json -w "%{http_code}" http://localhost:8080/tps/1)
verificar_codigo "5. Verificar actualizacion (GET /tps/1)" 200 "$codigo"
grep -q '"titulo":"tp3 terminado"' respuesta.json || { echo "[FAIL] El titulo actualizado no se guardo en la BD"; exit 1; }

# 6.Eliminar el TP con ID 1 (debe devolver 204 No Content)
codigo=$(curl -s -o respuesta.json -w "%{http_code}" -X DELETE http://localhost:8080/tps/1)
verificar_codigo "6. Eliminar TP con ID 1 (DELETE /tps/1)" 204 "$codigo"

# 7.Verificar eliminacion (debe devolver 404 Not Found)
codigo=$(curl -s -o respuesta.json -w "%{http_code}" http://localhost:8080/tps/1)
verificar_codigo "7. Verificar eliminacion (GET /tps/1 -> 404)" 404 "$codigo"

rm -f respuesta.json
echo "=== Todos los tests de la API pasaron correctamente (PASS) ==="
