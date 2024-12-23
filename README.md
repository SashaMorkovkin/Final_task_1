# Веб-калькулятор
## Структура проекта:
+ cmd/ - директория с файлом main.go(точка входа в программу)
+ internal/ - директория где храниться сервер
+ pkg/rpn/ - директория где храниться код калькулятора

## Запуск
1. Склонируйте проект с github git clone ```https://github.com/SashaMorkovkin/Final_task_1```
2. Перейдите в папку с проектом и запустите сервер ```go run ./cmd/main.go```
> PS: Сервер работает на порте 8080

## Примеры запросов
+ Пример №1
    + Команда:
        >curl --location 'http://localhost:8080/api/v1/calculate' \
        >--header 'Content-Type: application/json' \
        >--data '{
        >"expression": "2+2*2"
        ?}'
    + Ответ:
        {"result": "6"}
+ Пример №2
    + Команда:
        >curl --location 'http://localhost:8080/api/v1/calculate' \
        >--header 'Content-Type: application/json' \
        >--data '{
        >"expression": "ФЫФЫЫ"
        >}'
    + Ответ:
        >{"error":"Expression is not valid"}
+ В других случаях ответ :
    >{"error":"Internal server error"}
