
# FaceBulba

  


Запросы отправляемые frontend сервером к backend серверу. В качестве наглядной демонстрации и для возможности самому протестировать приведены curl запросы
## Хендлеры 🪝

	Хочется уточнить, что зачастую ответом при правильном вводе является JSON конструкция вида {"message":"TEXT HERE"}. При ошибке ответом является конструкция вида {"error":"ERROR MESSAGE HERE"}. В ситуациях являющихся исключениям будет приведен пример успешного ответа.
  

- POST /api/register

		Добавляет данные пользователя в базу данных

	- Пример curl запроса:

		curl -X POST http://localhost:5000/api/register \
		-H "Content-Type: application/json" \
		-d '{
			"name": "Oleg",
			"surname": "Sazanovich",
			"username": "sazan4ik",
			"email": "sazan@mail.ru",
			"password": "password"
		}'


- POST /api/login

		Возвращает JWT токен, необходимый для действий, связанных с аккаунтом пользователя.

	- пример curl запроса:

		curl -X POST http://localhost:5000/api/login \
		-H "Content-Type: application/json" \
		-d '{"user":"sazan4ik","password":"password"}'

	- Ответ:
				{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJnby1zb2NpYWwuand0Z28uaW8iLCJleHAiOjE3MTA5NDgxNTksImlzcyI6Imp3dGdvLmlvIiwidXNlciI6InNhemFuNGlrIn0.2xB64gb7ImifgsnycURDVf9fI_g2NJ4jdzPIqt7ktNA"}
  
- GET /api/is-auth/:token

		Проверяет зарегестрирован ли тот или иной пользователь

	- Пример curl запроса:
	
		curl -X GET \
		http://localhost:5000/api/is-auth/eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJnby1zb2NpYWwuand0Z28uaW8iLCJleHAiOjE3MTA5NDgxNTksImlzcyI6Imp3dGdvLmlvIiwidXNlciI6InNhemFuNGlrIn0.2xB64gb7ImifgsnycURDVf9fI_g2NJ4jdzPIqt7ktNA

	- Ответ:
		  {"isAuthenticated":true}



- POST /api/get-batch/:batchNumber

		Возвращает страницу постов

	-  Пример curl запроса:
			curl -X GET http://localhost:5000/api/get-batch/1
			
	- Ответ:
		{"response":[{"text":"Hello World","files_url":null,"tags":["code"]},{"text":"Hello World","files_url":null,"tags":["code"]},{"text":"Hello World","files_url":null,"tags":["code"]}]}

- POST /api/get-by-id/:postID

		Возвращает конкретный пост

	-  Пример curl запроса:
			curl -X GET http://localhost:5000/api/get-by-id/65f9937351a6a04a00407ba0
			
	- Ответ:
		  {"response":{POST DATA HERE}}
