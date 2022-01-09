# restless-program

this is a version of https://github.com/squeezesky/project-gorm-order program. But I have placed it separately because the methods in main classes are completely 
different it these programs and also I like this version more,than the first one.


 Data-base directory is the main program for the server's part in the restless service,
it takes request from a client and gives a response depends of a request method and context.
Warehouse directory is a secondary program that tells data-base program what amount of product
the warehouse has
Also it's a restless program, therefore it'll answer in JSON format. But I'm not sure yet about requests, for example if a client wants to delete an order probably
he should give a JSON formatted data to delete something, but perhaps this operation is too simple to do it, and it a lot easier to just get an id to a server
1. if the method is GET and a client didn't give any id then it'll show every order from a server's database (curl -v localhost:8080 -XGET)
2. if the method is GET and a client gave an id then it'll show detailed information about this order (curl -v localhost:8080/1 -XGET)
3. if the method is PUT //TODO
4. if the method is DELETE client should give the id of an order and the order will be deleted (curl -v localhost:8080/1 -XDELETE)


