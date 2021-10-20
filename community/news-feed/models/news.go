package models

import (
	"fmt"

	real "github.com/brianvoe/gofakeit/v6"
)
  

type NewsArticle struct{
	Title string `json:"title"`
	Author string `json:"author"`
	Date string `json:"date"`
	Description string `json:"description"`
	Body string `json:"body"`
}

type NewsBrief struct{
	Title string `json:"title"`
	Description string `json:"description"`
}


func (n NewsArticle) New() *NewsArticle {
	return &NewsArticle{
		Title: fmt.Sprintf("%s %s %s %s %s",real.Noun(),real.Verb(),real.Adverb(),real.Preposition(),real.Noun()),
		Author: real.Name(),
		Date: fmt.Sprintf("%d-%s-%d",real.Day(),real.MonthString(),real.Year()),
		Description:fmt.Sprintf("%s %s,%s %s %s.",real.Word(),real.Word(),real.BuzzWord(),real.HipsterWord(),real.Word()),
		Body: real.HipsterParagraph(5,5,5,","),
	}
}
func (n *NewsArticle) Breif() *NewsBrief{
	article := n.New()
	return &NewsBrief{
		Title: article.Title,
		Description: article.Description,
	}
}

func (n NewsArticle) GenArticles(amount int) []*NewsArticle{
	var articles []*NewsArticle
	for i:=1; i <=amount; i++{
		articles = append(articles,n.New())
	}
	return articles
}
