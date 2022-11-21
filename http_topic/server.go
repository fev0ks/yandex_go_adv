package main

import (
	"log"
	"net/http"
)

type MyHandler struct {
	Templ []byte
}

func (h MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write(h.Templ)
}

func main() {
	//handler1 := MyHandler{
	//	Templ: []byte("Hola, Mundo"),
	//}
	//handler2 := http.HandlerFunc(SimpleFunc)
	//http.Handle("/qweqw", http.NotFoundHandler())
	http.HandleFunc("/", SimpleFunc)
	//mux := http.NewServeMux()
	//mux.HandleFunc("/api/m/", SimpleFunc3)
	//mux.HandleFunc("/api/m/q/", SimpleFunc2)

	server := &http.Server{
		//Handler: handler1,
		//Handler: middleware(mux),
		Addr: "localhost:8080",
	}

	// вызов ListenAndServe — блокирующий, последний в программе
	// возникающие ошибки на серверных машинах пишут в системный лог,
	// а не в стандартную консоль ошибок,
	// поэтому обычно вызывают вот так
	log.Println("Start server")
	log.Fatal(server.ListenAndServe())

	//http.Handle("/kek", http.NotFoundHandler())
	//log.Fatal(http.ListenAndServe(":8080", nil))
}

func SimpleFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Blabla1", "123")
	w.Write([]byte("Something"))
	w.Header().Add("Blabla2", "123")
	w.Header().Set("Blabla3", "123")
}

func SimpleFunc2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Something2"))
}
func SimpleFunc3(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Something3"))
}

// middleware принимает параметром Handler и возвращает тоже Handler.
func middleware(next http.Handler) http.Handler {
	// собираем Handler приведением типа
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// здесь пишем логику обработки
		// например, разрешаем запросы cross-domain
		// w.Header().Set("Access-Control-Allow-Origin", "*")
		// ...
		// замыкание — используем ServeHTTP следующего хендлера
		next.ServeHTTP(w, r)
	})
}
