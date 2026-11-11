package main

type envelope[T any] struct {
	recipient string
	payload   T
}

func createEnvelopes[T any](recipients []string, payload T) []envelope[T] {
	envelopes := []envelope[T]{}

	for _, recipient := range recipients {
		envelopes = append(envelopes, envelope[T]{
			recipient: recipient,
			payload:   payload,
		})
	}

	return envelopes
}
