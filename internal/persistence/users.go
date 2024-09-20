package persistence

type User struct {
	ID       int
	Username string
	Email    string
}

func createUserTable() error {
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL,
		email TEXT NOT NULL
	);`
	_, err := db.Exec(createTableQuery)
	return err
}

func GetAllUsers() ([]User, error) {
	query := "SELECT id, username, email FROM users"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func CreateUser(user User) (User, error) {
	stmt, err := db.Prepare("INSERT INTO users (username, email) VALUES (?, ?);")
	if err != nil {
		return User{}, err
	}
	defer stmt.Close()

	r, err := stmt.Exec(user.Username, user.Email)
	if err != nil {
		return User{}, err
	}

	id, _ := r.LastInsertId()
	user.ID = int(id)
	return user, nil
}

func DeleteUser(userId int) error {
	stmt, err := db.Prepare("DELETE FROM users WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userId)
	return err
}
