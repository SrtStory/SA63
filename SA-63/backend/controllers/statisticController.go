package controllers
 
import (
	"context" 
   	"strconv"
   
   "github.com/SrtStory/app/ent"
   "github.com/SrtStory/app/ent/employee" 
   "github.com/SrtStory/app/ent/contagious"
   "github.com/SrtStory/app/ent/area"
   "github.com/SrtStory/app/ent/carrier" 
   "github.com/SrtStory/app/ent/patient"
       
   "github.com/gin-gonic/gin"
)
 
// StatisticController defines the struct for the statistic controller
type StatisticController struct {
   client *ent.Client
   router gin.IRouter
}

// Statistic  defines the struct for the statistic controller
type Statistic struct {
	Employee   	int
	Contagious  int
	Patient   	int
	Carrier   	int
	Area   		int
}
// CreateStatistic handles POST requests for adding statistic entities
// @Summary Create statistic
// @Description Create statistic
// @ID create-statistic
// @Accept   json
// @Produce  json
// @Param Statistic body ent.Statistic true "Statistic entity"
// @Success 200 {object} ent.Statistic
// @Failure 400 {object} gin.H.
// @Failure 500 {object} gin.H
// @Router /statistic [post]
func (ctl *StatisticController) CreateStatistic(c *gin.Context) {
	obj := Statistic{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "statistic binding failed",
		})
		return
	}
	emp, err := ctl.client.Employee.
		Query().
		Where(employee.IDEQ(int(obj.Employee))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "employee not found",
		})
		return
	}

	ct, err := ctl.client.Contagious.
		Query().
		Where(contagious.IDEQ(int(obj.Contagious))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Contagious not found",
		})
		return
	}

	p, err := ctl.client.Patient.
		Query().
		Where(patient.IDEQ(int(obj.Patient))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "patient not found",
		})
		return
	}

	cr, err := ctl.client.Carrier.
		Query().
		Where(carrier.IDEQ(int(obj.Carrier))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "carrier not found",
		})
		return
	}

	a, err := ctl.client.Area.
		Query().
		Where(area.IDEQ(int(obj.Area))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "area not found",
		})
		return
	}
	

	s, err := ctl.client.Statistic.
		Create().
		SetEmployee(emp).	
		SetContagious(ct).
		SetPatient(p).
		SetCarrier(cr).
		SetArea(a).

		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, s)
}

// ListStatistic handles request to get a list of statistic entities
// @Summary List statistic entities
// @Description list statistic entities
// @ID list-statistic
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Statistic
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /statistics [get]
func (ctl *StatisticController) ListStatistic(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 100
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	statistics, err := ctl.client.Statistic.
		Query().
		WithEmployee().
		WithContagious().
		WithPatient().
		WithCarrier().
		WithArea().
	
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, statistics)
}

// NewStatisticController creates and registers handles for the statistic controller
func NewStatisticController(router gin.IRouter, client *ent.Client) *StatisticController {
	sc := &StatisticController{
		client: client,
		router: router,
	}
	sc.register()
	return sc
}

func (ctl *StatisticController) register() {
	statistics := ctl.router.Group("/statistics")
	statistics.POST("", ctl.CreateStatistic)
	statistics.GET("", ctl.ListStatistic)

}