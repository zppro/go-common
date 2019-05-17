package data

type Row interface{
	Get (keyPath string) (interface{}, error)

}
