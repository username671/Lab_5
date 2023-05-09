# Web server

.git/…
cmd/ - специальная папка по соглашениям go, где лежат компилируемые исходники программы на Go
 blog/
  main.go - основной файл
  handlers.go - файл с  описанием обработчиков запросов
static/
 css/ - папка для css стилей
  Styles.css
 img/ - папка для картинок
  mat-vogels.png
pages/ - папка для html страниц
 index.html - главная страница блога
 post.html - страница поста, сейчас это страница поста The Road Ahead
go.mod - мета описание проекта для Golang компилятора
README.md


go run cmd/blog/main.go cmd/blog/handlers.go