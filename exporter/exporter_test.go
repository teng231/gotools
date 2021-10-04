package exporter

import (
	"log"
	"testing"
	"time"
)

type User struct {
	Fullname string
	Id       string
	Birth    int64
	Phone    string
}

func Test_exporter(t *testing.T) {
	header := []string{"Id", "Tên", "Ngày sinh"}
	data := make([]*User, 0)
	data = append(data,
		&User{Id: "i1", Fullname: "te1", Birth: 24234234},
		&User{Id: "i2", Fullname: "te2", Birth: 2432234234})
	inExport := [][]interface{}{}
	for _, d := range data {
		inExport = append(inExport, []interface{}{
			d.Id, d.Fullname, time.Unix(d.Birth, 0).String(),
		})
	}
	exportCSV, err := Exporter(header, inExport)
	log.Print(exportCSV, err)
	MakeCsvFile("afdasfsadf", exportCSV)
}
func Test_exporter2(t *testing.T) {
	// u2 := User{Fullname: "te1", Phone: "03356575675"}
	// bin, err := proto.Marshal(u2.(proto.Message))
}
