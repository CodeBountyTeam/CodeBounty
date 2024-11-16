package handlers

import (
	"codebounty/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddProject(c *gin.Context) {
	/* Add a new project
	Request:  { id uint, title string, description string, link string, tags []string}
	Response: { message, error } */

	id, ok := getIdFromRequest(c)
	if !ok {
		return
	}

	project, ok := getProjectData(c)
	if !ok {
		return
	}

	if err := models.AddProject(id, project); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not add project",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Project created", "error": ""})
}

func UpdateProject(c *gin.Context) {

	/* Update a project
	Request: { id uint, title? string, description? string, link? string, tags? []string}
	Response: { message, error } */

	userid, ok := getIdFromRequest(c)
	if !ok {
		return
	}

	project, ok := getProjectData(c)
	if !ok {
		return
	}

	if err := models.UpdateProject(userid, project); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not update profile",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Project updated", "error": ""})
}

func GetProjectById(c *gin.Context) {
	/* Get a project by the project id
	Param: id uint
	Response: { project Project } */

	id, ok := getIdFromParam(c)
	if !ok {
		return
	}

	project, err := models.GetProjectById(id)
	if err != nil {
		errorMessage := fmt.Sprintf("Error getting project by id: %s: ", err)
		c.JSON(http.StatusNotFound, gin.H{"error": errorMessage})
		return
	}

	c.JSON(http.StatusOK, gin.H{"project": project, "error": ""})
}

func GetProjectsByUserId(c *gin.Context) {
	/* Get a projects by a user id
	Param: id uint
	Response: { projects []Project } */

	id, ok := getIdFromParam(c)
	if !ok {
		return
	}

	projects, err := models.GetProjectsByUserId(id)
	if err != nil {
		errorMessage := fmt.Sprintf("Error getting projects by user id: %s ", err)
		c.JSON(http.StatusNotFound, gin.H{"error": errorMessage})
		return
	}

	c.JSON(http.StatusOK, gin.H{"projects": projects, "error": ""})

}

/* --- utils --- */

func getProjectData(c *gin.Context) (models.Project, bool) {
	/* Fill a project struct with project data from the request */
	var project models.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid data",
		})
		return project, false
	}
	return project, true
}