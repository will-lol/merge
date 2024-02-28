package merge

type TwMerge interface {
	Merge(existing string, incoming string) (*string, error)
	Close()
}
