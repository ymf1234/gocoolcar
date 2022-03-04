package id

type AccountID string

func (a AccountID) String() string {
	return string(a)
}

// TripID defines trip id object.
type TripID string

func (t TripID) String() string {
	return string(t)
}

type IdentityID string

func (i IdentityID) String() string {
	return string(i)
}

type CarID string

func (c CarID) String() string {
	return string(c)
}

type BlobID string

func (b BlobID) String() string {
	return string(b)
}
