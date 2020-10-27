Консольную утилиту - `crawler`, которая рекурсивно парсить страницы,
находить в них ссылки на другие (ранее не спаршенные)
страницы и тоже их парсить.

## Docker image
https://hub.docker.com/r/yakovbadygin/goparses


## Запуск in Docker
`docker run yakovbadygin/goparses:latest -r 2 -root https://meduza.io/ -n 10`

## Параметры для запуска

* `n` - максимальное количество параллельных запросов

* `root` - начальная страница

* `r` - глубина рекурсии

* `user-agent` - заголовок `User-Agent`

Например: `./crawler -r 10 -root https://meduza.io/ -n 15`
