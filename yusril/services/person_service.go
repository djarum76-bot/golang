package services

import (
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"yusril/db"
	"yusril/helper"
	"yusril/models"

	"github.com/lib/pq"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func AddPersonWithImage(person models.Person, image *multipart.FileHeader, images []*multipart.FileHeader) (models.Response, error) {
	var res models.Response
	var err error
	arrImagesURL := []string{}

	conn := db.CreateCon()

	sqlStatement := "INSERT INTO persons (user_id, nickname, first_name, last_name, email, job, phone, date_of_birth, place_of_birth, place_of_residence, instagram, twitter, facebook, linkedin, major, college, organization, party, last_activity, first_meet, category, tags, image, images, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25)"

	//image
	src, err := image.Open()
	if err != nil {
		return res, err
	}
	defer src.Close()

	imageURL := "image/" + image.Filename

	dst, err := os.Create(imageURL)
	if err != nil {
		return res, err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return res, err
	}

	//images
	if len(images) != 0 {
		for _, image := range images {
			src, err := image.Open()
			if err != nil {
				return res, err
			}

			defer src.Close()

			imagesURL := "image/" + image.Filename

			dst, err := os.Create(imagesURL)
			if err != nil {
				return res, err
			}
			defer dst.Close()

			if _, err = io.Copy(dst, src); err != nil {
				return res, err
			}

			arrImagesURL = append(arrImagesURL, imagesURL)
		}
	} else {
		arrImagesURL = append(arrImagesURL, "")
	}

	_, err = conn.Exec(sqlStatement, person.UserID, person.Nickname, person.FirstName, person.LastName, pq.Array(person.Email), person.Job, person.Phone, person.DateOfBirth, person.PlaceOfBirth, pq.Array(person.PlaceOfResidence), person.Instagram, person.Twitter, person.Facebook, person.LinkedIn, pq.Array(person.Major), pq.Array(person.College), pq.Array(person.Organization), pq.Array(person.Party), pq.Array(person.LastActivity), person.FirstMeet, person.Category, pq.Array(person.Tags), imageURL, pq.Array(arrImagesURL), person.CreatedAt)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success Add Person"
	res.Data = nil

	return res, nil
}

func AddPersonWithoutImage(person models.Person, images []*multipart.FileHeader) (models.Response, error) {
	var res models.Response
	var err error
	arrImagesURL := []string{}

	conn := db.CreateCon()

	sqlStatement := "INSERT INTO persons (user_id, nickname, first_name, last_name, email, job, phone, date_of_birth, place_of_birth, place_of_residence, instagram, twitter, facebook, linkedin, major, college, organization, party, last_activity, first_meet, category, tags, image, images, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25)"

	//images
	if len(images) != 0 {
		for _, image := range images {
			src, err := image.Open()
			if err != nil {
				return res, err
			}

			defer src.Close()

			imagesURL := "image/" + image.Filename

			dst, err := os.Create(imagesURL)
			if err != nil {
				return res, err
			}
			defer dst.Close()

			if _, err = io.Copy(dst, src); err != nil {
				return res, err
			}

			arrImagesURL = append(arrImagesURL, imagesURL)
		}
	} else {
		arrImagesURL = append(arrImagesURL, "")
	}

	_, err = conn.Exec(sqlStatement, person.UserID, person.Nickname, person.FirstName, person.LastName, pq.Array(person.Email), person.Job, person.Phone, person.DateOfBirth, person.PlaceOfBirth, pq.Array(person.PlaceOfResidence), person.Instagram, person.Twitter, person.Facebook, person.LinkedIn, pq.Array(person.Major), pq.Array(person.College), pq.Array(person.Organization), pq.Array(person.Party), pq.Array(person.LastActivity), person.FirstMeet, person.Category, pq.Array(person.Tags), "", pq.Array(arrImagesURL), person.CreatedAt)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success Add Person"
	res.Data = nil

	return res, nil
}

func GetAllPerson(userID int) (models.Response, error) {
	var res models.Response
	var err error
	var person models.Person
	arrPerson := []models.Person{}

	conn := db.CreateCon()

	sqlStatement := "SELECT * FROM persons WHERE user_id = ($1) ORDER BY nickname"

	rows, err := conn.Query(sqlStatement, userID)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&person.ID, &person.UserID, &person.Nickname, &person.FirstName, &person.LastName, pq.Array(&person.Email), &person.Job, &person.Phone, &person.DateOfBirth, &person.PlaceOfBirth, pq.Array(&person.PlaceOfResidence), &person.Instagram, &person.Twitter, &person.Facebook, &person.LinkedIn, pq.Array(&person.Major), pq.Array(&person.College), pq.Array(&person.Organization), pq.Array(&person.Party), pq.Array(&person.LastActivity), &person.FirstMeet, &person.Category, pq.Array(&person.Tags), &person.Image, pq.Array(&person.Images), &person.CreatedAt)
		if err != nil {
			return res, err
		}

		email, err := helper.DecryptList(person.Email)
		if err != nil {
			return res, err
		}

		person.Email = email

		arrPerson = append(arrPerson, person)
	}

	res.Status = http.StatusOK
	res.Message = "Success Get All Person"
	res.Data = arrPerson

	return res, nil
}

