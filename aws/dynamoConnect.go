// package db

// import (
// 	"fmt"

// 	"github.com/aws/aws-sdk-go/aws"
// 	"github.com/aws/aws-sdk-go/aws/session"
// 	"github.com/aws/aws-sdk-go/service/dynamodb"
// 	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
// )

// var svc *dynamodb.DynamoDB

// func startSession() *dynamodb.DynamoDB {

// // 	sess := session.Must(session.NewSessionWithOptions(session.Options{
// //     SharedConfigState: session.SharedConfigEnable,
// // }))

// 	session, err := session.NewSession(&aws.Config{
// 		Region: aws.String("us-west-2"),
// 		Endpoint: aws.String("http://localhost:8000")},
// 	)
// 	if err != nil {
// 		fmt.Println("Error connecting to local dynamodb")
// 		fmt.Println(err)
// 	}
// 	return dynamodb.New(session)
// }

// func ConnectDB() {
// 	svc = startSession()
// 	if svc != nil {
// 		fmt.Println("Connected to local DynamoDb")
// 	}
// }

// const tableName = "KittenTest"

// func CreateTable() {
// 	svc = startSession()
// 	input := &dynamodb.CreateTableInput{
// 		AttributeDefinitions: []*dynamodb.AttributeDefinition{
// 		 {
// 			AttributeName: aws.String("Make"),
// 			AttributeType: aws.String("S"),
// 		 },
// 		 {
// 			AttributeName: aws.String("Model"),
// 			AttributeType: aws.String("S"),
// 		 },
// 	},
// 	KeySchema: []*dynamodb.KeySchemaElement{
// 		{
// 			AttributeName: aws.String("Make"),
// 			KeyType: aws.String("HASH"),
// 		},
// 		{
// 			AttributeName: aws.String("Model"),
// 			KeyType: aws.String("RANGE"),
// 		},
// 	},
// 	ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
// 		ReadCapacityUnits: aws.Int64(10),
// 		WriteCapacityUnits: aws.Int64(10),
// 	},
// 	TableName: aws.String(tableName),
// 	}

// 	_, err := svc.CreateTable(input)
// 	if err != nil {
// 		fmt.Printf("Got error calling Create table\n %s", err)
// 	}
// 	fmt.Println("created table", tableName)
// }

// type Item struct {
// 	Make string
// 	Model string
// 	// year int
// }

// func AddItemToTable() {
// 	item := Item{
// 		Make: "Toyota",
// 		Model: "Corlolla",
// 	}
// 	fmt.Println("item: ", item)
// 	av, err := dynamodbattribute.MarshalMap(item)
// 	if err != nil {
// 		fmt.Println("error making item")
// 		return
// 	}

// 	input := &dynamodb.PutItemInput{
// 		Item: av,
// 		TableName: aws.String(tableName),
// 	}

// 	_, err = svc.PutItem(input)
// 	if err != nil {
// 		fmt.Println("Error inputting Item", err)
// 		return
// 	}
// 	fmt.Println("Item inserted")
// }

// func GetTableItems() {
// 	svc = startSession()
// 	result, err := svc.GetItem(&dynamodb.GetItemInput{
// 		TableName: aws.String(tableName),
// 		Key: map[string]*dynamodb.AttributeValue{
// 			"Make": {
// 				S: aws.String("Toyota"),
// 			},
// 			"Model": {
// 				S: aws.String("Corolla"),
// 			},
// 		},
// 	})
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(result)
// }

package db

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var svc *dynamodb.DynamoDB

func startSession() *dynamodb.DynamoDB {
	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String("us-west-2"),
		Endpoint: aws.String("http://localhost:8000"),
	})
	if err != nil {
		fmt.Println("Error connecting to local DynamoDB")
		fmt.Println(err)
	}
	return dynamodb.New(sess)
}

func ConnectDB() {
	svc = startSession()
	if svc != nil {
		fmt.Println("Connected to local DynamoDB")
	}
}

const tableName = "KittenTest2"

func CreateTable() {
	svc = startSession()
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("Make"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("Model"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("Make"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("Model"),
				KeyType:       aws.String("RANGE"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String(tableName),
	}

	_, err := svc.CreateTable(input)
	if err != nil {
		fmt.Printf("Got error calling CreateTable: %s\n", err)
		return
	}
	fmt.Println("Created table:", tableName)
}

type Item struct {
	Make  string
	Model string
	// year int
}

func AddItemToTable() {
	item := Item{
		Make:  "Toyota",
		Model: "Corolla",
	}
	fmt.Println("Item:", item)

	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		fmt.Println("Error marshaling item:", err)
		return
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		fmt.Println("Error inserting item:", err)
		return
	}

	fmt.Println("Item inserted successfully!")
}

func GetTableItems() {
	svc = startSession()
	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Make": {
				S: aws.String("Toyota"),
			},
			"Model": {
				S: aws.String("Corolla"),
			},
		},
	})
	if err != nil {
		fmt.Println("Error retrieving item:", err)
		return
	}
	fmt.Println(result)
}
