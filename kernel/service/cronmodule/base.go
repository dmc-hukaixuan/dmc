package cronmodule

type CornModule interface {
    Run(taskID int, taskData string)
}

func CornRun(CornType string) CornModule {
    switch CornType {
    case "auto_report":
        return &Report{}
    case "cacheDelete":
        return &Report{}
    default:
        return &Report{}
    }
}
