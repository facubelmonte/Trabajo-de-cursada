package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
    "os"

    _ "github.com/lib/pq"
    "trabajo-cursada/db/sqlc"
    "trabajo-cursada/handlers"
)

func main() {
    dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
        os.Getenv("POSTGRES_USER"),
        os.Getenv("POSTGRES_PASSWORD"),
        os.Getenv("POSTGRES_DB"),
    )

    conn, err := sql.Open("postgres", dsn)
    if err != nil {
        log.Fatalf("failed to connect to DB: %v", err)
    }
    defer conn.Close()

    handlers.Queries = db.New(conn)
    fmt.Println("Base de datos configurada y sqlc listo")

    http.Handle("/", http.FileServer(http.Dir("./static")))
    http.HandleFunc("/tps", handlers.TpsHandler)
    http.HandleFunc("/tps/", handlers.TpHandler)

    fmt.Println("Servidor escuchando en localhost 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
