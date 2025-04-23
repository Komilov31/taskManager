Сервис для асинхронного выполнения долгих I/O-bound задач (3–5 минут) с HTTP API.
Хранит состояние в памяти, не требует внешних зависимостей (БД, очередей).
Особенности

✅ Асинхронная обработка – задачи выполняются в фоне без блокировки клиента.
✅ Масштабируемая архитектура – легко добавлять новые типы задач.

API Endpoints
1. Создать задачу

POST /newtask

Пример запроса:

bash
```
curl -X POST http://localhost:8080/newtask
```
Ответ:
```
json
{
  "id": 1,
  "status": "pending"
  "result": "not ready yet"
}
```
2. Проверить статус

GET /status?id=<task_id>

Пример запроса:

bash
```
curl http://localhost:8080/status?id=<task_id>
```
Ответ:
```
json
{
  "id": 1,
  "status": "finished",
}
```
3. Проверить статус

GET /result?id=<task_id>

Пример запроса:
bash
```
curl http://localhost:8080/result?id=<task_id>
```
Ответ:
```
json
{
  "id": 1,
  "result": "this is result of your task",
}
```
Архитектура
Слои

    Transport – HTTP-роутинг (gorilla/mux)

    Service – бизнес-логика (менеджер задач).

    Repository – хранение состояния задач в памяти (sync.Map + RWMutex).
