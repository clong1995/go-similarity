package similarity

import (
	"fmt"
	"math"
)

// NGramFeatures Generate 2-Gram Features
func NGramFeatures(s string, n int) map[string]int {
	features := make(map[string]int)
	runes := []rune(s)

	// Generate n-grams
	for i := 0; i <= len(runes)-n; i++ {
		gram := string(runes[i : i+n])
		features[gram]++
	}
	return features
}

// Cosine Similarity between two feature vectors
func Cosine(featuresA, featuresB map[string]int) float64 {
	// Calculate dot product and magnitudes
	var dotProduct, magnitudeA, magnitudeB float64
	for key, valA := range featuresA {
		valB := featuresB[key]
		dotProduct += float64(valA * valB)
		magnitudeA += float64(valA * valA)
	}
	for _, valB := range featuresB {
		magnitudeB += float64(valB * valB)
	}

	// Avoid division by zero
	if magnitudeA == 0 || magnitudeB == 0 {
		return 0.0
	}

	// Compute cosine similarity
	return dotProduct / (math.Sqrt(magnitudeA) * math.Sqrt(magnitudeB))
}

// Main function to demonstrate saving and comparing features
func main() {
	// Step 1: Save A's features
	strA := "你好世界"
	nGramA := NGramFeatures(strA, 2) // 2-Gram

	// Step 2: Compare with B
	strB := "你好地球"
	nGramB := NGramFeatures(strB, 2) // 2-Gram

	// Compute similarity
	similarity := Cosine(nGramA, nGramB)

	fmt.Printf("Similarity between '%s' and '%s': %.2f%%\n", strA, strB, similarity*100)
}
