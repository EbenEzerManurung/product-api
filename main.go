package main

import (
    "flag"
    "fmt"
    "log"
    "os"

    "product-api/config"
    "product-api/database"
    "product-api/route"
    "product-api/seeder"
)


func main() {
    if len(os.Args) < 2 {
        printUsage()
        os.Exit(1)
    }

    command := os.Args[1]

    switch command {
    case "serve":
        runServer()
    case "fresh:migrate":
        runFreshMigrate()
    default:
        printUsage()
        os.Exit(1)
    }
}

func printUsage() {
    fmt.Println("Penggunaan:")
    fmt.Println("  go run main.go serve                     # Jalankan HTTP server")
    fmt.Println("  go run main.go fresh:migrate             # Drop semua tabel lalu migrate ulang")
    fmt.Println("  go run main.go fresh:migrate --seeder    # Drop, migrate ulang, lalu jalankan seeder")
}

func runFreshMigrate() {
    fs := flag.NewFlagSet("fresh:migrate", flag.ExitOnError)
    seedFlag := fs.Bool("seeder", false, "Jalankan seeder setelah migrasi")
    _ = fs.Parse(os.Args[2:])

    cfg := config.Load()
    db := database.Connect(cfg)

    database.FreshMigrate(db)

    if *seedFlag {
        seeder.RunAll(db)
    }

    fmt.Println("✅ Fresh migration berhasil.")
}

func runServer() {
    cfg := config.Load()
    db := database.Connect(cfg)

    router := route.SetupRouter(db, cfg)

    addr := fmt.Sprintf(":%s", cfg.ServerPort)
    log.Printf("🚀 Server berjalan di %s", addr)
    if err := router.Run(addr); err != nil {
        log.Fatalf("Gagal menjalankan server: %v", err)
    }
}