package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Directory to store all collections
const collectionsDir = "notestool"

func init() {
	// Create the collections directory if it doesn't exist
	if _, err := os.Stat(collectionsDir); os.IsNotExist(err) {
		os.Mkdir(collectionsDir, 0755)
	}
}

// Get the file path for a collection
func getCollectionFilePath(collection string) string {
	return fmt.Sprintf("%s/%s.txt", collectionsDir, collection)
}

// Create a new collection if it doesn't exist
func createCollectionIfNotExists(collectionName string) {
	collectionFilePath := getCollectionFilePath(collectionName)
	if _, err := os.Stat(collectionFilePath); os.IsNotExist(err) {
		createCollection(collectionName)
	}
}

// Create a new collection
func createCollection(collectionName string) {
	collectionFilePath := getCollectionFilePath(collectionName)
	file, err := os.Create(collectionFilePath)
	if err != nil {
		fmt.Println("Error creating collection:", err)
		return
	}
	file.Close()
	fmt.Printf("Collection '%s' created successfully.\n", collectionName)
}

// Load notes from a collection file
func loadNotes(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		return []string{}
	}
	defer file.Close()

	var notes []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		notes = append(notes, scanner.Text())
	}
	return notes
}

// Save notes to a collection file
func saveNotes(filename string, notes []string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	defer file.Close()

	for _, note := range notes {
		file.WriteString(note + "\n")
	}
}

// Display notes from a collection
func showNotes(notes []string) {
	if len(notes) == 0 {
		fmt.Println("No notes available.")
	} else {
		for idx, note := range notes {
			fmt.Printf("%03d - %s\n", idx+1, note)
		}
	}
}

// Prompt user to add a note
func addNote() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the note text:")
	note, _ := reader.ReadString('\n')
	return strings.TrimSpace(note)
}

// Prompt user to delete a note
func deleteNote() int {
	fmt.Println("Enter the number of the note to remove or 0 to cancel:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	index, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return -1
	}
	return index
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the name of the collection: ")
	collectionName, _ := reader.ReadString('\n')
	collectionName = strings.TrimSpace(collectionName)

	fmt.Print("Welcome to the notes tool \n")

	for {
		fmt.Println("\nSelect operation:\n1. Show notes.\n2. Add a note.\n3. Delete a note.\n4. Exit.")

		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		switch choice {
		case 1:
			collection := collectionName // Entered collection name
			createCollectionIfNotExists(collection)
			collectionFilePath := getCollectionFilePath(collection)
			notes := loadNotes(collectionFilePath)
			showNotes(notes)
		case 2:
			collection := collectionName // Entered collection name
			createCollectionIfNotExists(collection)
			collectionFilePath := getCollectionFilePath(collection)
			notes := loadNotes(collectionFilePath)
			note := addNote()
			notes = append(notes, note)
			saveNotes(collectionFilePath, notes)
		case 3:
			collection := collectionName // Entered collection name
			createCollectionIfNotExists(collection)
			collectionFilePath := getCollectionFilePath(collection)
			notes := loadNotes(collectionFilePath)
			index := deleteNote()
			if index >= 1 && index <= len(notes) {
				notes = append(notes[:index-1], notes[index:]...)
				saveNotes(collectionFilePath, notes)
			} else if index != 0 {
				fmt.Println("Invalid index.")
			}
		case 4:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please select a valid operation.")
		}
	}
}
