Целевая система:
    Windows 10 22H2

Процесс настройки среды выполнения и запуск:


    - Установить данный проект (скорее всего если вы это читаете, то он уже установлен)

    - Установить последнюю версию браузера Chrome
        https://www.google.com/intl/ru_ru/chrome/

    - Установить Node.js и npm по ссылке (
            Если будут сложности с установкой - откройте https://nodejs.org/en/learn/getting-started/how-to-install-nodejs
        ):
        https://nodejs.org/en/download

    - Установить MongoDB MongoDB Compass:

        Перейдите по ссылке и установите приложение (при установке установить галочку на MongoDB Compass) (
            Если будут сложности с установкой - откройте https://www.mongodb.com/docs/manual/administration/install-community/
        ):
            https://fastdl.mongodb.org/windows/mongodb-windows-x86_64-7.0.7-signed.msi

        В MongoDB установить соединение с mongodb://localhost:27017 (туториал по MongoDB https://youtu.be/pmjHPOPwX2A?si=yiNuT8pHedDaxWXf)

    - Установить Git bash:
        https://git-scm.com/download/win
    ( Туториал по работе с Git bash: https://youtu.be/guok2Jj-TAM?si=HFMFjvYHJDKLcUac )



    Перед запуском следует убедится, что не запущеный программы, которые используют локальные порты 3000 и 5000.
    Проверить можно перейдя по ссылкам:
        http://localhost:3000/
        http://localhost:5000/
    Если есть сообщение "Не удается получить доступ к сайту" - все в порядке и можно переходить к следующуму этапу.
    В ином случае следует освободить данные порты.



    - Запустить backend:

        - Открыть Git bash
        - Перейти в корневую директорию проекта ( face-bulba )
        - Ввести команду:
            ./faceBulba.exe

    - Запустить frontend:

        - Открыть Git bash
        - Перейти в директорию проекта face-bulba/frontend
        - Ввести команду:
            npm start
