package channels

//OK Controla la salida de una go func del repository_user_crud
func OK(done <-chan bool) bool {
	select {
	case ok := <-done:
		if ok {
			return true
		}
	}
	return false
}
