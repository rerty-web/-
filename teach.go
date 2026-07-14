package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		html := `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>Правила чата</title>
    <style>
        body { font-family: Arial, sans-serif; max-width: 700px; margin: 50px auto; padding: 20px; background: #f9f9f9; }
        h1 { color: #222; }
        li { margin: 10px 0; }
    </style>
</head>
<body>
    <h1>📜 Правила нашего чата</h1>
    <ol>
        <li>Будьте вежливы друг к другу</li>
        <li>Не спамить и не рекламировать</li>
        <li>Обсуждаем темы по делу</li>
        <li>За нарушение — предупреждение или бан</li>
    </ol>
    <p><small>Правила обновляются при необходимости</small></p>
</body>
</html>
`
		fmt.Fprint(w, html)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Сервер запущен на порту", port)
	http.ListenAndServe(":"+port, nil)
}
