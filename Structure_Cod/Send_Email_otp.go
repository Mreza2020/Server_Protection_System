package Server_Protection_System

import (
	"bytes"
	"fmt"
	"strconv"
	"text/template"

	"gopkg.in/gomail.v2"
)

// # Send_Email
//
// This function is for sending emails [otp].
//
// # Parameters:
//
// - ch chan string
//
// - name_file_Email string for Example Email_sign_in.html
//
// - otp_cod string for Example 123456
//
// - SetHeader_server for Example Company name
//
// - SetHeader_Client for Example Recipient's email
//
// # NOTE:
//
// This code is only for the OTP system.
//
// It gives an output to your html code which is for displaying the code.
//
// # Returns:
//
// It returns three states:
//
// 1- err in template execution
//
// 2- Failed to send email
//
// 3- Email sent successfully
func Send_Email_otp(ch chan string, name_file_Email string, otp_cod string, SetHeader_server string, SetHeader_Client string) {
	data := struct {
		Code string
	}{
		Code: otp_cod,
	}

	htmlData, _ := template.ParseFiles(name_file_Email)

	var htmlContent bytes.Buffer
	if err := htmlData.Execute(&htmlContent, data); err != nil {
		ch <- "err in template execution"
	}

	m := gomail.NewMessage()
	Email := Env_password_Loaded_string("Email")
	Email_pass := Env_password_Loaded_string("Email_pass")
	Email_smtp_server := Env_password_Loaded_string("Email_smtp_server")
	Email_smtp_server_port := Env_password_Loaded_string("Email_smtp_server_port")
	Port, _ := strconv.Atoi(Email_smtp_server_port) // Convert string to int

	m.SetHeader("From", m.FormatAddress(Email, SetHeader_server))

	m.SetHeader("To", SetHeader_Client)

	m.SetHeader("Subject", "احراز هویت")
	emailBody := fmt.Sprintf("کد احراز هویت شما {%s}", otp_cod)
	m.SetBody("text/plain", emailBody)

	m.AddAlternative("text/html", htmlContent.String())

	d := gomail.NewDialer(Email_smtp_server, Port, Email, Email_pass)

	if err := d.DialAndSend(m); err != nil {
		ch <- "Failed to send email"
	}
	ch <- "Email sent successfully"
}
