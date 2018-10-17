package dbops

import (
	"fmt"
	"strconv"
	"testing"
	"time"
	"video_server/api/defs"
)

var temp_v_info defs.VideoInfo

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

func TestVideoWorkFlow(t *testing.T) {
	clearTables()
	t.Run("PrepareUser", testAddUser)
	t.Run("AddVideo", testAddVideoInfo)
	t.Run("GetVideo", testGetVideoInfo)
	t.Run("DelVideo", testDeleteVideoInfo)
	t.Run("RegetVideo", testRegetVideoInfo)
	clearTables()
}

func testAddVideoInfo(t *testing.T) {
	vi, err := AddNewVideo(1, "my-video")
	if err != nil {
		t.Errorf("Error of AddVideoInfo: %v", err)
	}
	temp_v_info = *vi
}

func testGetVideoInfo(t *testing.T) {
	res, err := GetVideoInfo(temp_v_info.Id)
	if err != nil {
		t.Errorf("Error of GetVideoInfo: %v", err)
	}

	if res == nil || res.Id != temp_v_info.Id ||
			res.AuthorId != temp_v_info.AuthorId || res.Name != temp_v_info.Name ||
			res.DisplayCtime != temp_v_info.DisplayCtime {
		t.Errorf("Error of GetVideoInfo: %v", err)
	}
}

func testDeleteVideoInfo(t *testing.T) {
	err := DeleteVideoInfo(temp_v_info.Id)
	if err != nil {
		t.Errorf("Error of DeleteVideoInfo: %v", err)
	}
}

func testRegetVideoInfo(t *testing.T) {
	vi, err := GetVideoInfo(temp_v_info.Id)
	if err != nil || vi != nil {
		fmt.Println("????????????", vi)
		t.Errorf("Error of RegetVideoInfo: %v", err)
	}
}

func TestCommentsWorkFlow(t *testing.T) {
	clearTables()
	t.Run("AddUser", testAddUser)
	t.Run("AddVideo", testAddVideoInfo)
	t.Run("AddComments", testAddComments)
	t.Run("ListComments", testListComments)
	clearTables()
}

func testAddComments(t *testing.T) {
	vid := temp_v_info.Id
	aid := temp_v_info.AuthorId
	content := "I like this video"

	err := AddNewComments(vid, aid, content)
	if err != nil {
		t.Errorf("Error of AddComments: %v", err)
	}
}

func testListComments(t *testing.T) {
	vid := temp_v_info.Id
	from := 1514764800 // 2018/1/1 8:0:0
	to, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano() / 1e9, 10))
	res, err := ListComments(vid, from, to)
	if err != nil {
		t.Errorf("Error of ListComments: %v", err)
	}
	for i, ele := range res {
		fmt.Printf("comment: %d, %+v\n", i, ele)
	}
}