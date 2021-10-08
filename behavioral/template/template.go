package template

// it allows to define a skeleton of an algorithm or operation comprised of different steps
// some steps can be implemented in different ways, so it lets you to override those without changing the overall algorithm’s structure

// similar to the strategy pattern, but only with some part of the algorithm

// example -> OTP = One Time Password
// steps:
// 1- Generate a random n digit number
// 2- Save this number in the cache for later verification
// 3- Prepare the content
// 4- Send the notification
// 5- Publish the metrics

type otpTemplate interface {
	randomOTP(int) string
	saveOTP(string)
	getMessage(string) string
	sentNotification(string) error
	publishMetric()
}

type otp struct {
	template otpTemplate
}

// generateOTP is the template method, the algorithm is common for every otp
func (o *otp) generateOTP(otpLength int) error {
	s := o.template.randomOTP(otpLength)
	o.template.saveOTP(s)
	message := o.template.getMessage(s)
	if err := o.template.sentNotification(message); err != nil {
		return err
	}
	o.template.publishMetric()
	return nil
}

// otpTemplate -> defines the set of methods that any otp type should implement
// sms and email -> are the implementors of otpTemplate interface
// otp -> defines the template method generateOTP(), otp embeds otpTemplate interface.
// The combination of otpTemplate interface and otp struct provides the capabilities of Abstract Class in go