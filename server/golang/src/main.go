package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

const APP_VERSION = "0.2.0"

var port = 8080

func calcAdd(a, b int) int {
	r := a + b
	return r
}

func calcSub(a, b int) int {
	r := a - b
	return r
}

func calcMultiply(a, b int) int {
	r := a * b
	return r
}

func calcDivide(a, b int) (int, bool, string) {
	if b == 0 {
		return 0, false, "Can't divide by Zero !"
	}
	r := a / b
	return r, true, ""
}

// Функция для преобразования параметров HTTP-запроса
func parseParams(r *http.Request, w http.ResponseWriter) (bool, int, int) {
	a, err := strconv.Atoi(r.URL.Query().Get("a"))
	if err != nil {
		http.Error(w, "Invalid parameter 'a'", http.StatusBadRequest)
		return false, 0, 0
	}

	b, err := strconv.Atoi(r.URL.Query().Get("b"))
	if err != nil {
		http.Error(w, "Invalid parameter 'b'", http.StatusBadRequest)
		return false, 0, 0
	}

	return true, a, b
}

// Структура для хранения результата и времени выполнения
type TaskResult struct {
	Result  int
	Elapsed time.Duration
}

// Функция для имитации долгого запроса
func longRunningTask(a int, resultChan chan<- TaskResult) {
	startTime := time.Now()
	var i_counter = 0
	var l_counter = 0
	for i_loop1 := 0; i_loop1 < a; i_loop1++ {
		// fmt.Printf("\nloop1: %#v", i_loop1)
		for i_loop2 := 0; i_loop2 < 32000; i_loop2++ {
			for i_loop3 := 0; i_loop3 < 32000; i_loop3++ {
				i_counter++
				l_counter++
				if i_counter > 50 {
					i_counter = 0
				}
			}
		}
	}
	elapsedTime := time.Since(startTime) // Вычисляем время выполнения
	resultChan <- TaskResult{Result: l_counter, Elapsed: elapsedTime}
}

func longHandler(w http.ResponseWriter, r *http.Request) {
	a, err := strconv.Atoi(r.URL.Query().Get("a"))
	if err != nil {
		http.Error(w, "Invalid parameter 'a'", http.StatusBadRequest)
		return
		// a = 5
	}

	resultChan := make(chan TaskResult)
	go longRunningTask(a, resultChan)

	res := <-resultChan

	response := fmt.Sprintf("a = %d , longRunningTask(a) = %d\nelapsed = %v\n", a, res.Result, res.Elapsed)
	w.Write([]byte(response))
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	ok, a, b := parseParams(r, w)
	if !ok {
		return // Если преобразование не удалось, ошибка уже отправлена
	}
	res := calcAdd(a, b)

	response := fmt.Sprintf("a = %d, b = %d\na + b = %d\n", a, b, res)
	w.Write([]byte(response))
}

func addApiHandler(w http.ResponseWriter, r *http.Request) {
	ok, a, b := parseParams(r, w)
	if !ok {
		return // Если преобразование не удалось, ошибка уже отправлена
	}
	res := calcAdd(a, b)

	// Формируем ответ
	response := map[string]int{"result": res}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jsonResponse))
}

func subHandler(w http.ResponseWriter, r *http.Request) {
	ok, a, b := parseParams(r, w)
	if !ok {
		return // Если преобразование не удалось, ошибка уже отправлена
	}
	res := calcSub(a, b)

	// Формируем ответ
	response := fmt.Sprintf("a = %d, b = %d\na - b = %d\n", a, b, res)
	w.Write([]byte(response))
}

func subApiHandler(w http.ResponseWriter, r *http.Request) {
	ok, a, b := parseParams(r, w)
	if !ok {
		return // Если преобразование не удалось, ошибка уже отправлена
	}
	res := calcSub(a, b)

	// Формируем ответ
	response := map[string]int{"result": res}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jsonResponse))
}

func divHandler(w http.ResponseWriter, r *http.Request) {
	ok, a, b := parseParams(r, w)
	if !ok {
		return // Если преобразование не удалось, ошибка уже отправлена
	}
	res, ok, info := calcDivide(a, b)
	if !ok {
		http.Error(w, info, http.StatusBadRequest)
		return // Если вычисление не удалось, ошибка уже отправлена
	}

	// Формируем ответ
	response := fmt.Sprintf("a = %d, b = %d\na / b = %d\n", a, b, res)
	w.Write([]byte(response))
}

func divApiHandler(w http.ResponseWriter, r *http.Request) {
	ok, a, b := parseParams(r, w)
	if !ok {
		return // Если преобразование не удалось, ошибка уже отправлена
	}
	res, ok, info := calcDivide(a, b)
	if !ok {
		http.Error(w, info, http.StatusBadRequest)
		return // Если вычисление не удалось, ошибка уже отправлена
	}

	// Формируем ответ
	response := map[string]int{"result": res}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jsonResponse))
}

func mulHandler(w http.ResponseWriter, r *http.Request) {
	ok, a, b := parseParams(r, w)
	if !ok {
		return // Если преобразование не удалось, ошибка уже отправлена
	}
	res := calcMultiply(a, b)

	// Формируем ответ
	response := fmt.Sprintf("a = %d, b = %d\na * b = %d\n", a, b, res)
	w.Write([]byte(response))
}

func mulApiHandler(w http.ResponseWriter, r *http.Request) {
	ok, a, b := parseParams(r, w)
	if !ok {
		return // Если преобразование не удалось, ошибка уже отправлена
	}
	res := calcMultiply(a, b)

	// Формируем ответ
	response := map[string]int{"result": res}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jsonResponse))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем hostname
	hostname, err := os.Hostname()
	if err != nil {
		http.Error(w, "Failed to get hostname", http.StatusInternalServerError)
		return
	}

	// Получаем текущее время в UTC
	currentTime := time.Now().UTC().Format(time.RFC3339)

	// Формируем ответ
	response := fmt.Sprintf("Hello World from Golang!\n DateTime (UTC): \"%s\"\n My hostname is \"%s\"\n", currentTime, hostname)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(response))
}

func startServer() *http.Server {
	// Запускаем сервер на порту 8080
	log.Printf("Server is starting on port %d ...", port)
	server := &http.Server{
		Addr: ":" + strconv.Itoa(port),
	}

	// Регистрируем обработчики для URL
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/long", longHandler)
	http.HandleFunc("/api/divide", divApiHandler)
	http.HandleFunc("/api/plus", addApiHandler)
	http.HandleFunc("/api/minus", subApiHandler)
	http.HandleFunc("/api/multiply", mulApiHandler)
	http.HandleFunc("/divide", divHandler)
	http.HandleFunc("/minus", subHandler)
	http.HandleFunc("/multiply", mulHandler)
	http.HandleFunc("/plus", addHandler)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %s\n", err)
		}
	}()

	return server
}

func main() {
	// runtime.GOMAXPROCS(4)  // equal "export GOMAXPROCS=4; ./app.bin.local"
	// log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Printf("Calc-Server (v%s, golang)", APP_VERSION)
	server := startServer()

	// Канал для получения сигналов ОС
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Ожидаем сигнал
	sig := <-sigChan
	log.Printf("Received signal: %v. Shutting down gracefully...\n", sig)

	// Создаем контекст с таймаутом 3 секунды
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Пытаемся завершить сервер
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Server shutdown error: %v\n", err)
	} else {
		log.Println("Server stopped gracefully")
	}
}
