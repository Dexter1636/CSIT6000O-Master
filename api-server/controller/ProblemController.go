package controller

import (
	"context"
	"errors"
	"log"
	"net/http"
	"ustoj-master/common"
	"ustoj-master/model"
	"ustoj-master/utils"
	"ustoj-master/vo"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IProblemController interface {
	ProblemList(c *gin.Context)
	ProblemDetail(c *gin.Context)
}

type ProblemController struct {
	DB  *gorm.DB
	Ctx context.Context
}

func NewProblemController() IProblemController { // Similar to the interface of service
	return ProblemController{DB: common.GetDB(), Ctx: common.GetCtx()}
}

func (ctl ProblemController) ProblemList(c *gin.Context) {
	var req vo.ProblemListRequest
	var problem, p model.Problem
	code := vo.OK

	defer func() {
		resp := vo.ProblemListResponse{
			ProblemID:         problemID,
			Status:            status,
			Difficulty:        difficulty,
			Acceptance:        acceptance,
			Global_Acceptance: globalAcceptance,
		}
		c.JSON(http.StatusOK, resp)
		utils.LogReqRespBody(req, resp, "CreateMember")
	}()

	if err := c.ShouldBindJSON(&req); err != nil {
		code = vo.UnknownError
		log.Println("CreateMember: ShouldBindJSON error")
		return
	}

	//problem = model.Problem{ProblemID: req.ProblemID, Status: req.Status, RoleId: 1}

	if err := ctl.DB.Where("user_name = ?", req.Username).Take(&u).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctl.DB.Create(&user)
			log.Println("CreateMember:Successfully create, username:" + user.Username)
			return
		} else {
			code = vo.UnknownError
			log.Println("CreateMember:Unknown-error while creating")
			return
		}
	}

	//User existed
	code = vo.UserHasExisted
	log.Println("CreateMember:UserExisted")
	return
}

func (ctl UserController) Login(c *gin.Context) {
	//TODO implement me
	panic("implement me")
	var req vo.LoginRequest
	var user, u model.User
	code := vo.OK
	defer func() {
		resp := vo.LoginResponse{
			Code: code,
		}
		c.JSON(http.StatusOK, resp)
		utils.LogReqRespBody(req, resp, "XXXXXXXXXXX")
	}()
	if err := c.ShouldBindJSON(&req); err != nil {
		code = vo.UnknownError
		log.Println("Login: ShouldBindJSON error")
		return
	}
	user = model.User{Username: req.Username, Password: req.Password, RoleId: 1}
	if err := ctl.DB.Where("user_name = ?", req.Username).Take(&u).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctl.DB.Create(&user)
			log.Println("CreateMember:Successfully create, username:" + user.Username)
			return
		} else {
			code = vo.UnknownError
			log.Println("CreateMember:Unknown-error while creating")
			return
		}
	}
}

func (ctl UserController) Logout(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
