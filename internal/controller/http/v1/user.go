package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github/go-clean-template/internal/entity"
	"github/go-clean-template/internal/usecase"
	"github/go-clean-template/pkg/logger"

	"google.golang.org/protobuf/encoding/protojson"
)

type userRoutes struct {
	t usecase.User
	l logger.Interface
}

func newUserRoutes(handler *gin.RouterGroup, t usecase.User, l logger.Interface) {
	r := &userRoutes{t, l}

	h := handler.Group("/users")
	{
		h.GET("/create", r.CreateUser)
		h.POST("/:id", r.GetUserById)
		h.PUT("/update/:id", r.UpdateUser)
		h.DELETE("/delete/:id", r.DeleteUser)
	}
}


// CreateUser
// @Router /users/create [post]
// @Summary create user
// @Tags User
// @Description Insert a new post with provided details
// @Accept json
// @Produce json
// @Param PostDetails body entity.User true "Create user"
// @Success 201 {object} entity.User
// @Failure 400 {object} response
// @Failure 500 {object} response
func (p *userRoutes) CreateUser(c *gin.Context) {
	var (
		body       entity.User
		jspMarshal protojson.MarshalOptions
	)
	jspMarshal.UseProtoNames = true

	err := c.BindJSON(&body)
	if err != nil {
		p.l.Error(err, "http - v1 - create post")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}

	body.Id = uuid.New().String()

	user, err := p.t.CreateUser(c.Request.Context(), &body)
	if err != nil {
		p.l.Error(err, "http - v1 - create user")
		errorResponse(c, http.StatusInternalServerError, "create user service problems")

		return
	}

	c.JSON(http.StatusOK, user)
}

// Update User
// @Router /users/update/{id} [put]
// @Summary update user
// @Tags User
// @Description Update user
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param UserInfo body entity.User true "Update User"
// @Success 201 {object} entity.User
// @Failure 400 {object} response
// @Failure 500 {object} response
func (p *userRoutes) UpdateUser(c *gin.Context) {
	var (
		body        entity.User
		jspbMarshal protojson.MarshalOptions
	)
	id := c.Param("id")

	jspbMarshal.UseProtoNames = true
	err := c.ShouldBindJSON(&body)
	if err != nil {
		p.l.Error(err, "http - v1 - update user")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}

	body.Id = id
	response, err := p.t.UpdateUser(c.Request.Context(), &body)
	if err != nil {
		p.l.Error(err, "http - v1 - update user")
		errorResponse(c, http.StatusInternalServerError, "update user service problems")

		return
	}

	c.JSON(http.StatusOK, response)
}


// Get User By Id
// @Router /users/{id} [get]
// @Summary get post by id
// @Tags User
// @Description Get user
// @Accept json
// @Produce json
// @Param id path string true "Id"
// @Success 201 {object} entity.User
// @Failure 400 {object} response
// @Failure 500 {object} response
func (p *userRoutes) GetUserById(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	post, err := p.t.GetUser(c.Request.Context(), id)
	if err != nil {
		p.l.Error(err, "http - v1 - get post")
		errorResponse(c, http.StatusInternalServerError, "get user service problems")

		return
	}

	c.JSON(http.StatusOK, post)
}

// Delete User
// @Router /user/delete/{id} [delete]
// @Summary delete user
// @Tags User
// @Description Delete user
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 201 {object} entity.MessageResponse
// @Failure 400 {object} response
// @Failure 500 {object} response
func (p *userRoutes) DeleteUser(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	err := p.t.DeleteUser(c.Request.Context(), id)
	if err != nil {
		p.l.Error(err, "http - v1 - delete user")
		errorResponse(c, http.StatusInternalServerError, "delete user service problems")

		return
	}

	c.JSON(http.StatusOK, entity.MessageResponse{
		Message: "user was successfully deleted",
	})
}
