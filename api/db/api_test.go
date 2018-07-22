package db

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
	err := CreateUser("admin", "admin", "user", 1533540012,"")
	if err != nil {
		t.Errorf("Err of CreateUser:%v", err)
	}
}

func testGetUser(t *testing.T) {
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
	pwd, err := GetUser("admin")
	if err != nil {
		t.Errorf("Err of regetUser:%v", err)
	}
	if pwd != "" {
		t.Errorf("Delete user test failed")
	}
}

func TestPostWorkFlow(t *testing.T) {
	t.Run("Create", testAddPost)
	t.Run("Get", testGetPost)
	t.Run("Delete", testDeletePost)
	t.Run("Reget", testRegetPost)
}

func testAddPost(t *testing.T) {
	_, err := AddPost("title", "content", "anime", "publish")
	if err != nil {
		t.Errorf("%v", err)
	}
}

func testGetPost(t *testing.T) {
	_, err := GetPost(1)
	if err != nil {
		t.Errorf("%v", err)
	}
}


func testDeletePost(t *testing.T) {
	err := DeletePost(1)
	if err != nil {
		t.Errorf("%v", err)
	}
}

func testRegetPost(t *testing.T) {
	_, err := GetPost(1)
	if err != nil {
		t.Errorf("Err of regetUser:%v", err)
	}
}