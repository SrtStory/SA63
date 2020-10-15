package controllers
 
import (
   "context"   
   "strconv"

   "github.com/SrtStory/app/ent"
   "github.com/SrtStory/app/ent/contagious"
   "github.com/gin-gonic/gin"
)

// ContagiousController defines the struct for the contagious controller
type ContagiousController struct {
   client *ent.Client
   router gin.IRouter
}

// Contagious defines the struct for the contagious controller
type Contagious struct {
	Name     string
}


// CreateContagious handles POST requests for adding contagious entities
// @Summary Create contagious
// @Description Create contagious
// @ID create-contagious
// @Accept   json
// @Produce  json
// @Param contagious body ent.Contagious true "Contagious entity"
// @Success 200 {object} ent.Contagious
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /contagiouss [post]
func (ctl *ContagiousController) CreateContagious(c *gin.Context) {
	obj := ent.Contagious{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "contagious binding failed",
		})
		return
	}

	ct, err := ctl.client.Contagious.
		Create(). 
		SetName(obj.Name).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, ct)
}

// GetContagious handles GET requests to retrieve a contagious entity
// @Summary Get a contagious entity by ID
// @Description get contagious by ID
// @ID get-contagious
// @Produce  json
// @Param id path int true "Contagious ID"
// @Success 200 {object} ent.Contagious
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /contagiouss/{id} [get]
func (ctl *ContagiousController) GetContagious(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	ct, err := ctl.client.Contagious.
		Query().
		Where(contagious.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, ct)
}

// ListContagious handles request to get a list of contagious entities
// @Summary List contagious entities
// @Description list contagious entities
// @ID list-contagious
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Contagious
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /contagiouss [get]
func (ctl *ContagiousController) ListContagious(c *gin.Context) {
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

	contagiouss, err := ctl.client.Contagious.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, contagiouss)
}

// NewContagiousController creates and registers handles for the Contagious controller
func NewContagiousController(router gin.IRouter, client *ent.Client) *ContagiousController {
	ctc := &ContagiousController{
		client: client,
		router: router,
	}
	ctc.register()
	return ctc
}

func (ctl *ContagiousController) register() {
	contagiouss := ctl.router.Group("/contagiouss")

	contagiouss.GET("", ctl.ListContagious)
	contagiouss.POST("", ctl.CreateContagious)
	contagiouss.GET(":id", ctl.GetContagious)
}