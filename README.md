# Shortify 🚀

Shortify — это сервис для создания коротких URL. Он использует MongoDB для хранения информации о коротких и длинных URL и предоставляет API для создания и перенаправления на короткие URL.

## Установка и запуск 🛠️

### 1. Клонирование репозитория

Склонируйте этот репозиторий на вашу локальную машину:

```bash
git clone https://github.com/yourusername/shortify.git
cd shortify
```
### 2. Создание файла окружения

Создайте файл .env в корневом каталоге и добавьте следующее:
```
MONGO_INITDB_ROOT_USERNAME=root
MONGO_INITDB_ROOT_PASSWORD=123qwe
MONGO_URI=mongodb://root:123qwe@mongodb:27017
MONGO_PORTS=27017
WEB_PORTS=8080
```
### 3. Запуск приложения
Используйте Docker Compose для запуска приложения и базы данных MongoDB:
```
docker-compose up --build
```
### 4. Доступ к документации API
Документация Swagger доступна по адресу http://localhost:8080/swagger/index.html. 📜

## API Эндпоинты 🌐

### Создать короткий URL

- **Метод:** `POST /api/v1/shortify`
- **Описание:** Создает новый короткий URL из длинного.
- **Параметры:**
  - `url` (параметр запроса): Длинный URL для укорочения.
- **Ответы:**
  - `201 Created`: Успешно создан короткий URL.
  - `400 Bad Request`: Неверный URL или отсутствующий параметр.

### Перенаправить короткий URL

- **Метод:** `GET /{shortUrl}`
- **Описание:** Перенаправляет на длинный URL, используя короткий URL.
- **Параметры:**
  - `shortUrl` (параметр пути): Короткий URL для перенаправления.
- **Ответы:**
  - `302 Found`: Успешное перенаправление.
  - `404 Not Found`: Короткий URL не найден.

## Логирование 📜

Логи обрабатываются с использованием Zap для структурированного и эффективного логирования.  
Стандартные логи Gin отключены, чтобы избежать избыточности. 🗑️

## Вклад 🤝

1. Сделайте форк репозитория
2. Создайте новую ветку для вашей функции или исправления ошибки
3. Отправьте пулл-реквест

Мы приветствуем вклад от сообщества! 🎉
