package main

import (
	"fmt"
	"sort"
)

type BrainrotMeme struct {
	Name       string
	TrendLevel int
	Category   string
	Views      float64
}

func FindTopTrending(memes []BrainrotMeme, minViews float64) []BrainrotMeme {
	var result []BrainrotMeme

	for _, meme := range memes {
		if meme.Views > minViews {
			result = append(result, meme)
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].TrendLevel > result[j].TrendLevel
	})

	return result
}

func CalculateCategoryImpact(memes []BrainrotMeme) map[string]float64 {
	result := make(map[string]float64)

	for _, meme := range memes {
		result[meme.Category] += meme.Views
	}

	return result
}

func FilterByComplexCondition(memes []BrainrotMeme) []string {
	var result []string

	for _, meme := range memes {
		if meme.TrendLevel >= 7 || (meme.Views > 50 && meme.Category == "Sigma") {
			result = append(result, meme.Name)
		}
	}

	return result
}

func main() {
	memes := []BrainrotMeme{
		{"Subo Dance", 8, "Subo Bratik", 42.5},
		{"TUN SAHUR LOOP", 9, "TUNTUNTUNSAHUR", 61.2},
		{"Sigma Walk", 6, "Sigma", 75.0},
		{"Skibidi Toilet XL", 10, "Skibidi", 120.3},
		{"Mewing Tutorial", 7, "Mewing", 38.9},
		{"Sigma Stare", 5, "Sigma", 55.4},
		{"Random Brainrot", 4, "Other", 20.1},
	}

	fmt.Println("Топ трендовых мемов (просмотры > 40):")
	topTrending := FindTopTrending(memes, 40)
	for _, meme := range topTrending {
		fmt.Printf("- %s | Trend: %d | Views: %.1fM\n",
			meme.Name, meme.TrendLevel, meme.Views)
	}

	fmt.Println("\nСуммарные просмотры по категориям:")
	categoryImpact := CalculateCategoryImpact(memes)
	for category, views := range categoryImpact {
		fmt.Printf("- %s: %.1fM просмотров\n", category, views)
	}

	fmt.Println("\nМемы, прошедшие сложный фильтр:")
	filtered := FilterByComplexCondition(memes)
	for _, name := range filtered {
		fmt.Println("-", name)
	}
}
