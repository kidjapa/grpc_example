package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "google.golang.org/grpc"
    "log"
    "net/http"
    "server_main_land/proto"
    "strconv"
)

type ResponseError struct {
    Error string `json:"error"`
}

type ResponseSuccess struct {
    Response string `json:"response"`
}

func main() {
    conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
    if err != nil {
        panic(err)
    }

    client := proto.NewAddServiceClient(conn)

    g := gin.Default()

    g.GET("/add/:a/:b", func(ctx *gin.Context) {
        a, err := strconv.ParseUint(ctx.Param("a"), 10, 64)
        if err != nil {
            ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter a"})
            return
        }
        b, err := strconv.ParseUint(ctx.Param("b"), 10, 64)
        if err != nil {
            ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter B"})
            return
        }

        req := proto.Request{A: int64(a), B: int64(b)}
        if response, err := client.Add(ctx, &req); err == nil {
            ctx.JSON(http.StatusOK, gin.H{"result": fmt.Sprint(response)})
        }else{
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    })

    g.GET("/multiply/:a/:b", func(ctx *gin.Context) {
        a, err := strconv.ParseUint(ctx.Param("a"), 10, 64)
        if err != nil {
            ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter a"})
            return
        }
        b, err := strconv.ParseUint(ctx.Param("b"), 10, 64)
        if err != nil {
            ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter B"})
            return
        }

        req := proto.Request{A: int64(a), B: int64(b)}
        if response, err := client.Multiply(ctx, &req); err == nil {
            ctx.JSON(http.StatusOK, gin.H{"result": fmt.Sprint(response)})
        }else{
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    })

    if err := g.Run(":8080"); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }

}

//type AddServiceRoutes struct {
//    client *proto.AddServiceClient
//}
//
//func (add *AddServiceRoutes) fill(e *echo.Echo, client *proto.AddServiceClient) error {
//    add.client = client
//    e.GET("/add/:a/:b", add.addRoute)
//    e.GET("/multiply/:a/:b", add.multiplyRoute)
//    return nil
//}
//
//func (add *AddServiceRoutes) addRoute(c echo.Context) error {
//    a, err := strconv.ParseUint(c.Param("a"),10, 64)
//    if err != nil {
//        return c.String(http.StatusInternalServerError, "Error: a not defined")
//    }
//
//    b, err := strconv.ParseUint(c.Param("a"),10, 64)
//    if err != nil {
//        return c.String(http.StatusInternalServerError, "Error: a not defined")
//    }
//
//    req := &proto.Request{A: int64(a), B: int64(b)}
//
//    if response, err := add.client.Add(c, req); err == nil  {
//        c.JSON()
//    }
//
//
//    return c.String(http.StatusOK, strconv.FormatUint(result, 10))
//}
//
//func (add *AddServiceRoutes) multiplyRoute(c echo.Context) error {
//    a, err := strconv.ParseUint(c.Param("a"),10, 64)
//    if err != nil {
//        return c.String(http.StatusInternalServerError, "Error: a not defined")
//    }
//
//    b, err := strconv.ParseUint(c.Param("a"),10, 64)
//    if err != nil {
//        return c.String(http.StatusInternalServerError, "Error: a not defined")
//    }
//    result := a * b
//    return c.String(http.StatusOK, strconv.FormatUint(result, 10))
//}

