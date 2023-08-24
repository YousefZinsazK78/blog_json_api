# blog json api


in this blog api backend we want to work with restful api mysql database and using mysql driver in golang
<br />
and using fiber golang framework and useing middlewares and routers ...

packages : 
   fiber v2 : github.com/gofiber/fiber/v2
   

mysql tables:
1. post table
   1. title (string) 
   2. description (string)
   3. author (string)
   4. likes ([]int)
   5. comments ([]string)
   6. createdAt (datetime)
   7. updatedAt (datetime)

2. user table
   1. fullname
   2. username
   3. password
   4. email
   5. (private) admin , guest ,writer user permission
      1. rule writer => write , read , update , delete blog
      2. rule admin => write , read , update , delete blog, delete comments, delete users
      3. rule guest => read blog , read comments
   