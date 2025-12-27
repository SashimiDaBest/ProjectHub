package handlers

import (
    "context"
    // "fmt"
    "projecthub-backend/db"
    "projecthub-backend/models"
)

func CreateUser(user models.User) error {
    _, err := db.Pool.Exec(context.Background(),
        "INSERT INTO users (name, email) VALUES ($1, $2)",
        user.Name, user.Email)
    return err
}

func GetUsers() ([]models.User, error) {
    rows, err := db.Pool.Query(context.Background(),
        "SELECT id, name, email FROM users")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []models.User
    for rows.Next() {
        var u models.User
        if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
            return nil, err
        }
        users = append(users, u)
    }
    return users, nil
}