func GetPerson(userID int, ID int) (models.Response, error) {
	var res models.Response
	var err error
	var person models.Person

	conn := db.CreateCon()

	sqlStatement := "SELECT * FROM persons WHERE user_id = ($1) AND id = ($2)"

	err = conn.QueryRow(sqlStatement, userID, ID).Scan(&person.ID, &person.UserID, &person.Nickname, &person.FirstName, &person.LastName, pq.Array(&person.Email), &person.Job, &person.Phone, &person.DateOfBirth, &person.PlaceOfBirth, pq.Array(&person.PlaceOfResidence), &person.Instagram, &person.Twitter, &person.Facebook, &person.LinkedIn, pq.Array(&person.Major), pq.Array(&person.College), pq.Array(&person.Organization), pq.Array(&person.Party), pq.Array(&person.LastActivity), &person.FirstMeet, &person.Category, pq.Array(&person.Tags), &person.Image, pq.Array(&person.Images), &person.CreatedAt)
	if err != nil {
		return res, err
	}

	email, err := helper.DecryptList(person.Email)
	if err != nil {
		return res, err
	}

	person.Email = email

	res.Status = http.StatusOK
	res.Message = "Success Get All Person"
	res.Data = person

	return res, nil
}

func UpdatePerson(person models.Person) (models.Response, error) {
	var res models.Response
	var err error

	conn := db.CreateCon()

	sqlStatement := "UPDATE persons SET nickname = $1, first_name = $2, last_name = $3, email = $4, job = $5, phone = $6, date_of_birth = $7, place_of_birth = $8, place_of_residence = $9, instagram = $10, twitter = $11, facebook = $12, linkedin = $13, major = $14, college = $15, organization = $16, party = $17, last_activity = $18, first_meet = $19, category = $20, tags = $21 WHERE id = $22"

	_, err = conn.Exec(sqlStatement, person.Nickname, person.FirstName, person.LastName, pq.Array(person.Email), person.Job, person.Phone, person.DateOfBirth, person.PlaceOfBirth, pq.Array(person.PlaceOfResidence), person.Instagram, person.Twitter, person.Facebook, person.LinkedIn, pq.Array(person.Major), pq.Array(person.College), pq.Array(person.Party), pq.Array(person.Organization), pq.Array(person.LastActivity), person.FirstMeet, person.Category, pq.Array(person.Tags), person.ID)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success Update Person"
	res.Data = nil

	return res, nil
}

func UpdateImage(image *multipart.FileHeader, id int) (models.Response, error) {
	var res models.Response
	var err error

	conn := db.CreateCon()

	sqlStatement := "UPDATE persons SET image = $1 WHERE id = $2"

	//image
	src, err := image.Open()
	if err != nil {
		return res, err
	}
	defer src.Close()

	imageURL := "image/" + image.Filename

	dst, err := os.Create(imageURL)
	if err != nil {
		return res, err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return res, err
	}

	_, err = conn.Exec(sqlStatement, imageURL, id)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success Update Image"
	res.Data = nil

	return res, nil
}

