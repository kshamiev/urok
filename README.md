# urok

urok

### go.mod

добавление replace в go.mod

    go mod edit -replace github.com/user_name/repo_name=local_path

удаление replace из go.mod

    go mod edit -dropreplace github.com/user_name/repo_name
