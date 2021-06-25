# simpleServer

Запуск на macos

* `brew install mongodb-community`
* `brew services start mongodb-community`
*  `export PORT=27017`
*  `go run -mod=mod main.go`


Методы:

* AddProdcut : 
curl -X POST \                                  
-d '{ 
    "title": "fist",
    "image": ["google.com/iamge.jpg"],          
    "price": 3,
    "Description": "Hello world"
    }' \
http://localhost:27017/AddProduct

* GetPage 
curl -X GET http://localhost:27017/Pages/?page=0&s=1
Где page - номер страницы, s - тип сортировки
s = 1 - по цене, возрастание
s = 2 - по цене, убывание
s = 3 - по дате добавления, возрастание
s = 4 - по дате добавления, убывания

* GetAll
curl -X GET http://localhost:27017
Все обьявления на одной странице

curl -X GET http://localhost:27017/products/id
Где id - номер обьявления
