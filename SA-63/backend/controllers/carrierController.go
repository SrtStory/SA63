package controllers
 
import (
   "context"
   "strconv"

   "github.com/SrtStory/app/ent"
   "github.com/SrtStory/app/ent/carrier"
   "github.com/gin-gonic/gin"
)

// CarrierController defines the struct for the carrier controller
type CarrierController struct {
   client *ent.Client
   router gin.IRouter
}

// Carrier defines the struct for the carrier controller
type Carrier struct {
	Name     string
	
}


// CreateCarrier handles POST requests for adding carrier entities
// @Summary Create carrier
// @Description Create carrier
// @ID create-carrier
// @Accept   json
// @Produce  json
// @Param Carrier body ent.Carrier true "Carrier entity"
// @Success 200 {object} ent.Carrier
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /carriers [post]
func (ctl *CarrierController) CreateCarrier(c *gin.Context) {
	obj := ent.Carrier{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "carrier binding failed",
		})
		return
	}

	cr, err := ctl.client.Carrier.
		Create(). 
		SetName(obj.Name).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, cr)
}

// GetCarrier handles GET requests to retrieve a carrier entity
// @Summary Get a carrier entity by ID
// @Description get carrier by ID
// @ID get-carrier
// @Produce  json
// @Param id path int true "Carrier ID"
// @Success 200 {object} ent.Carrier
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /carriers/{id} [get]
func (ctl *CarrierController) GetCarrier(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	cr, err := ctl.client.Carrier.
		Query().
		Where(carrier.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, cr)
}

// ListCarrier handles request to get a list of carrier entities
// @Summary List carrier entities
// @Description list carrier entities
// @ID list-carrier
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Carrier
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /carriers [get]
func (ctl *CarrierController) ListCarrier(c *gin.Context) {
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

	carriers, err := ctl.client.Carrier.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, carriers)
}

// NewCarrierController creates and registers handles for the Carrier controller
func NewCarrierController(router gin.IRouter, client *ent.Client) *CarrierController {
	cc := &CarrierController{
		client: client,
		router: router,
	}
	cc.register()
	return cc
}

func (ctl *CarrierController) register() {
	carriers := ctl.router.Group("/carriers")

	carriers.GET("", ctl.ListCarrier)
	carriers.POST("", ctl.CreateCarrier)
	carriers.GET(":id", ctl.GetCarrier)
}