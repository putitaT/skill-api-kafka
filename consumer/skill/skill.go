package skill

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	// "net/http"
	"os"

	// "github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"github.com/putitaT/skill-api-kafka/consumer/util"
)

var db = connectDB()

func connectDB() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Connect to database error", err)
	}

	fmt.Println("okay")

	return db
}

func CreateSkill(data []byte) error {
	var skill util.SkillDB

	newSkill := util.SkillData{}
	errT := json.Unmarshal([]byte(data), &newSkill)
	if errT != nil {
		fmt.Println("Can't convert data")
		return errT
	}

	sql := "INSERT INTO skill (key, name, description, logo, tags) values ($1, $2, $3, $4, $5) RETURNING key, name, description, logo, tags"
	row := db.QueryRow(sql, newSkill.Key, newSkill.Name, newSkill.Description, newSkill.Logo, pq.Array(newSkill.Tags))
	err := row.Scan(&skill.Key, &skill.Name, &skill.Description, &skill.Logo, &skill.Tags)

	// var res map[string]any
	if err != nil {
		// res = map[string]any{
		// 	"status":  "error",
		// 	"message": "Skill already exists",
		// }
		fmt.Println("Skill already exists: ", err)
		// ctx.JSON(http.StatusNotFound, res)
		return err
	} else {
		// res = map[string]any{
		// 	"status": "success",
		// 	"data":   util.Skill(skill),
		// }
		// ctx.JSON(http.StatusOK, res)
		fmt.Println("Create Skill Success")
		return nil
	}
}

func updateSkillById(data []byte) error {
	var skill util.SkillDB

	newSkill, key, errT := util.ConvertSkillData(data)
	if errT != nil {
		return errT
	}

	row := db.QueryRow("UPDATE skill SET key=$1, name=$2, description=$3, logo=$4, tags=$5 WHERE key=$1 RETURNING key, name, description, logo, tags;",
		key, newSkill.Name, newSkill.Description, newSkill.Logo, pq.Array(newSkill.Tags))
	err := row.Scan(&skill.Key, &skill.Name, &skill.Description, &skill.Logo, &skill.Tags)

	if err != nil {
		// res = map[string]any{
		// 	"status":  "error",
		// 	"message": "not be able to update skill",
		// }
		fmt.Println("not be able to update skill: ", err)
		return err
		// ctx.JSON(http.StatusNotFound, res)
	} else {
		// res = map[string]any{
		// 	"status": "success",
		// 	"data":   util.Skill(skill),
		// }
		// ctx.JSON(http.StatusOK, res)
		fmt.Println("Update Skill Success")
		return nil
	}
}

func updateNameById(data []byte) error {
	var skill util.SkillDB

	newSkill, key, errT := util.ConvertSkillData(data)
	if errT != nil {
		return errT
	}

	row := db.QueryRow("UPDATE skill SET name=$2 WHERE key=$1 RETURNING key, name, description, logo, tags;", key, newSkill.Name)
	err := row.Scan(&skill.Key, &skill.Name, &skill.Description, &skill.Logo, &skill.Tags)

	if err != nil {
		// res = map[string]any{
		// 	"status":  "error",
		// 	"message": "not be able to update skill name",
		// }
		fmt.Println("not be able to update skill name: ", err)
		// ctx.JSON(http.StatusNotFound, res)
		return err
	} else {
		// res = map[string]any{
		// 	"status": "success",
		// 	"data":   util.Skill(skill),
		// }
		// ctx.JSON(http.StatusOK, res)
		fmt.Println("Update Skill Name Success")
		return nil
	}
}

func updateDescriptionById(data []byte) error {
	var skill util.SkillDB

	newSkill, key, errT := util.ConvertSkillData(data)
	if errT != nil {
		return errT
	}
	row := db.QueryRow("UPDATE skill SET description=$2 WHERE key=$1 RETURNING key, name, description, logo, tags;", key, newSkill.Description)
	err := row.Scan(&skill.Key, &skill.Name, &skill.Description, &skill.Logo, &skill.Tags)

	if err != nil {
		// res = map[string]any{
		// 	"status":  "error",
		// 	"message": "not be able to update skill description",
		// }
		fmt.Println("not be able to update skill description: ", err)
		return err
		// ctx.JSON(http.StatusNotFound, res)
	} else {
		// res = map[string]any{
		// 	"status": "success",
		// 	"data":   util.Skill(skill),
		// }
		// ctx.JSON(http.StatusOK, res)
		fmt.Println("Update Skill Description Success")
		return nil
	}
}

