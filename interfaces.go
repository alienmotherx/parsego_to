package exporter

// SocialMedia : Social Media Interface
type SocialMedia interface {
	Feed() []string
	Fame() int
}
