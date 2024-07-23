package skill

import "github.com/gin-gonic/gin"

func Router(r *gin.Engine) {
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
