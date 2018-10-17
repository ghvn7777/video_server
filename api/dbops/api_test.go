package dbops

import (
	"testing"
)

// init(dblogin, truncate tables) -> run tests -> clear data(truncate tables)

func clearTables() {
	dbConn.Exec("truncate users")  // 清空表
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Del", testDeleteUser)
	t.Run("Reget", testRegetUser)
}

func testAddUser(t *testing.T) {
	err := AddUserCredential("kaka", "123")
	if err != nil {
		t.Errorf("Error of AddUser : %v\n", err)
	}
}

func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("kaka")
	if pwd != "123" || err != nil {
		t.Errorf("Error of GetUser: %v", err)
	}
}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("kaka", "123")
	if err != nil {
		t.Errorf("Error of DeleteUser: %v", err)
	}
}

func testRegetUser(t *testing.T) {
	pwd, err := GetUserCredential("kaka")
	if err != nil {
		t.Errorf("Error of RegeUser: %v", err)
	}

	if pwd != "" {
		t.Errorf("Deleting user test failed: %v", err)
	}
}