package aws

import (
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/sirupsen/logrus"
)

var AWS *session.Session
var AccessKeyID string
var SecretAccessKey string
var MyRegion string

func ConnectAws() *session.Session {
	AccessKeyID = os.Getenv("AWS_ACCESS_KEY_ID")
	SecretAccessKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
	MyRegion = os.Getenv("AWS_REGION")
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(MyRegion),
			Credentials: credentials.NewStaticCredentials(
				AccessKeyID,
				SecretAccessKey,
				"", // a token will be created when the session it's used.
			),
		})
	if err != nil {
		panic(err)
	}
	AWS = sess
	logrus.Trace("AWS Connected!")
	return sess
}

func composeS3Url(filePath string) string {
	S3B := os.Getenv("S3_BUCKET")
	if filePath[:1] != "/" {
		filePath = "/"+filePath
	}
	return "https://" + S3B + "." + "s3-" + MyRegion + ".amazonaws.com" + filePath
}

func ComposeOneeCDNUrl(filePath string) string {
	if filePath[:1] != "/" {
		filePath = "/"+filePath
	}
	return "https://cdn.onee.id" + filePath
}

func UploadToS3(filePath string, savePath *string) (string, error)  {

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	awsPath := filePath
	if savePath != nil {
		awsPath = *savePath
	}

	S3B := os.Getenv("S3_BUCKET")
	uploader := s3manager.NewUploader(AWS)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(S3B),
		ACL:    aws.String("public-read"),
		Key:    aws.String(awsPath),
		Body:   file,
	})

	if err != nil {
		logrus.Warn("FAIL Upload! to AWS S3 Bucket "+S3B)
	}

	return composeS3Url(awsPath), nil
}

func UploadImageS3(file io.Reader, savePath string) (string, error)  {
	S3B := os.Getenv("S3_BUCKET")
	uploader := s3manager.NewUploader(AWS)

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(S3B),
		ACL:    aws.String("public-read"),
		Key:    aws.String(savePath),
		Body:   file,
	})

	if err != nil {
		logrus.Warn("FAIL Upload! to AWS S3 Bucket "+S3B)
	}

	return composeS3Url(savePath), nil
}

func DeleteFileS3(savePath string) error {
	S3B := os.Getenv("S3_BUCKET")
	MyRegion = os.Getenv("AWS_REGION")

	svc := s3.New(AWS)
	// Delete the item
	_, err := svc.DeleteObject(&s3.DeleteObjectInput{Bucket: &S3B, Key: &savePath})
	if err != nil {
		logrus.Printf("Error occurred while waiting for object to be deleted, %+v\n", err)
		return err
	}
	err = svc.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: &S3B,
		Key:    &savePath,
	})
	if err != nil {
		logrus.Printf("Error occurred while waiting for object to be deleted, %+v\n", err)
		return err
	}

	return nil
}
