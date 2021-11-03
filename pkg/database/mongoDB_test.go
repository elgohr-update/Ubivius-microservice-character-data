package database

import (
	"context"
	"testing"

	"github.com/Ubivius/microservice-character-data/pkg/data"
	"github.com/google/uuid"
)

func TestMongoDBConnectionAndShutdownIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}

	mp := NewMongoCharacters()
	if mp == nil {
		t.Fail()
	}
	mp.CloseDB()
}

func TestMongoDBAddCharacterIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}

	character := &data.Character{
		Name:   "testName",
		UserID: "a2181017-5c53-422b-b6bc-036b27c04fc8",
	}

	mp := NewMongoCharacters()
	err := mp.AddCharacter(context.Background(), character)
	if err != nil {
		t.Errorf("Failed to add character to database")
	}
	mp.CloseDB()
}

func TestMongoDBUpdateCharacterIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}

	character := &data.Character{
		ID:     uuid.NewString(),
		UserID: "a2181017-5c53-422b-b6bc-036b27c04fc8",
	}

	mp := NewMongoCharacters()
	err := mp.UpdateCharacter(context.Background(), character)
	if err != nil {
		t.Fail()
	}
	mp.CloseDB()
}

func TestMongoDBGetCharactersIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}

	mp := NewMongoCharacters()
	characters := mp.GetCharacters(context.Background())
	if characters == nil {
		t.Fail()
	}

	mp.CloseDB()
}

func TestMongoDBGetCharacterByIDIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}

	mp := NewMongoCharacters()
	_, err := mp.GetCharacterByID(context.Background(), "a2181017-5c53-422b-b6bc-036b27c04fc8")
	if err != nil {
		t.Fail()
	}

	mp.CloseDB()
}

func TestMongoDBGetCharactersByUserIDIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Test skipped during unit tests")
	}

	mp := NewMongoCharacters()
	_, err := mp.GetCharactersByUserID(context.Background(), "a2181017-5c53-422b-b6bc-036b27c04fc8")
	if err != nil {
		t.Fail()
	}

	mp.CloseDB()
}
