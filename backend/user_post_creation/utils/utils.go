package utils

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"user_post_creation/model"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/gorilla/mux"
)

func Login(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	valid, err := ValidateUser(user.Email, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !valid {
		http.Error(w, fmt.Errorf("User not found or password mismatched").Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	fmt.Fprintln(w, "User logged in successfully")
}

func Signup(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := RegisterUser(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	fmt.Fprintln(w, "User logged in successfully")

}

func UserDetails(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userid := params["id"]
	userDetails, err := GetUserDetails(userid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userDetails)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["id"]

	// Parse the updated user details from the request body
	var updatedUser model.User
	err := json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = EditUserDetails(userID, updatedUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User updated successfully")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["id"]

	err := DeleteUSer(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User deleted successfully")
}

func UploadImage(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["id"]
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to get file from form data", http.StatusBadRequest)
		return
	}
	defer file.Close()
	fileUrl, err := PutFileInS3(file, handler)
	if err != nil {
		http.Error(w, "Faled to upload file in s3", http.StatusBadRequest)
		return
	}
	if err := StoreFileInUserDetails(fileUrl, userID); err != nil {
		http.Error(w, "Unable to store file in db", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fileUrl)
}

func CreateS3Session() (*session.Session, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-south1"),
	})
	return sess, err
}

func PutFileInS3(file multipart.File, handler *multipart.FileHeader) (string, error) {
	sess, err := CreateS3Session()
	if err != nil {
		return "", err
	}

	svc := s3.New(sess)
	bucketName := "s3export"
	fileName := handler.Filename
	input := &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
		Body:   file,
	}
	_, err = svc.PutObject(input)
	if err != nil {
		return "", err
	}
	fileURL := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucketName, fileName)
	return fileURL, nil
}
