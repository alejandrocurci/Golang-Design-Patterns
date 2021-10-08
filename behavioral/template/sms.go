package template

import "fmt"

type sms struct{}

func (s *sms) randomOTP(len int) string {
	fmt.Printf("SMS: generating OTP of length %v", len)
	return "1234"
}

func (s *sms) saveOTP(otp string) {
	fmt.Printf("SMS: saving OTP... %s", otp)
}

func (s *sms) getMessage(otp string) string {
	return "SMS OTP for login: " + otp
}

func (s *sms) sentNotification(msg string) error {
	fmt.Printf("SMS: sending notification... %s", msg)
	return nil
}

func (s *sms) publishMetric() {
	fmt.Println("SMS: pushing metric...")
}
