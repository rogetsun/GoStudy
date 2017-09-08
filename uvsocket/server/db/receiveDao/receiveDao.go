package receiveDao

import (
	"GoStudy/uvsocket/server/db"
	"fmt"
)

func Insert(frames [][]byte) (rows int, err error) {
	fmt.Println("save:", frames)
	rows = 0
	tx, err := db.Conn.Begin()
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer func(err error) {
		if err != nil {
			fmt.Println("rollback")
			tx.Rollback()
		} else {
			tx.Commit()
			fmt.Println("commit")
		}
	}(err)

	statm, err := tx.Prepare("insert into d_receive_cmd_msg values('1','2','3','4',5,6,current_timestamp(),?)")
	defer statm.Close()
	if err != nil {
		return 0, err
	}

	for _, frame := range frames {
		_, err := statm.Exec(fmt.Sprintf("%X", frame))
		if err != nil {
			return 0, err
		}
	}
	return len(frames), nil
}
