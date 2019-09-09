package helpers

import "github.com/google/uuid"

//request uuid
func GetUUid() string{
	var id uuid.UUID
	var err error

	for i:=0; i<3;i++ {
		id,err = uuid.NewUUID()
		if err == nil{
			return id.String()
		}
	}

	return ""
}