package utils

import (
	"math"
	"reflect"
	"testing"

	. "github.com/agiledragon/gomonkey"
	. "github.com/prashantv/gostub"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"golang-project-template/pkg/client"
	mock_client "golang-project-template/test/mock/client"

	. "github.com/golang/mock/gomock"
)

// TestSum1 testing unittest
func TestSum1(t *testing.T) {
	actual := Sum(10)
	assert.Equal(t, actual, 45, "should equal")
	require.Equal(t, actual, 45, "should equal")
	t.Log("TestSum1 Finished")
}

// TestSum2 GoConvey unittest
func TestSum2(t *testing.T) {
	actual := Sum(10)
	Convey("GoConvey Example", t, func() {
		So(actual, ShouldEqual, 45)
	})
	t.Log("TestSum2 Finished")
}

// TestSqrt1 GoStub global variable
func TestSqrt1(t *testing.T) {
	stubs := Stub(&MUL, float64(2))
	defer stubs.Reset()
	actual := Sqrt(float64(4.5))
	assert.Equal(t, actual, float64(3), "should equal")
	t.Log("TestSqrt1 Finished")
}

// TestSqrt2 GoStub function
func TestSqrt2(t *testing.T) {
	Exec := math.Sqrt
	stubs := StubFunc(&Exec, float64(3))
	defer stubs.Reset()
	actual := Exec(float64(4))
	assert.Equal(t, actual, float64(3), "should equal")
	t.Log("TestSqrt1 Finished")
}

// TestSqrt3 GoMonkey function
func TestSqrt3(t *testing.T) {
	Exec := math.Sqrt
	guard := ApplyFunc(
		Exec,
		func(_ float64) float64 {
			return float64(3)
		})
	defer guard.Reset()
	actual := Exec(float64(4))
	assert.Equal(t, actual, float64(3), "should equal")
	t.Log("TestSqrt3 Finished")
}

// TestShow1 GoMock method
func TestShow1(t *testing.T) {
	//创建mock控制器
	//定义mock对象的作用域和生命周期
	ctrl := NewController(t)
	defer ctrl.Finish()

	//mock对象注入控制器
	//多个mock对象则注入同一个控制器
	mockClient := mock_client.NewMockTestClientInterface(ctrl)

	//mock对象行为定义
	mockClient.EXPECT().Get().Return("test")

	// InOrder(
	// 	mockClient.EXPECT().Get().Return("test1"),
	// 	mockClient.EXPECT().Get().Return("test2"),
	// 	mockClient.EXPECT().Get().Return("test3"),
	// )

	actual := Show(mockClient)
	assert.Equal(t, actual, "test", "should equal")
	t.Log("TestShow1 Finished")
}

// TestModify1 GoMock method
func TestModify1(t *testing.T) {
	//创建mock控制器
	//定义mock对象的作用域和生命周期
	ctrl := NewController(t)
	defer ctrl.Finish()

	//mock对象注入控制器
	//多个mock对象则注入同一个控制器
	mockClient := mock_client.NewMockTestClientInterface(ctrl)

	//mock对象行为定义
	mockClient.EXPECT().Update("name").Return("testname")

	actual := Modify(mockClient, "name")
	assert.Equal(t, actual, "testname", "should equal")
	t.Log("TestModify1 Finished")
}

// TestModify2 GoMonkey method
func TestModify2(t *testing.T) {
	var mock_client *client.Client
	guard := ApplyMethod(reflect.TypeOf(mock_client), "Update", func(_ *client.Client, _ string) string {
		return "testname"
	})
	defer guard.Reset()
	c := client.NewClient()
	actual := Modify(c, "name")
	assert.Equal(t, actual, "testname", "should equal")
	t.Log("TestModify2 Finished")
}

//
func TestIsEqual(t *testing.T) {
	guard := ApplyFunc(IsEqual, func(_, _ string) bool {
		return true
	})
	defer guard.Reset()
	actual := IsEqual("hello", "world")
	assert.Equal(t, actual, false, "should equal")
	t.Log("TestModify2 Finished")
}
