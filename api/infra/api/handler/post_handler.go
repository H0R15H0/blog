package handler

import (
	"api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostHandler interface {
	CreatePost(c *gin.Context)
	GetPosts(c *gin.Context)
	GetPost(c *gin.Context)
	// UpdatePost(c *gin.Context)
	DeletePost(c *gin.Context)
}

type postHandler struct {
    postUsecase usecase.PostUsecase
}

// NewPostHandler post handlerのコンストラクタ
func NewPostHandler(postUsecase usecase.PostUsecase) PostHandler {
    return &postHandler{postUsecase: postUsecase}
}

func (pH *postHandler) GetPost(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	e, err := pH.postUsecase.FindByID(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, e)
}

func (pH *postHandler) GetPosts(c *gin.Context) {

	e, err := pH.postUsecase.FindAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, e)
}

func (pH *postHandler) CreatePost(c *gin.Context) {
	e := usecase.PostParamsObject{}

	if err := c.Bind(&e); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	post, err := pH.postUsecase.Create(e)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, post)
}


func (pH *postHandler) DeletePost(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := pH.postUsecase.Delete(id); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "ok"})
}
