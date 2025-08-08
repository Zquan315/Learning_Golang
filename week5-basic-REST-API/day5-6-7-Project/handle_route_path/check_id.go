package handle_route_path


func CheckID(id string) bool {
	for _, v := range products {
		if id == v.ID {
			return true
		}
	}
	return false
}