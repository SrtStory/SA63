package controllers
	
	import (
	"context"
	"strconv"

	"github.com/SrtStory/app/ent"
	"github.com/SrtStory/app/ent/area"
	"github.com/gin-gonic/gin"
	)

// AreaController defines the struct for the area controller
	type AreaController struct {
	client *ent.Client
	router gin.IRouter
	}

	// Area defines the struct for the area controller
type Area struct {
	Name     string
	
}


// CreateArea handles POST requests for adding area entities
// @Summary Create area
// @Description Create area
// @ID create-area
// @Accept   json
// @Produce  json
// @Param area body ent.Area true "Area entity"
// @Success 200 {object} ent.Area
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /areas [post]
func (ctl *AreaController) CreateArea(c *gin.Context) {
	obj := ent.Area{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "area binding failed",
		})
		return
	}

	a, err := ctl.client.Area.
		Create(). 
		SetName(obj.Name).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, a)
}

// GetArea handles GET requests to retrieve a area entity
// @Summary Get a area entity by ID
// @Description get area by ID
// @ID get-area
// @Produce  json
// @Param id path int true "Area ID"
// @Success 200 {object} ent.Area
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /areas/{id} [get]
func (ctl *AreaController) GetArea(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	a, err := ctl.client.Area.
		Query().
		Where(area.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, a)
}

// ListArea handles request to get a list of area entities
// @Summary List area entities
// @Description list area entities
// @ID list-area
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Area
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /areas [get]
func (ctl *AreaController) ListArea(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	areas, err := ctl.client.Area.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, areas)
}

// NewAreaController creates and registers handles for the Area controller
func NewAreaController(router gin.IRouter, client *ent.Client) *AreaController {
	ac := &AreaController{
		client: client,
		router: router,
	}
	ac.register()
	return ac
}

func (ctl *AreaController) register() {
	areas := ctl.router.Group("/areas")

	areas.GET("", ctl.ListArea)
	areas.POST("", ctl.CreateArea)
	areas.GET(":id", ctl.GetArea)
}