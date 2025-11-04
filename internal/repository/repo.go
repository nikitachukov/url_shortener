package repository

type ShortenerRepo interface {
	FindShortURL(long string) (string, bool)
	GetLongURL(short string) (string, bool)
	Set(short, long string) bool
	Ping() bool
}
