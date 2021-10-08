package template

import "fmt"

type email struct{}

func (e *email) randomOTP(len int) string {
	fmt.Printf("EMAIL: generating OTP of length %v", len)
	return "1234"
}

func (e *email) saveOTP(otp string) {
	fmt.Printf("EMAIL: saving OTP... %s", otp)
}

func (e *email) getMessage(otp string) string {
	return "EMAIL OTP for login: " + otp
}

func (e *email) sentNotification(msg string) error {
	fmt.Printf("EMAIL: sending notification... %s", msg)
	return nil
}

func (e *email) publishMetric() {
	fmt.Println("EMAIL: pushing metric...")
}
