package assert

type testcommon interface {
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fail()
	FailNow()
	Failed() bool
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Helper()
	Log(args ...interface{})
	Logf(format string, args ...interface{})
	Name() string
	Skip(args ...interface{})
	SkipNow()
	Skipf(format string, args ...interface{})
	Skipped() bool
}

func AssertEqual(t testcommon, a, b interface{}) {
	// 标记调用函数为helper函数，打印文件信息或日志，不会追溯这个函数
	t.Helper()
	if a != b {
		t.Errorf("Not Equal. %d %d", a, b)
	}
}
