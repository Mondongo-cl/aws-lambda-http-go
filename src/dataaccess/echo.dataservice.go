package dataaccess

var (
	data       = map[int32]MessageRow{1: {1, "Hard Coded"}}
	next int32 = 1
)

func Add(message string) MessageRow {
	next++
	data[next] = MessageRow{Id: int32(next), Message: message}
	return data[next]

}
func Remove(id int32) {
	delete(data, id)
}

func Find(id int32) string {
	r := data[id]
	return r.Message
}
