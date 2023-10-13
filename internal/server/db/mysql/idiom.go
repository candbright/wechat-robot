package mysql

import (
	"github.com/candbright/wechat-robot/internal/server/db/mysql/model"
	"github.com/candbright/wechat-robot/internal/server/db/options"
	"github.com/candbright/wechat-robot/internal/server/repo"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (DB *DB) AddIdiom(data repo.Idiom) error {
	var before model.Idiom
	if err := DB.Where(model.Idiom{Word: data.Word}).First(&before).Error; err == nil {
		return errors.Errorf("word %s exist", data.Word)
	} else if err != gorm.ErrRecordNotFound {
		return errors.WithStack(err)
	}
	id := uuid.NewString()
	idiom := model.Idiom{
		Id:          id,
		Word:        data.Word,
		Pinyin:      data.Pinyin,
		Abbr:        data.Abbr,
		Explanation: data.Explanation,
	}
	quote := model.Quote{
		Id:   id,
		Text: data.Quote.Text,
		Book: data.Quote.Book,
	}
	source := model.Source{
		Id:   id,
		Text: data.Source.Text,
		Book: data.Source.Book,
	}
	tx := DB.Begin()
	if err := tx.Create(&idiom).Error; err != nil {
		tx.Rollback()
		return errors.WithStack(err)
	}
	if err := tx.Create(&quote).Error; err != nil {
		tx.Rollback()
		return errors.WithStack(err)
	}
	if err := tx.Create(&source).Error; err != nil {
		tx.Rollback()
		return errors.WithStack(err)
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return errors.WithStack(err)
	}
	return nil
}

func (DB *DB) GetIdioms(opts ...options.Option) ([]repo.Idiom, error) {
	var results []repo.Idiom
	var idioms []model.Idiom
	if err := DB.options(opts...).Find(&idioms).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	for _, idiom := range idioms {
		result := repo.Idiom{
			Word:        idiom.Word,
			Pinyin:      idiom.Pinyin,
			Abbr:        idiom.Abbr,
			Explanation: idiom.Explanation,
		}

		var quote model.Quote
		var source model.Source
		if err := DB.options(options.WhereId(idiom.Id)).Take(&quote).Error; err == nil && quote.Text != "" {
			result.Quote = repo.Quote{Book: quote.Book, Text: quote.Text}
		}
		if err := DB.options(options.WhereId(idiom.Id)).Take(&source).Error; err == nil && source.Text != "" {
			result.Source = repo.Source{Book: source.Book, Text: source.Text}
		}
		results = append(results, result)
	}
	return results, nil
}
