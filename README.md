# 🔗 URL Shortener

[![Go Version](https://img.shields.io/badge/Go-1.25-00ADD8?logo=go)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen)](http://makeapullrequest.com)

**URL Shortener** — лёгкий и быстрый сервис сокращения ссылок на Go с веб-интерфейсом, REST API и SQLite.

## Интерфейс
<img width="535" height="552" alt="image" src="https://github.com/user-attachments/assets/5958f42e-fd96-4fc3-86d7-c6d1ff319557" />

---

## ✨ Возможности

| Функция | Описание |
|---------|----------|
| 🔗 Сокращение ссылок | Превращает длинные URL в короткие |
| 🏷️ Кастомные алиасы | Можно задать своё короткое имя |
| 🔐 Basic Auth | Только авторизованные пользователи создают ссылки |
| 🌐 Веб-интерфейс | Простая форма для создания ссылок |
| 📡 REST API | Для интеграции с другими сервисами |
| 💾 SQLite | Всё хранится в одном файле |
| 🧪 Тесты | Unit и интеграционные тесты |
| ⚙️ CI/CD | GitHub Actions автоматическая проверка |

---

## 🛠️ Технологии

| Компонент | Технология |
|-----------|------------|
| Язык | Go 1.25 |
| Роутер | chi/v5 |
| База данных | SQLite |
| Фронтенд | HTML + CSS + JS |
| Тесты | testify |
| CI/CD | GitHub Actions |

---

## 🚀 Быстрый старт

### Установка

```bash
# Клонируйте репозиторий
git clone https://github.com/Tim73916/url-shortener.git
cd url-shortener

# Установите зависимости
go mod download

# Создайте конфигурационный файл
cp config/example.yaml config/local.yaml

Настройка
Отредактируйте config/local.yaml:

yaml
env: "local"
storage_path: "./storage/storage.db"
http_server:
  address: "localhost:8080"
  timeout: 4s
  idle_timeout: 60s
  user: "Marshal"      # ваш логин
  password: "mypass"   # ваш пароль

# Через Makefile
make run

# Или напрямую
go run cmd/url-shortener/main.go
После запуска откройте в браузере: http://localhost:8080
