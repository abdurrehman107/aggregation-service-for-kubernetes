package models

type NodeHealth struct {
    NodeName      string
    Healthy       bool
    UnhealthyPods []string
}

type PodSummary struct {
    Namespace   string
    Running    int
    Pending    int
    Terminated int
}
