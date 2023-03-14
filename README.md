# Тестовое задание в компанию Interweb

Задача - написать телеграм бота на Go. Использовать Postgres для хранения данных.

Бот должен поддерживать две команды:
Информация - по запросу юзера искать какую-то простую информацию в открытых источниках. Какую информацию - на твой выбор. Может быть погода в городе из запроса, текущее время в городе из запроса, курс валюты из запроса, и т.д.
Статистика - показать юзеру когда был его первый запрос, сколько всего запросов было. И можно добавить ещё показателей, которые мы можем получить из хранимых данных.

Бот должен быть упакован в Docker контейнер. Для удобства разработки и тестирования можно приложить docker-compose файл.

Обрати внимание на обработку ошибок и логирование.

# Quick Start

Для начала надо описать сервисы в docker-compose файле, необходимо создать в корне проекта docker-compose.yml.

``` 
  telegram:
    build: ./telegram-bot-service/
    command: ./telegram-bot-service/wait-for-postgres.sh db ./telegram-bot-service
    ports:
      - "8081:3000"
    volumes:
      - ./telegram-bot-service:/go/src/app/telegram-bot-service
      - ./telegram-bot-service/vendor:/go/src
    working_dir: /go/src/app/telegram-bot-service
    depends_on:
      - "postgres"
    environment:
      - POSTGRES_USER=postgres
      - POSTGRES_HOST=postgres
      - POSTGRES_PORT=5432
      - POSTGRES_PASS=postgres
      - POSTGRES_DB=postgres
      - TELEGRAM_BOT_TOKEN=[BOT_TOKEN]

  postgres:
    image: postgres
    restart: always
    environment:
      - POSTGRES_PASSWORD=postgres
      - POSTGRES_USER=postgres
      - POSTGRES_DB=postgres
    volumes:
      - ./.pg/postgres/data:/var/lib/postgresql/data
    ports:
      - "5432:5432"
```

Затем выполнить команду ```docker-compose up --build telegram``` . После поднятия postgres, надо создать таблицы:
```
CREATE TABLE user_requests
(
    id           SERIAL PRIMARY KEY,
    chat_id      INTEGER      NOT NULL,
    request      VARCHAR(100) NOT NULL,
    request_date TIMESTAMP    NOT NULL DEFAULT NOW(),
    result_json  JSON         NOT NULL
);
```

# Переменные окружения

Для работы telegram-service, нужно привязать токен бота [BOT_TOKEN] в docker-compose.yml
