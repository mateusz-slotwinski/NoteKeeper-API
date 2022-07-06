package services

import (
	context "context"
	time "time"

	bson "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mongo "go.mongodb.org/mongo-driver/mongo"

	Database "NoteKeeperAPI/src/database"
	Models "NoteKeeperAPI/src/database/models"
	Helpers "NoteKeeperAPI/src/helpers"
	Requests "NoteKeeperAPI/src/types/requests"
)

var NoteCollection *mongo.Collection = Database.GetCollection(Database.DB, "notes")

type NotesService struct{}

func (v NotesService) GetNotes(UserID string) []Models.Note {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var notes []Models.Note

	cursor, err := NoteCollection.Find(ctx, bson.D{
		{Key: "Author", Value: UserID},
	})
	Helpers.PrintError(err)

	err = cursor.All(ctx, &notes)
	Helpers.PrintError(err)

	return notes
}

func (v NotesService) GetSingleNote(NoteID string) *Models.Note {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ID, err := primitive.ObjectIDFromHex(NoteID)
	Helpers.PrintError(err)

	var note Models.Note
	err = NoteCollection.FindOne(ctx, bson.M{"_id": ID}).Decode(&note)
	if err == mongo.ErrNoDocuments {
		return nil
	}

	return &note
}

func (v NotesService) CreateNote(noteData Requests.CreateNote) *Models.Note {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var note Models.Note

	UserID, err := primitive.ObjectIDFromHex(noteData.Author)
	Helpers.PrintError(err)

	checkAccount, err := AccountCollection.CountDocuments(ctx, bson.M{"_id": UserID})
	if err != nil || checkAccount == 0 {
		return nil
	}

	_, err = NoteCollection.InsertOne(ctx, bson.D{
		{Key: "Title", Value: noteData.Title},
		{Key: "Content", Value: noteData.Content},
		{Key: "Author", Value: noteData.Author},
	})
	Helpers.PrintError(err)

	return &note
}

func (v NotesService) UpdateNote(reqData Requests.CreateNote, NoteID string) *Models.Note {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ID, err := primitive.ObjectIDFromHex(NoteID)
	Helpers.PrintError(err)

	var note Models.Note

	_ = NoteCollection.FindOneAndReplace(ctx, bson.M{"_id": ID}, bson.D{
		{Key: "Title", Value: reqData.Title},
		{Key: "Content", Value: reqData.Content},
		{Key: "Author", Value: reqData.Author},
	}).Decode(&note)

	if err == mongo.ErrNoDocuments {
		return nil
	}

	return &note
}

func (v NotesService) DeleteNote(NoteID string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ID, err := primitive.ObjectIDFromHex(NoteID)
	Helpers.PrintError(err)

	_, err = NoteCollection.DeleteOne(ctx, bson.M{"_id": ID})

	return err == nil
}
