package golang

import (
	"testing"
)

func TestCaptcha(t *testing.T) {
	c := []int{1, 1, 1, 1}

	rs := getCaptcha(c)
	if rs != "1 + one" {
		t.Error("it should be 1 + one")
	}

}

func TestCaptcha2(t *testing.T) {
	c := []int{2, 1, 1, 1}

	rs := getCaptcha(c)

	if rs != "one + 1" {
		t.Error("it should be one + 1")
	}

}
func TestCaptcha3(t *testing.T) {
	c := []int{1, 1, 2, 1}

	rs := getCaptcha(c)

	if rs != "1 - one" {
		t.Error("it should be 1 - one")
	}

}
func TestCaptcha4(t *testing.T) {
	c := []int{2, 2, 2, 1}

	rs := getCaptcha(c)
	if rs != "two - 1" {
		t.Error("it should be two - 1")
	}

}
func TestCaptcha5(t *testing.T) {
	c := []int{1, 1, 3, 1}

	rs := getCaptcha(c)

	if rs != "1 * one" {
		t.Error("it should be 1 * one")
	}

}
func TestCaptcha6(t *testing.T) {
	c := []int{2, 2, 3, 1}

	rs := getCaptcha(c)
	if rs != "two * 1" {
		t.Error("it should be two * 1")
	}

}
func TestCaptcha7(t *testing.T) {
	c := []int{2, 1, 3, 1}

	rs := getCaptcha(c)

	if rs != "one * 1" {
		t.Error("it should be one * 1")
	}

}
func TestCaptcha8(t *testing.T) {
	c := []int{1, 2, 3, 1}

	rs := getCaptcha(c)
	if rs != "2 * one" {
		t.Error("it should be 2 * one")
	}

}