func updateLogoById(data []byte) error {
	var skill util.SkillDB

	newSkill, key, errT := util.ConvertSkillData(data)
	if errT != nil {
		return errT
	}

	row := db.QueryRow("UPDATE skill SET logo=$2 WHERE key=$1 RETURNING key, name, description, logo, tags;", key, newSkill.Logo)
	err := row.Scan(&skill.Key, &skill.Name, &skill.Description, &skill.Logo, &skill.Tags)

	if err != nil {
		// res = map[string]any{
		// 	"status":  "error",
		// 	"message": "not be able to update skill logo",
		// }
		fmt.Println("not be able to update skill logo: ", err)
		return err
		// ctx.JSON(http.StatusNotFound, res)
	} else {
		// res = map[string]any{
		// 	"status": "success",
		// 	"data":   util.Skill(skill),
		// }
		// ctx.JSON(http.StatusOK, res)
		fmt.Println("Update Skill Logo Success")
		return nil
	}
}

func updateTagById(data []byte) error {
	var skill util.SkillDB

	newSkill, key, errT := util.ConvertSkillData(data)
	if errT != nil {
		return errT
	}

	row := db.QueryRow("UPDATE skill SET tags=$2 WHERE key=$1 RETURNING key, name, description, logo, tags;", key, pq.Array(newSkill.Tags))
	err := row.Scan(&skill.Key, &skill.Name, &skill.Description, &skill.Logo, &skill.Tags)

	if err != nil {
		// res = map[string]any{
		// 	"status":  "error",
		// 	"message": "not be able to update skill tags",
		// }
		fmt.Println("not be able to update skill tags: ", err)
		return err
		// ctx.JSON(http.StatusNotFound, res)
	} else {
		// res = map[string]any{
		// 	"status": "success",
		// 	"data":   util.Skill(skill),
		// }
		// ctx.JSON(http.StatusOK, res)
		fmt.Println("Update Skill Tags Success")
		return nil
	}
}

func deleteSkillById(data []byte) error {
	_, key, errT := util.ConvertSkillData(data)
	if errT != nil {
		return errT
	}

	allSkill := []util.SkillData{}
	rows, err := db.Query("SELECT key, name, description, logo, tags FROM skill ORDER BY key")
	if err != nil {
		fmt.Println("can't get all skill", err)
	} else {
		for rows.Next() {
			var skill = util.SkillDB{}
			err := rows.Scan(&skill.Key, &skill.Name, &skill.Description, &skill.Logo, &skill.Tags)
			if err != nil {
				fmt.Println("can't Scan row into variable", err)
			}
			allSkill = append(allSkill, util.Skill(skill))
		}
	}

	if checkExistKey(allSkill, key) {
		if _, err := db.Exec("DELETE FROM skill WHERE key=$1;", key); err != nil {
			// res = map[string]any{
			// 	"status":  "error",
			// 	"message": "not be able to delete skill",
			// }
			// ctx.JSON(http.StatusNotFound, res)
			return err
		} else {
			// res = map[string]any{
			// 	"status":  "success",
			// 	"message": "Skill deleted",
			// }
			// ctx.JSON(http.StatusOK, res)
			fmt.Println("Delete Skill Success")
			return nil
		}
	} else {
		// res = map[string]any{
		// 	"status":  "error",
		// 	"message": "Skill key invalid",
		// }
		// ctx.JSON(http.StatusNotFound, res)
		fmt.Println("Skill key invalid")
		return nil
	}
}

func checkExistKey(allSkill []util.SkillData, key string) bool {
	for _, v := range allSkill {
		if v.Key == key {
			return true
		}
	}
	return false
}
