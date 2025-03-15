package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

var res int
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
	// Получаем параметры a и b из URL
	// и Преобразуем строки в числа
	a, err := strconv.Atoi(r.URL.Query().Get("a"))
	if err != nil {
		http.Error(w, "Invalid parameter 'a'", http.StatusBadRequest)
		return false, 0, 0
	}

	b, err := strconv.Atoi(r.URL.Query().Get("b"))
	if err != nil {
		http.Error(w, "Invalid parameter 'a'", http.StatusBadRequest)
		return false, 0, 0
	}

	return true, a, b
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

func main() {
	// Регистрируем обработчики для URL
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/api/divide", divApiHandler)
	http.HandleFunc("/api/plus", addApiHandler)
	http.HandleFunc("/api/minus", subApiHandler)
	http.HandleFunc("/api/multiply", mulApiHandler)
	http.HandleFunc("/divide", divHandler)
	http.HandleFunc("/minus", subHandler)
	http.HandleFunc("/multiply", mulHandler)
	http.HandleFunc("/plus", addHandler)

	// Запускаем сервер на порту 8080
	var sPort = fmt.Sprintf("%d", port)
	fmt.Printf("Server is running on port %d...\n", port)
	if err := http.ListenAndServe(":"+sPort, nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
