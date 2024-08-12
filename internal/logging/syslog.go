package logging

import (
	"fmt"
	"net"
)

type SyslogMessage struct {
	Priority byte
	Version  string
	Hostname string
	AppName  string
	Msg      string
}

func ConnectToSyslogServer(hostPort string) (net.Conn, error) {
	conn, err := net.Dial("udp", hostPort)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to syslog server: %w", err)
	}
	return conn, nil
}

func SendMessage(conn net.Conn, msg SyslogMessage) error {
	// Format the syslog message according to RFC 3164 or RFC 5424
	formattedMsg := fmt.Sprintf("%c %s %s[%d]: %s\n",
		msg.Priority,
		msg.Hostname,
		msg.AppName,
		msg.Version,
		msg.Msg,
	)

	_, err := conn.Write([]byte(formattedMsg))
	if err != nil {
		return fmt.Errorf("failed to send syslog message: %w", err)
	}

	return nil
}
