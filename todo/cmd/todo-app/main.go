package main

import (
    "log"
    "net/http"
    "github.com/adsyandex/otus_shool/todo/internal/api"
    "github.com/adsyandex/otus_shool/todo/internal/storage"
    "github.com/adsyandex/otus_shool/todo/internal/task"
    "github.com/gorilla/mux"
)

func main() {
    store := &storage.FileStorage{}
    taskService := task.New(store)
    
    router := mux.NewRouter()
    api.RegisterRoutes(router, taskService)
    
    log.Println("Starting server on :8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}