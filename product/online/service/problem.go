package service

import (
	"encoding/json"
	"log"
	"net/http"
	"online/define"
	"online/helper"
	"online/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetProblemList
// @Tag 公共方法
// @Summary 问题列表
// @Param page query int false "page"
// @Param size query int false "size"
// @Param keyword query string false "keyword"
// @Param category_identity query string false "category_identity"
// @Success 200 {string} json "{"code":"200", "msg":"","data":""}"
// @Router /problem-list [get]
func GetProblemList(c *gin.Context) {
	size, _ := strconv.Atoi(c.DefaultQuery("size", define.DefaultSize))
	page, err := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	if err != nil {
		log.Println("GetProblemList page strconv Error :", err)
		return
	}
	page = (page - 1) * size
	var count int64
	keyword := c.Query("keyword")
	categoryIdentity := c.Query("category_identity")
	tx := models.GetProblemList(keyword, categoryIdentity)
	list := make([]*models.ProblemBasic, 0)
	err = tx.Count(&count).Omit("content").Offset(page).Limit(size).Find(&list).Error
	if err != nil {
		log.Println("Get Problem List Error :", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": map[string]interface{}{
			"list":  list,
			"count": count,
		},
	})
}

// GetProblemDetail
// @Tag 公共方法
// @Summary 问题详情
// @Param identity query string false "problem_identity"
// @Success 200 {string} json "{"code":"200", "msg":"","data":""}"
// @Router /problem-detail [get]
func GetProblemDetail(c *gin.Context) {
	identity := c.Query("identity")
	if identity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "问题唯一标识不能为空",
		})
		return
	}
	data := new(models.ProblemBasic)
	err := models.DB.Where("identity = ?", identity).
		Preload("ProblemCategories").Preload("ProblemCategories.CategoryBasic").
		First(&data).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "问题不存在",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get ProblemDetail Error:" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": data,
	})
}

// GetProblemCreate
// @Tag 管理员私有方法
// @Summary 问题创建
// @Param authorization header string true "authorization"
// @Param title formData string true "title"
// @Param content formData string true "content"
// @Param max_runtime formData string false "max_runtime"
// @Param max_mem formData string false "max_mem"
// @Param category_ids formData array flase "category_ids"
// @Param test_cases formData array true "test_cases"
// @Success 200 {string} json "{"code":"200", "msg":"","data":""}"
// @Router /problem-create [post]
func GetProblemCreate(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")
	max_runtime, _ := strconv.Atoi(c.PostForm("max_runtime"))
	max_mem, _ := strconv.Atoi(c.PostForm("max_mem"))
	category_ids := c.PostFormArray("category_ids")
	test_cases := c.PostFormArray("test_cases")

	if title == "" || content == "" || len(category_ids) == 0 || len(test_cases) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数不能为空",
		})
		return
	}
	identity := helper.GetUUID()
	data := &models.ProblemBasic{
		Identity:   identity,
		Title:      title,
		Content:    content,
		MaxRuntime: max_runtime,
		MaxMem:     max_mem,
	}

	//处理分类
	CategoryBasics := make([]*models.ProblemCategory, 0)
	for _, id := range category_ids {
		category_id, _ := strconv.Atoi(id)
		CategoryBasics = append(CategoryBasics, &models.ProblemCategory{
			ProblemId:  data.ID,
			CategoryId: uint(category_id),
		})
	}
	data.ProblemCategories = CategoryBasics
	//处理测试用例
	testCaseBasics := make([]*models.TestCase, 0)
	for _, testCase := range test_cases {
		//例: {"input":"1 2\n", "output":"3\n"}
		caseMap := make(map[string]string)
		err := json.Unmarshal([]byte(testCase), &caseMap)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "测试用例格式错误",
			})
			return
		}

		if _, ok := caseMap["input"]; !ok {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "测试用例格式错误",
			})
			return
		}

		if _, ok := caseMap["output"]; !ok {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "测试用例格式错误",
			})
			return
		}

		testCaseBasic := &models.TestCase{
			Identity:        helper.GetUUID(),
			ProblemIdentity: identity,
			Input:           caseMap["input"],
			Output:          caseMap["output"],
		}
		testCaseBasics = append(testCaseBasics, testCaseBasic)
	}
	data.TestCases = testCaseBasics

	//创建问题
	err := models.DB.Create(data).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Problem Create Error:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"identity": identity,
		},
	})
}
