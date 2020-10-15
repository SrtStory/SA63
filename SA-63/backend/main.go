package main

import (
	"context"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"

	"github.com/SrtStory/app/controllers"
	"github.com/SrtStory/app/ent"
)

// Employees  defines the struct for the employees
type Employees struct {
	Employee []Employee
}

// Employee  defines the struct for the employee
type Employee struct {
	Name   string
	Email  string
	UserID string
}

// Contagiouss  defines the struct for the contagiouss
type Contagiouss struct {
	Contagious []Contagious
}

// Contagious  defines the struct for the  contagious
type Contagious struct {
	Name string
}

// Patients  defines the struct for the patients
type Patients struct {
	Patient []Patient
}

// Patient  defines the struct for the patient
type Patient struct {
	Name string
}

// Carriers  defines the struct for the carriers
type Carriers struct {
	Carrier []Carrier
}

// Carrier  defines the struct for the carrier
type Carrier struct {
	Name string
}

// Areas  defines the struct for the areas
type Areas struct {
	Area []Area
}

// Area  defines the struct for the area
type Area struct {
	Name string
}

// @title SUT SA Example API Statistic
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	client, err := ent.Open("sqlite3", "file:contagion.db?&cache=shared&_fk=1")
		if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewStatisticController(v1, client)
	controllers.NewEmployeeController(v1, client)
	controllers.NewContagiousController(v1, client)
	controllers.NewPatientController(v1, client)
	controllers.NewCarrierController(v1, client)
	controllers.NewAreaController(v1, client)

	

	// Set Employees Data
	employees := Employees{
		Employee: []Employee{
			Employee{"Dai Suki", "Daisuki@gmail.com", "emp01"},
			Employee{"Shin Taro", "Shintaro@gmail.com", "emp02"},
		},
	}

	for _, emp := range employees.Employee {
		client.Employee.
			Create().
			SetName(emp.Name).
			SetEmail(emp.Email).
			SetUserID(emp.UserID).
			Save(context.Background())
	}

	// Set Contagiouss Data
	contagiouss := Contagiouss{
		Contagious: []Contagious{
			Contagious{"โรคไข้เลือดออก"},
			Contagious{"โรคเอดส์"},
			Contagious{"โรคCovid-19"},
			Contagious{"โรคพิษสุนัขบ้า"},
		},
	}

	for _, ct := range contagiouss.Contagious {
		client.Contagious.
			Create().
			SetName(ct.Name).
			Save(context.Background())
	}

	// Set Patients Data
	patients := Patients{
		Patient: []Patient{
			Patient{"Male"},
			Patient{"Female"},			
		},
	}

	for _, p := range patients.Patient {
		client.Patient.
			Create().
			SetGender(p.Name).
			Save(context.Background())
	}

	// Set Carriers Data
	carriers := Carriers{
		Carrier: []Carrier{
			Carrier{"ยุงลาย"},
			Carrier{"มนุษย์"},
			Carrier{"สุนัข"},
			Carrier{"แมว"},
		},
	}

	for _, cr := range carriers.Carrier {
		client.Carrier.
			Create().
			SetName(cr.Name).
			Save(context.Background())
	}

	// Set Areas Data
	areas := Areas{
		Area: []Area{
			Area{"กรุงเทพมหานคร"},
			Area{"กระบี่"},
			Area{"กาญจนบุรี"},
			Area{"กาฬสินธุ์"},
			Area{"กำแพงเพชร"},
			Area{"ขอนแก่น"},
			Area{"จันทบุรี"},
			Area{"ฉะเชิงเทรา"},
			Area{"ชลบุรี"},
			Area{"ชัยนาท"},
			Area{"ชัยภูมิ"},
			Area{"ชุมพร"},
			Area{"เชียงราย"},
			Area{"เชียงใหม่"},
			Area{"ตรัง"},
			Area{"ตราด"},
			Area{"ตาก"},
			Area{"นครนายก"},
			Area{"นครปฐม"},
			Area{"นครพนม"},
			Area{"นครราชสีมา"},
			Area{"นครศรีธรรมราช"},
			Area{"นครสวรรค์"},
			Area{"นนทบุรี"},
			Area{"นราธิวาส"},
			Area{"น่าน"},
			Area{"บึงกาฬ"},
			Area{"บุรีรัมย์"},
			Area{"ปทุมธานี"},
			Area{"ประจวบคีรีขันธ์"},
			Area{"ปราจีนบุรี"},
			Area{"ปัตตานี"},
			Area{"พระนครศรีอยุธยา"},
			Area{"พังงา"},
			Area{"พัทลุง"},
			Area{"พิจิตร"},
			Area{"พิษณุโลก"},
			Area{"เพชรบุรี"},
			Area{"เพชรบูรณ์"},
			Area{"แพร่"},
			Area{"พะเยา"},
			Area{"ภูเก็ต"},
			Area{"มหาสารคาม"},
			Area{"มุกดาหาร"},
			Area{"แม่ฮ่องสอน"},
			Area{"ยะลา"},
			Area{"ยโสธร"},
			Area{"ร้อยเอ็ด"},
			Area{"ระนอง"},
			Area{"ระยอง"},
			Area{"ราชบุรี"},
			Area{"ลพบุรี"},
			Area{"ลำปาง"},
			Area{"ลำพูน"},
			Area{"เลย"},
			Area{"ศรีสะเกษ"},
			Area{"สกลนคร"},
			Area{"สงขลา"},
			Area{"สตูล"},
			Area{"สมุทรปราการ"},
			Area{"สมุทรสงคราม"},
			Area{"สมุทรสาคร"},
			Area{"สระแก้ว"},
			Area{"สระบุรี"},
			Area{"สิงห์บุรี"},
			Area{"สุโขทัย"},
			Area{"สุพรรณบุรี"},
			Area{"สุราษฎร์ธานี"},
			Area{"สุรินทร์"},
			Area{"หนองคาย"},
			Area{"หนองบัวลำภู"},
			Area{"อ่างทอง"},
			Area{"อุดรธานี"},
			Area{"อุทัยธานี"},
			Area{"อุตรดิตถ์"},
			Area{"อุบลราชธานี"},
			Area{"อำนาจเจริญ"},

		},
	}

	for _, a := range areas.Area {
		client.Area.
			Create().
			SetName(a.Name).
			Save(context.Background())
	}

	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}