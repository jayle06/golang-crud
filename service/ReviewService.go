package service

import (
	"ocg.com/train/model"
	"ocg.com/train/repository"
)

func GetAllReview() []*model.Review {
	return repository.Reviews.FindAllReview()
}

func CreateReview(review *model.Review) *model.Review {
	newReview := repository.Reviews.CreateNewReview(review)
	bookId := newReview.BookId
	book, _ := FindBookById(bookId)
	UpdateBookRating(book)
	return newReview
}

func UpdateReView(Id int64, newReview *model.Review) *model.Review {
	review, _ := repository.Reviews.FindReviewById(Id)
	review = CreateReview(newReview)
	return review
}

func GetReviewByBookId(bookId int64) map[int64][]*model.Review {
	return repository.Reviews.GetReviewByBookId(bookId)
}
