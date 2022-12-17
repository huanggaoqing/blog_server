package dao

import (
	"blog_server/db"
	"blog_server/dto/request"
	"blog_server/module/dbModule"
	"blog_server/resp"
	"gorm.io/gorm"
)

type ArticleDao struct {
	db *gorm.DB
}

func (a *ArticleDao) Save(data *dbModule.Article) (string, error) {
	err := a.db.Create(data).Error
	if err != nil {
		return "", err
	}
	return data.ArticleId, err
}

func (a *ArticleDao) GetList(params *request.GetArticleListRequest) ([]*dbModule.ArticleItem, error) {
	data := make([]*dbModule.ArticleItem, 0)
	err := a.db.Raw("SELECT article_id,article_title,is_stick,user_name,`blog_article`.`create_time`,`blog_article`.`update_time` FROM `blog_article` LEFT JOIN `blog_user` ON blog_article.user_id = blog_user.user_id WHERE `blog_article`.`user_id` = ? ORDER BY is_stick DESC, create_time DESC LIMIT ? OFFSET ?",
		params.UserId,
		params.Size,
		(params.Page-1)*params.Size,
	).Scan(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, resp.NOT_ARTICLES
		}
		return nil, err
	}
	return data, nil
}

func (a *ArticleDao) GetListCount(userId int) (int, error) {
	count := 0
	err := a.db.Model(&dbModule.Article{}).Select("count(article_id)").Where("user_id = ?", userId).Scan(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (a *ArticleDao) GetDetail(articleId string) (*dbModule.ArticleItem, error) {
	data := &dbModule.ArticleItem{}
	err := a.db.Raw(
		"SELECT article_id,article_content,article_title,is_stick,user_name,group_id,`blog_article`.`create_time`,`blog_article`.`update_time` FROM `blog_article` LEFT JOIN `blog_user` ON blog_article.user_id = blog_user.user_id WHERE article_id = ?",
		articleId,
	).Scan(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, resp.PARAMS_ERROR
		}
		return nil, err
	}
	return data, nil
}

func NewArticleDao() *ArticleDao {
	return &ArticleDao{db: db.GetDataBase()}
}
