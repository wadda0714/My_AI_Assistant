package handler

import(
        "github.com/wadda0714/My_AI_Assistant/usecase"
        "github.com/wadda0714/My_AI_Assistant/domain/model"
        "github.com/wadda0714/My_AI_Assistant/domain/query"

)

func CmdHandler(cmd string) error{

        switch cmd {
        case "start":
                return usecase.StartConversation(model.StartConversation{

}
