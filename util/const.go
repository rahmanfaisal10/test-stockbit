package util

const (
	//total worker/pekerja
	TOTALWORKER = 100
	//jumlah koneksi idle
	DBMAXIDLECONNS = 4
	//jumlah koneksi database dalam pool || menjaga supaya koneksi database terbuka
	DBMAXCONNS = 100

	CONTENTTYPEJSONANDCHARSETUTF8VALUE = "application/json; charset=utf-8"
	CONTENTTYPEKEY                     = "Content-Type"
)
