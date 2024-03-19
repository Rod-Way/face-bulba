# FaceBulba

## Хэндлеры - api

POST /api/add_user
curl

POST /api/sign_in
пример curl запроса:
curl -X POST -H "Content-Type: application/json" -d '{"user":"johndoe","password":"password"}' http://localhost:5000/api/login
Пример ответа:
{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJnby1zb2NpYWwuand0Z28uaW8iLCJleHAiOjE3MTA3ODUyMDgsImlzcyI6Imp3dGdvLmlvIiwidXNlciI6ImpvaG5kb2UifQ.DxBs6dWZHFu1rTvpBBchyTCDB5UzieVPBOX3TMOPYSQ"}

GET /api/auth/:token
Пример curl запроса:
curl http://localhost:5000/api/auth/eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJnby1zb2NpYWwuand0Z28uaW8iLCJleHAiOjE3MTA3NjY0OTMsImlzcyI6Imp3dGdvLmlvIiwidXNlciI6IiJ9.RbH6YqS2qYl5KzEtvao5Zlk6gfhzoyvg7q2RX0U-BT8

Пример ответа:
{"isAuthenticated":true}

POST /api/post/create
Пример curl запроса:
curl -X POST http://localhost:5000/api/posts/create -H "Content-Type: application/json" -H "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJnby1zb2NpYWwuand0Z28uaW8iLCJleHAiOjE3MTA3ODUyMDgsImlzcyI6Imp3dGdvLmlvIiwidXNlciI6ImpvaG5kb2UifQ.DxBs6dWZHFu1rTvpBBchyTCDB5UzieVPBOX3TMOPYSQ" -d '{"content":"Hello, World!", "tags":["code"]}'
