package conf

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Worker struct {
	ID         string
	Secret     string
	Alias      string
	CreateTime metav1.Time
	Agents     []WorkerAgent
	Visitors   []WorkerVisitor
}

type Session struct {
	ID        string
	User      string
	Alias     string
	Kind      string
	OS        string
	Arch      string
	Kernel    string
	HostName  string
	Variant   string
	LoginTime metav1.Time
}

type WorkerAgent struct {
	Name       string
	Alias      string
	User       string
	CreateTime metav1.Time
	Services   []WorkerAgentService
}

type WorkerAgentService struct {
	Name       string
	Type       string
	SK         string
	Local_IP   string
	Local_Port int32
	CreateTime metav1.Time
}

type WorkerVisitor struct {
	Name       string
	Alias      string
	User       string
	CreateTime metav1.Time
	Services   []WorkerVisitorService
}

type WorkerVisitorService struct {
	Name        string
	Type        string
	Role        string
	Server_Name string
	SK          string
	Bind_Addr   string
	Bind_Port   int32
	CreateTime  metav1.Time
}
