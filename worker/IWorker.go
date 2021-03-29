package worker

//IWorker ...
type IWorker interface {
	Run()
	Push(data *EntryData)
	DeviceExist(string) bool
	NewDevice(string)
}
