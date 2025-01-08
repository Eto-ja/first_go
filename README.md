## Калькулятор для вычислений "+-*/()" используя односимвольные числа.

Чтобы запустить сервер используйте команду
```bash
go run cmd/main.go
```

Для проверки можно использовать вот такой запрос в терминале (программа сработает успешно):
```bash
$headers = @{
    "Content-Type" = "application/json"
}
$body = @{
    "expression" = "2+2"
} | ConvertTo-Json
 
 Invoke-WebRequest -Uri http://localhost:8181/api/v1/calculate -Method Post -Headers $headers -Body $body
```

В поле expression заводите своё арифметическое выражение.

Если хотите получить ошибку "Expression is not valid", то необходимо ввести такой запрос:
```bash
$headers = @{
    "Content-Type" = "application/json"
}
$body = @{
    "expression" = "2+"
} | ConvertTo-Json
 
 Invoke-WebRequest -Uri http://localhost:8181/api/v1/calculate -Method Post -Headers $headers -Body $body
```

 Если хотите получить ошибку "Internal server error", то необходимо ввести такой запрос:
```bash
$headers = @{
    "Content-Type" = "application"
}
$body = @{
    "expression" = "2+2"
} | ConvertTo-Json
 
 Invoke-WebRequest -Uri http://localhost:8181/api/v1/calculate -Method Post -Headers $headers -Body $body
```



 На этом всё, приятного пользования
