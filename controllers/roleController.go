package controllers

import (
	"fmt"
	"go-admin/database"
	"go-admin/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func AllRoles(c *fiber.Ctx) error {
	var roles []models.Role

	database.DB.Find(&roles)

	return c.JSON(roles)
}

func CreateRole(c *fiber.Ctx) error {
	var roleDto fiber.Map

	if err := c.BodyParser(&roleDto); err != nil {
		return err
	}

	list := roleDto["permissions"].([]interface{})
	permissions := make([]models.Permission, len(list))

	for i, permissionId := range list {
		id, _ := strconv.Atoi(permissionId.(string))

		permissions[i] = models.Permission{
			Id: uint(id),
		}
	}

	role := models.Role{
		Name:        roleDto["name"].(string),
		Permissions: permissions,
	}

	database.DB.Create(&role)

	return c.JSON(role)
}

func GetRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role{
		Id: uint(id),
	}

	database.DB.Find(&role)

	return c.JSON(role)

}

type RoleDto struct {
	Permissions []uint `json:"permissions"`
	Name        string `json:"name"`
}

func UpdateRole(c *fiber.Ctx) error {
	roleId, _ := strconv.Atoi(c.Params("id"))

	var roleDto RoleDto

	if err := c.BodyParser(&roleDto); err != nil {
		return err
	}

	fmt.Printf("%+v", roleDto)
	list := roleDto.Permissions
	permissions := make([]models.Permission, len(list))

	for i, permissionId := range list {
		id := permissionId
		permissions[i] = models.Permission{
			Id: uint(id),
		}
	}

	database.DB.Table("role_permissions").Where("role_id = ?", roleId).Delete(nil)

	role := models.Role{
		Id:          uint(roleId),
		Name:        roleDto.Name,
		Permissions: permissions,
	}

	database.DB.Model(&role).Updates(role)

	return c.JSON(role)
}

func DeleteRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role{
		Id: uint(id),
	}
	if err := c.BodyParser(&role); err != nil {
		return err
	}

	database.DB.Model(&role).Delete(role)

	return nil
}
