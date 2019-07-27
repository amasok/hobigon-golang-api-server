package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/yyh-gl/hobigon-golang-api-server/domain/model"
	"github.com/yyh-gl/hobigon-golang-api-server/infra/gateway"
)

type response struct {
	Success  string
	TaskList []model.Task
}

func TaskHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// TODO: コンテキストの設定方法・場所のベストプラクティスが分かり次第修正
	ctx := r.Context()
	ctx = context.WithValue(ctx, "params", ps)

	taskGateway := gateway.NewtaskGateway()

	var todayTasks []model.Task
	boardIDList := [3]string{os.Getenv("MAIN_BOARD_ID"), os.Getenv("TECH_BOARD_ID"), os.Getenv("WORK_BOARD_ID")}
	for _, boardID := range boardIDList {
		lists, err := taskGateway.GetListsByBoardID(ctx, boardID)
		if err != nil {
			// TODO: ロガーに差し替え
			fmt.Println("v===== ERROR =====v")
			fmt.Println(err)
			fmt.Println("^===== ERROR =====^")
			return
		}

		for _, list := range lists {
			// TODO: 今後必要があれば動的に変更できる仕組みを追加
			if list.Name ==  "TODO" || list.Name == "WIP" {
				taskList, err := taskGateway.GetTasksFromList(ctx, *list)
				if err != nil {
					// TODO: ロガーに差し替え
					fmt.Println("v===== ERROR =====v")
					fmt.Println(err)
					fmt.Println("^===== ERROR =====^")
					return
				}

				switch list.Name {
				case "TODO":
					// TODOリストからは今日のタスクのみ出力
					tasks := taskList.GetTodayTasks()
					todayTasks = append(todayTasks, tasks...)
				case "WIP":
					// WIPリストにあるタスクは全て出力
					todayTasks = append(todayTasks, taskList.Tasks...)
				}
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response{ Success: "ok", TaskList: todayTasks}); err != nil {
		// TODO: エラーハンドリングをきちんとする
		http.Error(w, "Internal Server Error", 500)
		return
	}
}
