package reader

type EtcdReader struct {

}

func (this *EtcdReader) Name() string {
	panic("implement me")
}

func (this *EtcdReader) Read() error {
	panic("implement me")
}

func (this *EtcdReader) GetWatchFunc() (WatchFunc, error) {
	panic("implement me")
}

func (this *EtcdReader) Serialize() error {
	panic("implement me")
}