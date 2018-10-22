package taskrunner

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"video_server/scheduler/dbops"
)

func VideoClearDispatcher(dc dataChan) error {
	res, err := dbops.ReadVideoDeletionRecord(3)
	if err != nil {
		log.Printf("Video clear dispatcher error: %v", err)
		return err
	}

	if len(res) == 0 {
		return errors.New("All tasks finished")
	}

	for _, id := range res {
		fmt.Printf("add %v to dc chan\n", id)
		dc <- id
	}
	return nil
}

func deleteVideo(vid string) error {
	path, _ := filepath.Abs(VIDEO_DIR + vid)
	log.Println(path)
	err := os.Remove(path)
	if err != nil && !os.IsNotExist(err) { // 文件不存在是正常的，已经删除过了
		log.Printf("Deleting video error: %v", err)
		return err
	}
	return nil
}
func VideoClearExecutor(dc dataChan) error {
	errMap := &sync.Map{}
	var err error

	flag := make(chan int, 10) // 最多 10 个 go 协程同时跑
	forloop:
		for {
			select {
			case vid := <-dc:
				flag <- 0
				go func(id interface{}) error {
					if err := deleteVideo(id.(string)); err != nil {
						_ = <-flag
						errMap.Store(id, err)
						return err
					}
					if err := dbops.DelVideoDeletionRecord(id.(string)); err != nil {
						_ = <-flag
						errMap.Store(id, err)
						return err
					}
					_ = <-flag
					return nil
				}(vid)
			default:
				break forloop
			}
		}
	<-flag // 等待所有协程结束
	errMap.Range(func(k, v interface{}) bool {
		err = v.(error)
		if err != nil {
			return false
		}
		return true
	})
	return err
}