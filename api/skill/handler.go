package skill

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/putitaT/skill-api-kafka/api/database"
	"github.com/putitaT/skill-api-kafka/api/util"
)

var db = database.ConnectDB()

type skillHandler struct {
	repository skillRepository
}

func NewHandler(repository skillRepository) skillHandler {
	return skillHandler{repository: repository}
}

func (handler *skillHandler) GetSkillByKeyHandler(ctx *gin.Context) {
	key := ctx.Param("key")
	row := handler.repository.getSkillByKey(key)

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

func (handler *skillHandler) GetAllSkillHandler(ctx *gin.Context) {
	rows, err := handler.repository.getAllSkill()
	var res map[string]any
	if err != nil {
		res = map[string]any{
			"status":  "error",
			"message": err,
		}
		fmt.Printf("Error: %v\n", err)
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
