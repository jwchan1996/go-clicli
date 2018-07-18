package dbOpt

import "testing"

func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate posts")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()

}

func TestUserWorkFlow(t *testing.T) {
	t.Run("Create", testCreateUser)
	t.Run("Get", testGetUser)
	t.Run("Delete", testDeleteUser)
	t.Run("Reget", testRegetUser)
}

func testCreateUser(t *testing.T) {
	err := CreateUserCredential("admin", "admin", "user", 1533540012)
	if err != nil {
		t.Errorf("Err of CreateUser:%v", err)
	}
}

func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("admin")
	if pwd != "admin" {
		t.Errorf("%v", pwd)
	}
	if err != nil {
		t.Errorf("Err of GetUser:%v", err)
	}
}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("admin", "admin")
	if err != nil {
		t.Errorf("Err of DeleteUser:%v", err)
	}
}

func testRegetUser(t *testing.T) {
	pwd, err := GetUserCredential("admin")
	if err != nil {
		t.Errorf("Err of regetUser:%v", err)
	}
	if pwd != "" {
		t.Errorf("Delete user test failed")
	}
}
