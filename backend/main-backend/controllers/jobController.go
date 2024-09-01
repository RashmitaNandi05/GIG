package controllers

import (
    "context"
    "net/http"
    "time"
    "github.com/ritankarsaha/backend/database"
    "github.com/ritankarsaha/backend/models"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateJob(c *gin.Context) {
    var job models.Job
    if err := c.BindJSON(&job); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    job.ID = primitive.NewObjectID()
    job.EmployerID, _ = primitive.ObjectIDFromHex(c.GetString("user_id"))
    job.CreatedAt = time.Now().Unix()

    collection := database.GetCollection("jobs")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    _, err := collection.InsertOne(ctx, job)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating job"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Job created successfully"})
}

func GetJobs(c *gin.Context) {
    var jobs []models.Job
    collection := database.GetCollection("jobs")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    cursor, err := collection.Find(ctx, bson.M{})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching jobs"})
        return
    }

    if err := cursor.All(ctx, &jobs); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing jobs"})
        return
    }

    c.JSON(http.StatusOK, jobs)
}

func ApplyJob(c *gin.Context) {
    // jobID := c.Param("jobId")
    userID, _ := primitive.ObjectIDFromHex(c.GetString("user_id"))

    application := models.Application{
        ID:        primitive.NewObjectID(),
        // JobID:     primitive.ObjectIDFromHex(jobID),
        EmployeeID: userID,
        AppliedAt: time.Now().Unix(),
    }

    collection := database.GetCollection("applications")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    _, err := collection.InsertOne(ctx, application)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error applying to job"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Applied to job successfully"})
}