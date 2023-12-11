# go_learning

## Gorutines

### Channels
https://habr.com/ru/articles/490336/  

#### Unbuffered channels
- Чтение или запись данных в канал блокирует горутину и контроль передается свободной горутине

#### Buffered channels
- Когда размер буфера больше 0, горутина не блокируется до тех пор, пока буфер не будет заполнен. Когда буфер заполнен, любые значения отправляемые через канал, добавляются к буферу, отбрасывая предыдущее значение, которое доступно для чтения  
- Операция чтения на буферизированном канале является жадной, таким образом, как только операция чтения началась, она не будет завершена до полного опустошения буфера. Это означает, что горутина будет считывать буфер канала без блокировки до тех пор, пока буфер не станет пустым  

## RabbitMQ
Docker commands:  
```docker pull rabbitmq:3.9.29-management-alpine```   
```docker run -d -p 15672:15672 --hostname my-rabbit --name some-rabbit rabbitmq:3.9.29-management-alpine```   
http://localhost:15672/  
Default User/password: ```guest/guest``` 

## Project launching  
Inspect RabbitMQ container and get it ip address (```172.17.0.2:5672```)  
Inert it to the eventBus/reciever.go & eventBus/sender.go  
Run next commands in eventBus folder:    
```go run .``` or ```go build && ./eventBus```  
```curl localhost:8080``` - for new event through server  