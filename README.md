# Веб-сервис подсчёта арифметических выражений

## Описание
Этот проект реализует веб-сервис, который вычисляет арифметические выражения, переданные пользователем через HTTP-запрос.

## Запуск сервиса

2. Склонируйте проект с GitHub:
    ```bash
    git clone (https://github.com/imil23456/Calculate-service/tree/master)
    ```
3. Перейдите в папку проекта и запустите сервер:
    ```bash
    go run ./cmd/calc_service/...
    ```
4. Сервис будет доступен по адресу: [http://localhost:8080/api/v1/calculate](http://localhost:8080/api/v1/calculate).

## Примеры использования

## Успешный запрос

**Пример запроса через curl:**

```bash
curl http://localhost:8080/api/v1/calculate -X POST --header "Content-Type: application/json" --data "{\"expression\": \"2+2*2\"}"
```
**Пример ответа:**

```json
{
    "result": "6"
}
```
**HTTP-код ответа:** `200` (всё прошло успешно).

### Ошибка 422 (Некорректное выражение)

**Пример запроса через curl:**

```bash
curl http://localhost:8080/api/v1/calculate -X POST --header "Content-Type: application/json" --data "{\"expression\": \"25-0.251***\"}"
```

**Пример ответа:**

```json
{
    "error": "Expression is not valid"
}
```

**HTTP-код ответа:** `422` (Unprocessable Entity).



