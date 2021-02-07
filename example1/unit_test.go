package example1

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Do not use `panic` for testing
// Fail -> Continue to the next code
// FailNow -> Continue to the next test
// Error -> Error == Fail + Comment
// Fatal -> Fatal == FailNow + Comment
// Assert -> Assert == Error + detail from package
// Require -> Require == Fatal + detail from package

// Running test with `go test` + `-v` for detail
// Running test on all package use `go test ./..` + `-v` for detail

//! For Before and After Test -> `TestMain` is a Key
//! `TestMain` is just for only one package
// func TestMain(m *testing.M) {
// 	fmt.Println("Database Opened!")

// 	m.Run()

// 	fmt.Println("Database Closed!")
// }

//! `Benchmark` is a key
func BenchmarkTestHalo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Halo("Nana")
	}
}

func BenchmarkSubTest(b *testing.B) {
	b.Run("Halo", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Halo("Nana")
		}
	})
	b.Run("Halo2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Halo("Nana2")
		}
	})
}

func BenchmarkTableTest(b *testing.B) {
	//Todo: Can use table too, jsut tricky code!
}

//! Do not use return value for unit testing
//! `Test{function name}` key for unit testing
func TestHalo(t *testing.T) {
	result := Halo("Nana")

	if result != "Hello Nana!" {
		t.Error("Error!")
	}
}

func TestHaloAssert(t *testing.T) {
	result := Halo("Nana")

	if result != "Hello Nana!" {
		assert.Equal(t, "Hello Nana!", result, "Assert error!")
		// require.Equal(t, "Hello Nana!", result, "Require error!")
	}
}

func TestHaloSkip(t *testing.T) {
	result := Halo("Nana")

	if runtime.GOOS == "windows" {
		t.Skip("Can't run on windows OS")
	} else {
		assert.Equal(t, "Hello Nana!", result, "Skipping error!")
	}
}

//! Sub Test
func TestSubTest(t *testing.T) {
	t.Run("Test1", func(t *testing.T) {
		result := Halo("Test1")
		require.Equal(t, "Hello Test1!", result, "Sub Test1 error!")
	})
	t.Run("Test2", func(t *testing.T) {
		result := Halo("Test2")
		require.Equal(t, "Hello Test2!", result, "Sub Test2 error!")
	})
}

//! Table Test -> for easy subtest configuration
func TestTableTest(t *testing.T) {

	type test struct {
		FunctionName string
		Expected     interface{}
		Actual       interface{}
		Message      interface{}
	}

	sliceTest := []test{}
	sliceTest = append(sliceTest, test{FunctionName: "Test1", Expected: "Hello Test1!", Actual: Halo("Test1"), Message: "Error"})
	sliceTest = append(sliceTest, test{FunctionName: "Test2", Expected: "Hello Test2!", Actual: Halo("Test2"), Message: "Error"})
	sliceTest = append(sliceTest, test{FunctionName: "Test3", Expected: "Hello Test3!", Actual: Halo("Test3"), Message: "Error"})
	sliceTest = append(sliceTest, test{FunctionName: "Test4", Expected: "Hello Test4!", Actual: Halo("Test4"), Message: "Error"})

	for _, value := range sliceTest {
		t.Run(value.FunctionName, func(t *testing.T) {
			require.Equal(t, value.Expected, value.Actual, value.Message)
		})
	}
}
