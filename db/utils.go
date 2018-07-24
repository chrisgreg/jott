package db

import (
	"database/sql"
	"fmt"
)

func init() {
	fmt.Println("Imported DB Utils")
}

// CreateNewBlogPost Insert new blog post for user
func CreateNewBlogPost(db *sql.DB, userID int, title string, subtitle string) (bool, error) {
	stmtIns, err := db.Prepare("INSERT INTO Blog VALUES(null, ?, ?, ?, null)")

	if err != nil {
		return false, err
	}

	defer stmtIns.Close()

	_, err = stmtIns.Exec(userID, title, subtitle)
	if err != nil {
		return false, err
	}

	return true, nil
}
