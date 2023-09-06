package user_list

import (
	"github.com/gin-gonic/gin"
	"golang_rest_api/internal/util/db"
)

func Handle(ctx *gin.Context) {
	type Result struct {
		Id int64 `json:"id"`
		Name string `json:"name"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
		DeletedAt *string `json:"deleted_at"`
	}

	var result Result

	dbConnection, _ := db.GetInstance()

	rows, _ := dbConnection.Raw("SELECT * FROM users WHERE deleted_at is NULL").Rows()

	data := make([]Result, 0, 50)

	defer rows.Close()

	for i:=0; rows.Next(); i++ {
		dbConnection.ScanRows(rows, &result)
		data = append(data, result)
	}

	ctx.JSON(200, data)
}
