package valueHandler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/core/result"
	redisDB "server/databases/redis"
	"server/models/value"
	"strconv"
)

type indexType struct {
	index int
}

func GetAll(c *gin.Context) {
	val, _ := value.GetAll()
	c.JSON(http.StatusOK, result.SuccessData("Values fetch success.", val))
}

func GetCurrent(c *gin.Context) {
	val, _ := redisDB.HGetAll("values")
	c.JSON(http.StatusOK, result.SuccessData("Current value fetch success.", val))
}

func Create(c *gin.Context) {
	var ind indexType
	redisDB.HSet("values", strconv.Itoa(ind.index), "Nothing yet!")
	redisDB.Publish("insert", strconv.Itoa(ind.index))
	value.Create(&value.Value{
		Values: ind.index,
	})
	c.JSON(http.StatusCreated, result.Success("Created process success."))
}