func UpdateImages(images []*multipart.FileHeader, id int) (models.Response, error) {
	var res models.Response
	var err error
	arrImagesURL := []string{}

	conn := db.CreateCon()

	sqlStatement := "UPDATE persons SET images = $1 WHERE id = $2"

	//images
	for _, image := range images {
		src, err := image.Open()
		if err != nil {
			return res, err
		}

		defer src.Close()

		imagesURL := "image/" + image.Filename

		dst, err := os.Create(imagesURL)
		if err != nil {
			return res, err
		}
		defer dst.Close()

		if _, err = io.Copy(dst, src); err != nil {
			return res, err
		}

		arrImagesURL = append(arrImagesURL, imagesURL)
	}

	_, err = conn.Exec(sqlStatement, pq.Array(arrImagesURL), id)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success Update Images"
	res.Data = nil

	return res, nil
}

func DeletePerson(id int) (models.Response, error) {
	var res models.Response
	var err error

	conn := db.CreateCon()

	sqlStatement := "DELETE FROM persons WHERE id = $1"

	_, err = conn.Exec(sqlStatement, id)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success Delete Person"
	res.Data = nil

	return res, nil
}

func SearchPerson(query string, category string, tag string, userID int) (models.Response, error) {
	var res models.Response
	var person models.Person
	arrPerson := []models.Person{}
	search := "'%" + query + "%'"

	conn := db.CreateCon()

	if category != "" && tag != "" {
		sqlStatement := "SELECT * FROM persons WHERE user_id = ($1) AND (LOWER(persons.nickname) LIKE " + search + " OR LOWER(persons.first_name) LIKE " + search + " OR LOWER(persons.last_name) LIKE " + search + " OR LOWER(persons.job) LIKE " + search + " OR LOWER(persons.place_of_birth) LIKE " + search + " OR LOWER(ARRAY_TO_STRING(place_of_residence, ' ')) LIKE " + search + " OR LOWER(persons.instagram) LIKE " + search + " OR LOWER(persons.twitter) LIKE " + search + " OR LOWER(persons.facebook) LIKE " + search + " OR LOWER(persons.linkedin) LIKE " + search + " OR LOWER(ARRAY_TO_STRING(major, ' ')) LIKE " + search + " OR LOWER(ARRAY_TO_STRING(college, ' ')) LIKE " + search + " OR LOWER(ARRAY_TO_STRING(organization, ' ')) LIKE " + search + " OR LOWER(ARRAY_TO_STRING(party, ' ')) LIKE " + search + " OR LOWER(ARRAY_TO_STRING(last_activity, ' ')) LIKE " + search + " OR LOWER(persons.first_meet) LIKE " + search + " OR LOWER(persons.phone) LIKE " + search + ") AND LOWER(persons.category) = ($2) AND LOWER(ARRAY_TO_STRING(tags, ' ')) LIKE " + "'%" + tag + "%'" + " ORDER BY nickname"

		rows, err := conn.Query(sqlStatement, userID, category)
		if err != nil {
			return res, err
		}
		defer rows.Close()

		for rows.Next() {
			err = rows.Scan(&person.ID, &person.UserID, &person.Nickname, &person.FirstName, &person.LastName, pq.Array(&person.Email), &person.Job, &person.Phone, &person.DateOfBirth, &person.PlaceOfBirth, pq.Array(&person.PlaceOfResidence), &person.Instagram, &person.Twitter, &person.Facebook, &person.LinkedIn, pq.Array(&person.Major), pq.Array(&person.College), pq.Array(&person.Organization), pq.Array(&person.Party), pq.Array(&person.LastActivity), &person.FirstMeet, &person.Category, pq.Array(&person.Tags), &person.Image, pq.Array(&person.Images), &person.CreatedAt)
			if err != nil {
				return res, err
			}

			email, err := helper.DecryptList(person.Email)
			if err != nil {
				return res, err
			}

			person.Email = email

			arrPerson = append(arrPerson, person)
		}
	} else if category != "" {
		sqlStatement := "SELECT * FROM persons WHERE user_id = ($1) AND (LOWER(persons.nickname) LIKE " + search + " OR LOWER(persons.first_name) LIKE " + search + " OR LOWER(persons.last_name) LIKE " + search + " OR LOWER(persons.job) LIKE " + search + " OR LOWER(persons.place_of_birth) LIKE " + search + " OR LOWER(ARRAY_TO_STRING(place_of_residence, ' ')) LIKE " + search + " OR LOWER(persons.instagram) LIKE " + search + " OR LOWER(persons.twitter) LIKE " + search + " OR LOWER(persons.facebook) LIKE " + search + " OR LOWER(persons.linkedin) LIKE " + search + " OR LOWER(ARRAY_TO_STRING(major, ' ')) LIKE " + search + " OR LOWER(ARRAY_TO_STRING(college, ' ')) LIKE " + search + " OR LOWER(ARRAY_TO_STRING(organization, ' ')) LIKE " + search + " OR LOWER(ARRAY_TO_STRING(party, ' ')) LIKE " + search + " OR LOWER(ARRAY_TO_STRING(last_activity, ' ')) LIKE " + search + " OR LOWER(persons.first_meet) LIKE " + search + " OR LOWER(ARRAY_TO_STRING(tags, ' ')) LIKE " + search + " OR LOWER(persons.phone) LIKE " + search + ") AND LOWER(persons.category) = ($2) ORDER BY nickname"

		rows, err := conn.Query(sqlStatement, userID, category)
		if err != nil {
			return res, err
		}
		defer rows.Close()

		for rows.Next() {
			err = rows.Scan(&person.ID, &person.UserID, &person.Nickname, &person.FirstName, &person.LastName, pq.Array(&person.Email), &person.Job, &person.Phone, &person.DateOfBirth, &person.PlaceOfBirth, pq.Array(&person.PlaceOfResidence), &person.Instagram, &person.Twitter, &person.Facebook, &person.LinkedIn, pq.Array(&person.Major), pq.Array(&person.College), pq.Array(&person.Organization), pq.Array(&person.Party), pq.Array(&person.LastActivity), &person.FirstMeet, &person.Category, pq.Array(&person.Tags), &person.Image, pq.Array(&person.Images), &person.CreatedAt)
			if err != nil {
				return res, err
			}

			email, err := helper.DecryptList(person.Email)
			if err != nil {
				return res, err
			}

			person.Email = email

			arrPerson = append(arrPerson, person)
		}
	} else if tag != "" {
		sqlStatement := "SELECT * FROM persons WHERE user_id = ($1) AND (LOWER(persons.nickname) LIKE " + search + " OR LOWER(persons.first_name) LIKE " + search + " OR LOWER(persons.last_name) LIKE " + search + " OR LOWER(persons.job) LIKE " + search + " OR LOWER(persons.place_of_birth) LIKE " + search + " OR LOWER(ARRAY_TO_STRING(place_of_residence, ' ')) LIKE " + search + " OR LOWER(persons.instagram) LIKE " + search + " OR LOWER(persons.twitter) LIKE " + search + " OR LOWER(persons.facebook) LIKE " + search + " OR LOWER(persons.linkedin) LIKE " + search + " OR LOWER(ARRAY_TO_STRING(major, ' ')) LIKE " + search + " OR LOWER(ARRAY_TO_STRING(college, ' ')) LIKE " + search + " OR LOWER(ARRAY_TO_STRING(organization, ' ')) LIKE " + search + " OR LOWER(ARRAY_TO_STRING(party, ' ')) LIKE " + search + " OR LOWER(ARRAY_TO_STRING(last_activity, ' ')) LIKE " + search + " OR LOWER(persons.first_meet) LIKE " + search + " OR LOWER(persons.phone) LIKE " + search + ") AND LOWER(ARRAY_TO_STRING(tags, ' ')) LIKE " + "'%" + tag + "%'" + " ORDER BY nickname"

		rows, err := conn.Query(sqlStatement, userID)
		if err != nil {
			return res, err
		}
		defer rows.Close()

		for rows.Next() {
			err = rows.Scan(&person.ID, &person.UserID, &person.Nickname, &person.FirstName, &person.LastName, pq.Array(&person.Email), &person.Job, &person.Phone, &person.DateOfBirth, &person.PlaceOfBirth, pq.Array(&person.PlaceOfResidence), &person.Instagram, &person.Twitter, &person.Facebook, &person.LinkedIn, pq.Array(&person.Major), pq.Array(&person.College), pq.Array(&person.Organization), pq.Array(&person.Party), pq.Array(&person.LastActivity), &person.FirstMeet, &person.Category, pq.Array(&person.Tags), &person.Image, pq.Array(&person.Images), &person.CreatedAt)
			if err != nil {
				return res, err
			}

			email, err := helper.DecryptList(person.Email)
			if err != nil {
				return res, err
			}

			person.Email = email

			arrPerson = append(arrPerson, person)
		}
	} else {
		sqlStatement := "SELECT * FROM persons WHERE user_id = ($1) AND (LOWER(persons.nickname) LIKE " + search + " OR LOWER(persons.first_name) LIKE " + search + " OR LOWER(persons.last_name) LIKE " + search + " OR LOWER(persons.job) LIKE " + search + " OR LOWER(persons.place_of_birth) LIKE " + search + " OR LOWER(ARRAY_TO_STRING(place_of_residence, ' ')) LIKE " + search + " OR LOWER(persons.instagram) LIKE " + search + " OR LOWER(persons.twitter) LIKE " + search + " OR LOWER(persons.facebook) LIKE " + search + " OR LOWER(persons.linkedin) LIKE " + search + " OR LOWER(ARRAY_TO_STRING(major, ' ')) LIKE " + search + " OR LOWER(ARRAY_TO_STRING(college, ' ')) LIKE " + search + " OR LOWER(ARRAY_TO_STRING(organization, ' ')) LIKE " + search + " OR LOWER(ARRAY_TO_STRING(party, ' ')) LIKE " + search + " OR LOWER(ARRAY_TO_STRING(last_activity, ' ')) LIKE " + search + " OR LOWER(persons.first_meet) LIKE " + search + " OR LOWER(ARRAY_TO_STRING(tags, ' ')) LIKE " + search + " OR LOWER(persons.phone) LIKE " + search + ") ORDER BY nickname"

		rows, err := conn.Query(sqlStatement, userID)
		if err != nil {
			return res, err
		}
		defer rows.Close()

		for rows.Next() {
			err = rows.Scan(&person.ID, &person.UserID, &person.Nickname, &person.FirstName, &person.LastName, pq.Array(&person.Email), &person.Job, &person.Phone, &person.DateOfBirth, &person.PlaceOfBirth, pq.Array(&person.PlaceOfResidence), &person.Instagram, &person.Twitter, &person.Facebook, &person.LinkedIn, pq.Array(&person.Major), pq.Array(&person.College), pq.Array(&person.Organization), pq.Array(&person.Party), pq.Array(&person.LastActivity), &person.FirstMeet, &person.Category, pq.Array(&person.Tags), &person.Image, pq.Array(&person.Images), &person.CreatedAt)
			if err != nil {
				return res, err
			}

			email, err := helper.DecryptList(person.Email)
			if err != nil {
				return res, err
			}

			person.Email = email

			arrPerson = append(arrPerson, person)
		}
	}

	res.Status = http.StatusOK
	res.Message = "Success Get All Person"
	res.Data = arrPerson

	return res, nil
}

