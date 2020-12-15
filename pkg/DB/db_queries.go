package DB

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db_type = "mysql"
	db_user = os.Getenv("DB_USER")
	db_pwd  = os.Getenv("DB_PASS")
	db_addr = os.Getenv("DB_ADDR")
	db_name = os.Getenv("DB_NAME")

	// conn_string = db_user + ":" + db_pwd + "@(" + db_addr + ")/" + db_name + "?parseTime=true"
	conn_string = db_user + ":" + db_pwd + "@tcp(" + db_addr + ")/" + db_name //+ "?parseTime=true"
)

func connect_to_DB() *sql.DB {
	fmt.Printf("Connection string %v \n", conn_string)
	conn, err := sql.Open(db_type, conn_string)
	if err != nil {
		fmt.Println("Error connecting to DB, ", err)
	}
	return conn
}
func Check_ID_Unique(id string) bool {
	conn := connect_to_DB()
	defer conn.Close()
	// var id_unique bool
	// err := conn.QueryRow("call spUserIdTaken(?)", id).Scan(&id_unique)
	insert, err := conn.Query("INSERT INTO " + db_name + " VALUES (1,USR_1)")
	defer insert.Close()
	if err != nil {
		fmt.Println(err)
		return false
	}
	// return id_unique
	return true
}

// port : 3306
// username : root
// root@localhost:3306
// cry_root@127.0.0.1:3306
//jdbc:mysql://127.0.0.1:3306/?user=cry_root

/*
Stored procedure should look like this
PROCEDURE `spUserIdTaken`(IN p_id varchar(150))
BEGIN
    SELECT count(user_id)
    FROM User
    WHERE user_id = p_id
END

*/
