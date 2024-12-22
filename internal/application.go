package application

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/SashaMorkovkin/Final_task_1/pkg/rpn"
)

type Config struct {
	Addr string
}

type Application struct {
	config *Config
}

type Request struct {
	Expression string `json:"expression"`
}

type RequestBody struct {
	Expression string `json:"expression"`
}

type Response struct {
	Result string `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
}

func New() *Application {
	return &Application{
		config: ConfigFromEnv(),
	}
}

func ConfigFromEnv() *Config {
	config := new(Config)
	config.Addr = os.Getenv("PORT")
	if config.Addr == "" {
		config.Addr = "8080"
	}
	return config
}

func (a *Application) Run() error {
	for {
		log.Println("input expression")
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Println("failed to read expression from console")
		}
		text = strings.TrimSpace(text)
		if text == "exit" {
			log.Println("aplication was successfully closed")
			return nil
		}
		result, err := rpn.Calc(text)
		if err != nil {
			log.Println(text, " calculation failed wit error: ", err)
		} else {
			log.Println(text, "=", result)
		}
	}
}

func EvaluateExpression(expression string) (float64, error) {
	value, err := rpn.Calc(expression)
	if err != nil {
		return 0, fmt.Errorf("неожиданный тип результата")
	}
	return value, nil
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "метод не разрешен", http.StatusMethodNotAllowed)
		return
	}

	var request Request
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "неправильный формат запроса", http.StatusBadRequest)
		return
	}

	result, err := EvaluateExpression(request.Expression)
	if err != nil {
		if err.Error() == "неверное выражение" {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(Response{Error: "Expression is not valid"})
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(Response{Error: "Internal server error"})
		}
		return
	}

	response := Response{Result: fmt.Sprintf("%.2f", result)}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (a *Application) RunServer() error {
	http.HandleFunc("/", calculateHandler)
	return http.ListenAndServe(":"+a.config.Addr, nil)
}
