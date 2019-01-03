package lib

func CheckError(err error) {
    if err != nil {
        PrintError(err.Error())
    }
}
