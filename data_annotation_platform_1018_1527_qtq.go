// 代码生成时间: 2025-10-18 15:27:40
package main

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// Annotation represents a data annotation.
type Annotation struct {
    gorm.Model
    Label    string `gorm:"type:varchar(255);"`
    Note     string `gorm:"type:text;"`
}

// AnnotationService handles business logic for annotations.
type AnnotationService struct {
    db *gorm.DB
}

// NewAnnotationService creates a new AnnotationService instance.
func NewAnnotationService(db *gorm.DB) *AnnotationService {
    return &AnnotationService{db: db}
}

// CreateAnnotation adds a new annotation to the database.
func (s *AnnotationService) CreateAnnotation(annotation *Annotation) error {
    result := s.db.Create(&annotation)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

// GetAllAnnotations retrieves all annotations from the database.
func (s *AnnotationService) GetAllAnnotations() ([]Annotation, error) {
    var annotations []Annotation
    result := s.db.Find(&annotations)
    if result.Error != nil {
        return nil, result.Error
    }
    return annotations, nil
}

func main() {
    // Initialize a new SQLite database.
    db, err := gorm.Open(sqlite.Open("data_annotation.db"), &gorm.Config{})
    if err != nil {
        fmt.Println("failed to connect database:", err)
        return
    }

    // Migrate the schema.
    db.AutoMigrate(&Annotation{})

    // Create a new annotation service instance.
    annotationService := NewAnnotationService(db)

    // Example usage: creating and retrieving annotations.
    if err := annotationService.CreateAnnotation(&Annotation{Label: "Example", Note: "This is an example annotation."}); err != nil {
        fmt.Println("failed to create annotation:", err)
    }

    annotations, err := annotationService.GetAllAnnotations()
    if err != nil {
        fmt.Println("failed to retrieve annotations:", err)
    } else {
        for _, annotation := range annotations {
            fmt.Printf("Label: %s, Note: %s
", annotation.Label, annotation.Note)
        }
    }
}