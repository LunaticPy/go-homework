package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	// type=nil, val=nil != type=string val=""
	if err != nil {
		println("error")
		return
	}
	println("ok")
}

// error
