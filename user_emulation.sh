MAIN_PATH="http://localhost:5000/api"

get_token(){
    local register_data="$1"
    local login_data="$2"

    local register_url="$MAIN_PATH/register"
    local login_url="$MAIN_PATH/login"

    # Регистрация
    local register_res=$(curl -s -X POST "$register_url" -H "Content-Type: application/json" -d "$register_data")
    if [ $? -ne 0 ]; then
        echo "Ошибка при регистрации"
        return 1
    fi

    # Вход
    local login_res=$(curl -s -X POST "$login_url" -H "Content-Type: application/json" -d "$login_data")
    if [ $? -ne 0 ]; then
        echo "Ошибка при входе"
        return 1
    fi

    echo "$login_res"
}

create_post(){
    local token="$1"
    local post_data="$2"

    local post_url="$MAIN_PATH/posts/create"

    local res=$(curl -s -X POST "$post_url" -H "Content-Type: application/json" -H "Authorization: $token" -d "$post_data" --fail)
    if [ $? -ne 0 ]; then
        echo "Ошибка при создании поста"
        return 1
    fi
}

register_data='{"name":"Oleg", "surname":"Sazanovich", "username": "sazan4ik", "email":"sazanovi4@mail.ru", "password": "sazan4ik"}'
login_data='{"username": "sazan4ik", "password": "sazan4ik"}'

Oleg_TOKEN=$(get_token "$register_data" "$login_data")
if [ $? -eq 0 ]; then
    echo "TOKEN: $Oleg_TOKEN"
    post_data='{"text":"Сегодня был такой день, когда даже браузер стал мне задавать вопросы о своем смысле жизни. Я просто хотел протестировать кнопку Войти, а она начала философствовать. 🤔", "tags":["Тестирование", "Баги", "ФилософияКода"]}'
    create_post "$Oleg_TOKEN" "$post_data"
else
    echo "Не удалось получить токен"
fi

post_data='{"text":"Сегодня был такой день, когда даже браузер стал мне задавать вопросы о своем смысле жизни. Я просто хотел протестировать кнопку Войти, а она начала философствовать. 🤔", "tags":["Тестирование", "Баги", "ФилософияКода"]}'
create_post "$Oleg_TOKEN" "$post_data"

##################################################################################################################################

register_data='{"name":"Borislav", "surname":"Lapshinin", "username": "b0risBritva", "email":"BIGboris@mail.ru", "password": "b0risBritva"}'
login_data='{"username": "b0risBritva", "password": "b0risBritva"}'

Boris_TOKEN=$(get_token "$register_data" "$login_data")
if [ $? -eq 0 ]; then
    echo "TOKEN: $Boris_TOKEN"
else
    echo "Не удалось получить токен"

fi