package comment

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_mater/database"
	"shop_mater/models"
	"shop_mater/utils"
	"strconv"
	"time"
)

type CommentListRtnJson struct {
	Comment  string                  `json:"comment"`
	TypeId   int                     `json:"type_id"`
	ValueId  int                     `json:"value_id"`
	Id       int                     `json:"id"`
	AddTime  string                  `json:"add_time"`
	UserInfo models.User             `json:"user_info"`
	PicList  []models.CommentPicture `json:"pic_list"`
}

// List 评论列表
func List(c *gin.Context) {
	typeId := c.DefaultQuery("typeId", "0")
	valueId := c.DefaultQuery("valueId", "0")
	page := c.DefaultQuery("page", "1")
	size := c.DefaultQuery("size", "10")
	showType := c.DefaultQuery("showType", "0")

	intTypeId, _ := strconv.Atoi(typeId)
	intValueId, _ := strconv.Atoi(valueId)
	intShowType, _ := strconv.Atoi(showType)
	intPage, _ := strconv.Atoi(page)
	intSize, _ := strconv.Atoi(size)

	var comments []models.Comment
	var pageData utils.PageData

	var Total int64

	if intShowType != 1 {
		database.DB.Model(models.Comment{}).Where("type_id = ?", intTypeId).Where("value_id = ?", intValueId).Count(&Total).
			Offset((intPage - 1) * intSize).Limit(intSize).Find(&comments)
	} else {
		Total = database.DB.Raw("select c.* "+
			"from comment c "+
			"inner join comment_picture cp "+
			"on c.id = cp.comment_id "+
			"where c.type_id = ? and c.value_id = ? offset ? limit ? ", typeId, valueId, (intPage-1)*intSize, intSize).Scan(&comments).RowsAffected
	}

	pageData = utils.PageData{
		Count:     int(Total),
		Current:   intPage,
		TotalPage: (int(Total) + intSize - 1) / intSize,
		PageSize:  intSize,
		Data:      comments,
	}

	var rtnComments []CommentListRtnJson

	for _, item := range pageData.Data.([]models.Comment) {

		var users []models.User
		var commentPictures []models.CommentPicture

		database.DB.Model(models.User{}).
			Select("username", "avatar", "nickname").
			Where("id = ?", item.UserID).Find(&users)

		database.DB.Model(models.CommentPicture{}).Where("comment_id = ?", item.ID).Find(&commentPictures)

		rtnComments = append(rtnComments, CommentListRtnJson{
			Comment:  item.Content,
			TypeId:   int(item.TypeID),
			ValueId:  item.ValueID,
			Id:       int(item.ID),
			AddTime:  time.Unix(item.AddTime, 0).Format("2006-01-02 15:04:05 PM"),
			UserInfo: users[0],
			PicList:  commentPictures,
		})
	}
	pageData.Data = rtnComments
	c.JSON(http.StatusOK, utils.SuccessReturn(pageData))
}
