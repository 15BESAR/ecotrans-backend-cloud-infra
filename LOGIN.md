## This document consists of login spec

POST /login
    body
        {
            "username" : "<input username here>",
            "password" : "<input password here>"
        }
    message error
        301 : {"error" : "Username not found"}
        301 : {"error" : "Wrong password"}
        400 : {"error" : "bad request"}

        200 : {"userId" : <output user id here>, "token" : <jwt token>}
    
    *Token will be expired in 7 days, need to refresh

POST /refresh
    body
        {
            "token" : "<input jwt token here>"
        }
    message error
        401 : : not authorized

        200 : {"token" : <new jwt token>}
    *go to this path to refresh the token, by 15 minutes time mark

POST /register
    body
        {
            "username" : "<input username here>",
            "email" : "<input email here>",
            "password" : "<input password here>"
        }
For logout, just forget the token
