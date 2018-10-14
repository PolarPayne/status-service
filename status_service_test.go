package main

import (
	"net/http"
	"os"
	"strconv"
	"testing"
)

func clearEnv(t *testing.T) {
	envs := []string{"STATUS_CODE", "STATUS_MESSAGE", "HOST", "PORT"}

	for _, v := range envs {
		err := os.Unsetenv(v)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestGetStatusCodeDefault(t *testing.T) {
	clearEnv(t)

	v := getStatusCode()
	if v != 200 {
		t.Fail()
	}
}

func TestGetStatusMessageDefault(t *testing.T) {
	clearEnv(t)

	v := getStatusMessage(getStatusCode())
	if v != "OK" {
		t.Fail()
	}
}

func TestGetHostDefault(t *testing.T) {
	clearEnv(t)

	v := getHost()
	if v != "0.0.0.0" {
		t.Fail()
	}
}

func TestGetPortDefault(t *testing.T) {
	clearEnv(t)

	v := getPort()
	if v != 80 {
		t.Fail()
	}
}

func TestGetStatusCodeAllValid(t *testing.T) {
	clearEnv(t)

	test := func(i int) {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			os.Setenv("STATUS_CODE", strconv.Itoa(i))
			getStatusCode()
		})
	}

	for i := 200; i <= 299; i++ {
		test(i)
	}

	for i := 400; i <= 499; i++ {
		test(i)
	}

	for i := 500; i <= 599; i++ {
		test(i)
	}
}

func TestGetStatusCodeInvalidInteger(t *testing.T) {
	clearEnv(t)

	test := func(i int) {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Fail()
				}
			}()

			os.Setenv("STATUS_CODE", strconv.Itoa(i))
			getStatusCode()
		})
	}

	for i := 0; i <= 199; i++ {
		test(i)
	}

	for i := 300; i <= 399; i++ {
		test(i)
	}

	for i := 600; i <= 999; i++ {
		test(i)
	}

	// just straight up invalid status codes
	for i := -999; i <= -1; i++ {
		test(i)
	}

	for i := 1000; i <= 1999; i++ {
		test(i)
	}
}

func TestGetStatusCodeInvalid(t *testing.T) {
	clearEnv(t)

	test := func(v string) {
		t.Run(v, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Fail()
				}
			}()

			os.Setenv("STATUS_CODE", v)
			getStatusCode()
		})
	}

	testCases := []string{"", "ok", "200.0"}

	for _, v := range testCases {
		test(v)
	}
}

func TestGetStatusMessageDefaults(t *testing.T) {
	clearEnv(t)

	test := func(k int, v string) {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			if v != getStatusMessage(k) {
				t.Fail()
			}
		})
	}

	testCases := make(map[int]string)

	for i := 100; i <= 999; i++ {
		testCases[i] = http.StatusText(i)
	}

	for k, v := range testCases {
		test(k, v)
	}
}

func TestGetStatusMessageCustom(t *testing.T) {
	clearEnv(t)

	test := func(v string) {
		t.Run(v, func(t *testing.T) {
			os.Setenv("STATUS_MESSAGE", v)
			if v != getStatusMessage(200) {
				t.Fail()
			}
		})
	}

	testCases := []string{"", " ", "ok", "hello world", "0"}

	for _, v := range testCases {
		test(v)
	}
}
