package lib

import "github.com/google/uuid"

func ParseUUIDStrToBin(strUUID string) []byte {
	UUID, parseErr := uuid.Parse(strUUID)
	if parseErr != nil {
		return nil
	}
	binaryUUID, err := UUID.MarshalBinary()
	if err != nil {
		return nil
	}
	return binaryUUID
}

func ParseUUIDBinToStr(binaryUUID []byte) string {
	UUID, parseErr := uuid.ParseBytes(binaryUUID)
	if parseErr != nil {
		return parseErr.Error()
	}
	strUUID := UUID.String()

	return strUUID
}
