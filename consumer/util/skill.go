package util

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/lib/pq"
)

type SkillDB struct {
	Key         string         `json:"key"`
	Name        sql.NullString `json:"name"`
	Description sql.NullString `json:"description"`
	Logo        sql.NullString `json:"logo"`
	Tags        pq.StringArray `json:"tags"`
}

type SkillData struct {
	Key         string
	Name        string
	Description string
	Logo        string
	Tags        []string
}

type Message struct {
	Key  string
	Data SkillData
}

func Skill(val SkillDB) SkillData {
	var skill SkillData
	skill.Key = val.Key
	skill.Description = val.Description.String
	skill.Logo = val.Logo.String
	skill.Name = val.Name.String
	skill.Tags = val.Tags
	return skill
}

func ConvertSkillData(data []byte) (SkillData, string, error) {
	result := Message{}
	errT := json.Unmarshal([]byte(data), &result)
	if errT != nil {
		fmt.Println("Can't convert data:", errT)
		return SkillData{}, "", errT
	}

	newSkill := result.Data
	key := result.Key

	return newSkill, key, nil
}
