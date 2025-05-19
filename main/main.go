package main

import (
        "encoding/json"
        "fmt"
        "log"
        "net/http"
        )

// Kullanıcı modelimiz
        type User struct {
        Name  string `json:"name"`
        Email string `json:"email"`
        }

// Ana sayfa handler
        func homeHandler(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Merhaba, Go Web Sunucusuna Hoş Geldiniz!")
        }

// JSON verisi alan handler
        func userHandler(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
        http.Error(w, "Sadece POST metodu destekleniyor.", http.StatusMethodNotAllowed)
        return
        }

        var user User
        err := json.NewDecoder(r.Body).Decode(&user)
        if err != nil {
        http.Error(w, "Geçersiz JSON", http.StatusBadRequest)
        return
        }

        response := fmt.Sprintf("Kullanıcı alındı: %s <%s>", user.Name, user.Email)
        fmt.Fprintln(w, response)
        }

        func main() {
        http.HandleFunc("/", homeHandler)
        http.HandleFunc("/user", userHandler)

        fmt.Println("Sunucu 8080 portunda çalışıyor...")
        log.Fatal(http.ListenAndServe(":8080", nil))
        }
