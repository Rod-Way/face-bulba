Целевая система:
    Windows 10 22H2

Процесс настройки среды выполнения и запуск:

    - Установить последнюю версию браузера Chrome

    - Установить данный проект (скорее всего если вы это читаете, то он уже установлен)

    - Установить Node.js и npm по ссылке (
            Если будут сложности с установкой - откройте https://nodejs.org/en/learn/getting-started/how-to-install-nodejs
        ):
        https://nodejs.org/en/download

    - Установить MongoDB:

        Перейдите по ссылке и установите приложение (
            Если будут сложности с установкой - откройте https://www.mongodb.com/docs/manual/administration/install-community/
        ):
            https://fastdl.mongodb.org/windows/mongodb-windows-x86_64-7.0.7-signed.msi

        В MongoDB установить соединение с mongodb://localhost:27017 (туториал по MongoDB https://youtu.be/pmjHPOPwX2A?si=yiNuT8pHedDaxWXf)

    - Запустить backend:

        - Открыть терминал (далее команды будут описываться для PowerShell)
        - Перейти в директорию проекта по умолчанию ( face-bulba )
        - Ввести команду:
            ./faceBulba.exe

    - Запустить frontend:

         - Открыть терминал (далее команды будут описываться для PowerShell)
        - Перейти в директорию проекта face-bulba/frontend
        - Ввести команду:
            npm start

    Далее откроется страница в браузере. Это и есть сайт.
