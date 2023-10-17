# About

Go-To-Do-App is simple mobile and desktop app written on Go that work with [Go-To-Do-Api](https://github.com/GAMch1k/go-to-do-api), it can show, edit, delete and post tasks on server.

It was written just to learn more about Go.


# Development hints

Build for android:

    fyne package -os android -appID org.gamch1k.todo_app -icon img/Icon.png

Install builded .apk using adb

    adb install todo_app.apk
