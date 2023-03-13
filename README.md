# DevOps-Task2-GO-backend

#[backend]
#Adjust your backend application:
#when you execute (from your Terminal: Powershell or Bash) 
#curl -X GET http://localhost/api/hello -d'{"name":"Andrey"}'
#it should return 
#Hello, Andrey!


go run ./hellobackend.go
curl -X GET http://localhost/api/hello -d'{"name":"Andrey"}'