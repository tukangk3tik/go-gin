# Go Gin-Gonic example

Go application that implements Gin Gonic http web framework 

First, download the library:

```sh
go get 
```

To run this application, build and run the Go binary:

```sh
go build
./go-jwt
```

Now, using any HTTP client (like [Postman](https://www.getpostman.com/apps)) hit the end point to get all students:

```
GET http://localhost:2345/api/students
{
  "data": [
    {
      "student_id": "123456",
      "full_name": "Budi Santoso",
      "address": "Jl. By Pass No. 10",
      "gender": "Laki-laki"
    },
    {
      "student_id": "87234",
      "full_name": "Nadia Smith",
      "address": "Jl. Madio Santoso No. 10",
      "gender": "Perempuan"
    },
    {
      "student_id": "44678",
      "full_name": "Iwan Dani",
      "address": "Jl. Karya No. 6",
      "gender": "Laki-laki"
    }
  ]
}
```

Then test hit another end point to get 1 student data:

```
GET http://localhost:2345/api/students/123456
{
  "data": {
    "student_id": "123456",
    "full_name": "Budi Santoso",
    "address": "Jl. By Pass No. 10",
    "gender": "Laki-laki"
  }
}
```

In the end, thank you for visiting this repository.

