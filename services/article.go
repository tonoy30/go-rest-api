package services

import (
	"context"
	"log"

	"github.com/tonoy30/practice-rest/dbs"
	"github.com/tonoy30/practice-rest/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	COLLECTION_NAME = "practice_articles"
)

var ctx = context.TODO()
var collection = dbs.GetMongoCollection(COLLECTION_NAME)

func GetArticles() ([]models.Article, error) {
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Println(err)
	}
	defer cursor.Close(ctx)
	var articles []models.Article
	for cursor.Next(ctx) {
		var article models.Article
		if err := cursor.Decode(&article); err != nil {
			log.Println(err)
		}
		articles = append(articles, article)
	}
	return articles, err
}
func CreateArticle(article models.Article) (models.Article, error) {
	result, err := collection.InsertOne(ctx, article)
	if err != nil {
		log.Println(err)
	}
	article.ID = result.InsertedID.(primitive.ObjectID)
	return article, err
}

func GetArticleById(articleID string) (models.Article, error) {
	var article models.Article
	objectID, err := primitive.ObjectIDFromHex(articleID)
	if err != nil {
		log.Println(err)
	}
	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&article)
	if err != nil {
		log.Println(err)
	}
	return article, err
}

func DeleteArticle(articleID string) (bool, error) {
	objectID, err := primitive.ObjectIDFromHex(articleID)
	if err != nil {
		log.Println(err)
	}
	result, err := collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		log.Println(err)
	}
	return result.DeletedCount > 0, err
}

func UpdateArticle(articleID string, article models.Article) (models.Article, error) {
	objectID, err := primitive.ObjectIDFromHex(articleID)
	if err != nil {
		log.Println(err)
	}
	_, err = collection.UpdateOne(ctx, bson.M{"_id": objectID}, bson.M{"$set": article})
	if err != nil {
		log.Println(err)
	}
	result, err := GetArticleById(articleID)
	if err != nil {
		log.Println(err)
	}
	return result, err
}