func GetAllCategory(userID int) (models.Response, error) {
	var res models.Response
	var err error
	arrCategory := []string{"All"}
	var category string

	conn := db.CreateCon()

	sqlStatement := "SELECT DISTINCT category FROM persons WHERE user_id = ($1) AND category != ($2)"

	rows, err := conn.Query(sqlStatement, userID, "")
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&category)
		if err != nil {
			return res, err
		}

		if helper.IsUpper(category) {
			arrCategory = append(arrCategory, category)
		} else {
			arrCategory = append(arrCategory, cases.Title(language.Indonesian).String(category))
		}
	}

	res.Status = http.StatusOK
	res.Message = "Success Get All Category"
	res.Data = arrCategory

	return res, nil
}

func GetAllTags(userID int) (models.Response, error) {
	var res models.Response
	var err error
	arrTags := []string{}
	var tags string

	conn := db.CreateCon()

	sqlStatement := "SELECT DISTINCT LOWER(unnest(tags)) FROM persons WHERE user_id = ($1)"

	rows, err := conn.Query(sqlStatement, userID)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&tags)
		if err != nil {
			return res, err
		}

		arrTags = append(arrTags, cases.Title(language.Indonesian).String(tags))
	}

	res.Status = http.StatusOK
	res.Message = "Success Get All Tags"
	res.Data = arrTags

	return res, nil
}
