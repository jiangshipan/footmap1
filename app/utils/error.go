package utils

func CheckError(err error) {
	//file, err1 := os.Open("logs/log.log")
	//if err1 != nil {
	//	log.Fatal("open log file error")
	//}
	//defer file.Close()
	//if err != nil {
	//	log.New(file, "[error]", log.Llongfile)
	//}
	if err != nil {
		panic(err)
	}
}