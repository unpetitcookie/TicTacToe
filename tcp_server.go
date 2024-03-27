package main

import (
    "fmt"
    "net"
)

func handleConnection(conn net.Conn) {
    defer conn.Close()

    buf := make([]byte, 1024)
    _, err := conn.Read(buf)
    if err != nil {
        fmt.Println("Erreur lors de la lecture des données:", err)
        return
    }
    fmt.Println("Données reçues:", string(buf))
}

func server() {
    address := "localhost:8080"

    listener, err := net.Listen("tcp", address)
    if err != nil {
        fmt.Println("Erreur lors du démarrage du serveur:", err)
        return
    }
    defer listener.Close()

    fmt.Println("http://localhost:8080")

    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Erreur lors de l'acceptation de la connexion:", err)
            continue
        }

        go handleConnection(conn)
    }
}
