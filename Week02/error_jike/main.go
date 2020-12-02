package error_jike

import (
	"database/sql"
	"errors"
	"fmt"
	xerrors "github.com/pkg/errors"
	"os/exec"
)

func main() {
	str := service()
	fmt.Println(str)

}

func service() string {
	str,err := mid()
	if err != nil {
		if errors.As(err,sql.ErrNoRows) {
			return "查不到数据"
		}else {
			fmt.Println("main:%+v\n",err)
			return "其他错误"
		}
	}
	return str
}

func mid() (string,error) {
	return source()
}

func source() (string,error) {
	i,err := exec.LookPath("mysql_shell")
	return i,xerrors.Wrapf(err, "source error")
}