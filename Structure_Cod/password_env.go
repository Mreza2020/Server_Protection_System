package Server_Protection_System

import "os"

// # Env_password_Loaded_string
//
// # Parameters:
//
// - This function receives the password from env, so its inputs are based on the parameters specified beforehand.
//
// # Example:
//
// High speed in receiving passwords and receiving passwords with high security
func Env_password_Loaded_string(Type string) string {
	dbPassword := os.Getenv("DB_PASSWORD")
	Redis_Password_db := os.Getenv("DB_Redis")
	Redis_Password := os.Getenv("DB_Redis_password")
	Email := os.Getenv("Email")
	Email_pass := os.Getenv("Email_pass")
	Email_smtp_server := os.Getenv("Email_smtp_server")           //smtp.gmail.com
	Email_smtp_server_port := os.Getenv("Email_smtp_server_port") //465 ssl   in 587 starttls

	switch Type {
	case "dbPassword":
		return dbPassword
	case "DB_Redis":
		return Redis_Password_db
	case "DB_Redis_password":
		return Redis_Password
	case "Email":
		return Email
	case "Email_pass":
		return Email_pass
	case "Email_smtp_server":
		return Email_smtp_server
	case "Email_smtp_server_port":
		return Email_smtp_server_port

	default:
		return ""
	}

}
