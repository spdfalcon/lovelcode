package handlers

import (
	// "errors"
	// "fmt"
	// "os"
	// "strings"
	"time"

	"github.com/gofiber/fiber/v2"
	// "github.com/google/uuid"
	"gorm.io/gorm"

	"lovelcode/database"
	"lovelcode/models"
	"lovelcode/utils"
)

/////////////  public   //////////////////////////

// POST
func CreateContactUs(c *fiber.Ctx) error{

	var mb models.IContactUs
	if err:= c.BodyParser(&mb); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid json"})
	}

	// check check validation
	if err:=mb.Check(); err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":err.Error()})
	}
	
	// create ContactUs and fill it
	var ContactUs models.ContactUs
	ContactUs.Fill(&mb)
	ContactUs.TimeCreated = time.Now()

	if err:= database.DB.Create(&ContactUs).Error; err!=nil{
		return utils.ServerError(c, err)
	}
	return utils.JSONResponse(c, 201, fiber.Map{"msg":"successfully created"})
}

///////////////  admin  //////////////////////

// GET, admin , /?page=1&pageLimit=20
func GetAllUnseenContactUss(c *fiber.Ctx) error{

	page, pageLimit, err :=utils.GetPageAndPageLimitFromMap(c.Queries())
	if err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid query:"+err.Error()})
	}

	var ContactUss []models.OContactUs
	if err:= database.DB.Model(&models.ContactUs{}).Where(&models.ContactUs{IsSeen: true}).Offset((page-1)*pageLimit).Limit(pageLimit).Find(&ContactUss).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"no ContactUs found"})
		}
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"data":ContactUss})
}


// GET, admin , /?page=1&pageLimit=20
func GetAllContactUss(c *fiber.Ctx) error{

	page, pageLimit, err :=utils.GetPageAndPageLimitFromMap(c.Queries())
	if err!=nil{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid query:"+err.Error()})
	}

	var ContactUss []models.OContactUs
	if err:= database.DB.Model(&models.ContactUs{}).Offset((page-1)*pageLimit).Limit(pageLimit).Find(&ContactUss).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"no ContactUs found"})
		}
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"data":ContactUss})
}

// GET, admin, /:contactusId
func GetContactUs(c *fiber.Ctx) error{
	
	id := utils.GetIDFromParams(c, "contactusId")
	if id == 0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid id"})
	}
	
	var ContactUs models.ContactUs
	if err:= database.DB.First(&ContactUs, &models.ContactUs{ID: id}).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"ContactUs not found"})
		}
		return utils.ServerError(c, err)
	}

	return utils.JSONResponse(c, 200, fiber.Map{"data":ContactUs})
	
}

// DELETE, admin /:title
func DeleteContactUs(c *fiber.Ctx) error{
	
	id := utils.GetIDFromParams(c, "contactusId")
	if id == 0{
		return utils.JSONResponse(c, 400, fiber.Map{"error":"invalid id"})
	}

	if err:= database.DB.Delete(&models.ContactUs{}, id).Error; err!=nil{
		if err==gorm.ErrRecordNotFound{
			return utils.JSONResponse(c, 404, fiber.Map{"error":"ContactUs not found"})
		}
		return utils.ServerError(c, err)
	}
	
	return utils.JSONResponse(c, 200, fiber.Map{"msg":"successfuly deleted"})

}