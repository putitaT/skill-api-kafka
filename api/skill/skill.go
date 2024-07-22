package skill

import (
	"encoding/json"
	"fmt"
	"net/http"

	// "strings"

	"github.com/gin-gonic/gin"
	"github.com/putitaT/skill-api-kafka/api/database"
	"github.com/putitaT/skill-api-kafka/api/util"
)

var db = database.ConnectDB()

func SkillApi(r *gin.Engine) {
	r.GET("/api/v1/skills/:key", getSkillByKey)
	r.GET("/api/v1/skills", getAllSkill)
	r.POST("/api/v1/skills", createSkill)
	r.PUT("/api/v1/skills/:key", updateSkillById)
	r.PATCH("/api/v1/skills/:key/actions/name", updateNameById)
	r.PATCH("/api/v1/skills/:key/actions/description", updateDescriptionById)
	r.PATCH("/api/v1/skills/:key/actions/logo", updateLogoById)
	r.PATCH("/api/v1/skills/:key/actions/tags", updateTagById)
	r.DELETE("/api/v1/skills/:key", deleteSkillById)
}

func getSkillByKey(ctx *gin.Context) {
	key := ctx.Param("key")
	sql := "SELECT key, name, description, logo, tags FROM skill where key=$1"
	row := db.QueryRow(sql, key)

	var skill util.SkillDB
	var res map[string]any
	err := row.Scan(&skill.Key, &skill.Name, &skill.Description, &skill.Logo, &skill.Tags)
	if err != nil {
		res = map[string]any{
			"status":  "error",
			"message": "Skill not found",
		}
		ctx.JSON(http.StatusNotFound, res)
	} else {
		res = map[string]any{
			"status": "success",
			"data":   util.Skill(skill),
		}
		ctx.JSON(http.StatusOK, res)
	}
}

func getAllSkill(ctx *gin.Context) {
	rows, err := db.Query("SELECT key, name, description, logo, tags FROM skill ORDER BY key")
	var res map[string]any
	if err != nil {
		res = map[string]any{
			"status":  "error",
			"message": "Can't get all skill",
		}
		ctx.JSON(http.StatusNotFound, res)
	} else {
		skills := []util.SkillData{}
		for rows.Next() {
			var skill = util.SkillDB{}
			err := rows.Scan(&skill.Key, &skill.Name, &skill.Description, &skill.Logo, &skill.Tags)
			if err != nil {
				fmt.Println("can't Scan row into variable", err)
			}
			skills = append(skills, util.Skill(skill))
		}
		ctx.JSON(http.StatusOK, skills)
	}
}

func createSkill(ctx *gin.Context) {
	var newSkill util.SkillData

	if err := ctx.ShouldBindJSON(&newSkill); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusInternalServerError, map[string]error{"message": err})
		return
	}

	jsonData, err := json.Marshal(newSkill)
	if err != nil {
		fmt.Println("Error marshaling map:", err)
		ctx.JSON(http.StatusInternalServerError, map[string]error{"message": err})
		return
	}
	if err := Producer(jsonData, "create"); err != nil {
		ctx.JSON(http.StatusOK, map[string]error{"message": err})
		return
	}
	ctx.JSON(http.StatusOK, map[string]string{"message": "Adding Skill Data"})

}

func updateSkillById(ctx *gin.Context) {
	key := ctx.Param("key")
	var newSkill util.SkillData

	if err := ctx.ShouldBindJSON(&newSkill); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusInternalServerError, map[string]error{"message": err})
		return
	}

	jsonData, err := json.Marshal(map[string]any{"data": newSkill, "key": key})
	if err != nil {
		fmt.Println("Error marshaling map:", err)
		ctx.JSON(http.StatusInternalServerError, map[string]error{"message": err})
		return
	}
	if err := Producer(jsonData, "update"); err != nil {
		ctx.JSON(http.StatusOK, map[string]error{"message": err})
		return
	}
	ctx.JSON(http.StatusOK, map[string]string{"message": "Updating Skill Data"})

}

func updateNameById(ctx *gin.Context) {
	key := ctx.Param("key")
	var newSkill util.SkillData

	if err := ctx.ShouldBindJSON(&newSkill); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusInternalServerError, map[string]error{"message": err})
		return
	}

	jsonData, err := json.Marshal(map[string]any{"data": newSkill, "key": key})
	if err != nil {
		fmt.Println("Error marshaling map:", err)
		ctx.JSON(http.StatusInternalServerError, map[string]error{"message": err})
		return
	}
	if err := Producer(jsonData, "updateName"); err != nil {
		ctx.JSON(http.StatusOK, map[string]error{"message": err})
		return
	}
	ctx.JSON(http.StatusOK, map[string]string{"message": "Updating Skill Name Data"})

}

func updateDescriptionById(ctx *gin.Context) {
	key := ctx.Param("key")
	var newSkill util.SkillData

	if err := ctx.ShouldBindJSON(&newSkill); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusInternalServerError, map[string]error{"message": err})
		return
	}

	jsonData, err := json.Marshal(map[string]any{"data": newSkill, "key": key})
	if err != nil {
		fmt.Println("Error marshaling map:", err)
		ctx.JSON(http.StatusInternalServerError, map[string]error{"message": err})
		return
	}
	if err := Producer(jsonData, "updateDesc"); err != nil {
		ctx.JSON(http.StatusOK, map[string]error{"message": err})
		return
	}
	ctx.JSON(http.StatusOK, map[string]string{"message": "Updating Skill Description Data"})
}

func updateLogoById(ctx *gin.Context) {
	key := ctx.Param("key")
	var newSkill util.SkillData

	if err := ctx.ShouldBindJSON(&newSkill); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusInternalServerError, map[string]error{"message": err})
		return
	}

	jsonData, err := json.Marshal(map[string]any{"data": newSkill, "key": key})
	if err != nil {
		fmt.Println("Error marshaling map:", err)
		ctx.JSON(http.StatusInternalServerError, map[string]error{"message": err})
		return
	}
	if err := Producer(jsonData, "updateLogo"); err != nil {
		ctx.JSON(http.StatusOK, map[string]error{"message": err})
		return
	}
	ctx.JSON(http.StatusOK, map[string]string{"message": "Updating Skill Logo Data"})
}

func updateTagById(ctx *gin.Context) {
	key := ctx.Param("key")
	var newSkill util.SkillData

	if err := ctx.ShouldBindJSON(&newSkill); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusInternalServerError, map[string]error{"message": err})
		return
	}

	jsonData, err := json.Marshal(map[string]any{"data": newSkill, "key": key})
	if err != nil {
		fmt.Println("Error marshaling map:", err)
		ctx.JSON(http.StatusInternalServerError, map[string]error{"message": err})
		return
	}
	if err := Producer(jsonData, "updateTags"); err != nil {
		ctx.JSON(http.StatusOK, map[string]error{"message": err})
		return
	}
	ctx.JSON(http.StatusOK, map[string]string{"message": "Updating Skill Tags Data"})
}

func deleteSkillById(ctx *gin.Context) {
	key := ctx.Param("key")

	jsonData, err := json.Marshal(map[string]any{"data": util.SkillData{}, "key": key})
	if err != nil {
		fmt.Println("Error marshaling map:", err)
		ctx.JSON(http.StatusInternalServerError, map[string]error{"message": err})
		return
	}
	if err := Producer(jsonData, "delete"); err != nil {
		ctx.JSON(http.StatusOK, map[string]error{"message": err})
		return
	}
	ctx.JSON(http.StatusOK, map[string]string{"message": "Deleting Skill Data"})

}
