package main

// db 관련 구조체
type C_DB struct {
	s_id   string
	s_pwd  string
	s_ip   string
	s_port string
}

func (t *C_DB) Insert() {

	t.db_insert()
}

// 전역 변수로 올릴거
// func Insert

// db 관련 함수를 지역변수로 쭉 나열.
func (t *C_DB) db_insert() {

	// db insert
	// db ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
}

func (t *C_DB) db_delete() {

	// db insert
	// db ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
}
