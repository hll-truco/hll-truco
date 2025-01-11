package utils

import "math"

func equalizeMaps(p map[int]int, q map[int]int) {
	for key := range p {
		if _, ok := q[key]; !ok {
			q[key] = 0
		}
	}
	for key := range q {
		if _, ok := p[key]; !ok {
			p[key] = 0
		}
	}
}

func KL(p map[int]int, q map[int]int) float64 {
	// equalize maps
	equalizeMaps(p, q)

	// Calculate total counts for normalization
	pTotal := 0
	qTotal := 0
	for _, count := range p {
		pTotal += count
	}
	for _, count := range q {
		qTotal += count
	}

	// Initialize KL divergence
	kl := 0.0

	// Calculate KL divergence for each key in p
	for key, pCount := range p {
		// Skip if p(x) = 0 since 0 * log(0/q) = 0
		if pCount == 0 {
			continue
		}

		// Convert counts to probabilities
		pProb := float64(pCount) / float64(pTotal)

		// Get corresponding q count (0 if key doesn't exist in q)
		qCount := q[key]

		// Handle the case where q(x) = 0
		if qCount == 0 {
			// KL divergence is undefined when q(x) = 0 and p(x) > 0
			// Return positive infinity
			return math.Inf(1)
		}

		qProb := float64(qCount) / float64(qTotal)

		// Add to KL divergence: p(x) * log(p(x)/q(x))
		kl += pProb * math.Log(pProb/qProb)
	}

	return kl
}
