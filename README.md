# Ova-entertainment-api
<a href='https://github.com/jpoles1/gopherbadger' target='_blank'>![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Go%20Coverage-94%25-brightgreen.svg?longCache=true&style=flat)</a>

Запуска
`make run
`

Генерация моков
`mockgen -source=repo.go  -destination=mock_repo.go -package=repo
`

Создание миграции
`goose -s -dir migration create init sql
`

