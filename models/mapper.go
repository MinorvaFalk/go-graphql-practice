package models

// func MapBooktoGeneratedModel(books []Book) []*gqlModel.Book {
// 	var gqlBook []*gqlModel.Book

// 	for index, book := range books {
// 		gqlBook = append(gqlBook, &gqlModel.Book{
// 			ID:          fmt.Sprint(book.ID),
// 			Title:       book.Title,
// 			Description: &books[index].Description,
// 			Rating:      &books[index].Rating,
// 			AuthorID:    fmt.Sprint(book.AuthorRefer),
// 		})
// 	}

// 	return gqlBook
// }

// func MapSingleBooktoGeneratedModel(book Book) *gqlModel.Book {
// 	return &gqlModel.Book{
// 		ID:          fmt.Sprint(book.ID),
// 		Title:       book.Title,
// 		Description: &book.Description,
// 		Rating:      &book.Rating,
// 		AuthorID:    fmt.Sprint(book.AuthorRefer),
// 	}
// }

// func MapAuthortoGeneratedModel(authors []Author) []*gqlModel.Author {
// 	var gqlAuthor []*gqlModel.Author

// 	for _, author := range authors {
// 		gqlAuthor = append(gqlAuthor, &gqlModel.Author{
// 			ID:   fmt.Sprint(author.ID),
// 			Name: author.Name,
// 		})
// 	}

// 	return gqlAuthor
// }

// func MapInputBooktoBookModel(genBook gqlModel.BookInput) (Book, error) {
// 	authorID, err := strconv.Atoi(genBook.AuthorID)

// 	if err != nil {
// 		return Book{}, err
// 	}

// 	return Book{
// 		Title:       genBook.Title,
// 		Description: *genBook.Description,
// 		Rating:      *genBook.Rating,
// 		AuthorRefer: uint(authorID),
// 	}, nil
// }
