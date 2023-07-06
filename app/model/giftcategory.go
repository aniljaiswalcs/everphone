package model

type InterestandCategory map[string]string     // 1-1 mapping
type Category []string                         //all category
type MapofCategoryInterest map[string][]string // category with similar interest

// all defined category which can be use further
// func getAllCategories() Category {
// 	allcategory := []string{"entertainment", "sports", "wellness", "nature", "cooking", "party", "fun", "outdoor activity", "anilmal",
// 		"family", "learning", "photography"}
// 	return allcategory
// }

var interestandCategories InterestandCategory
var mapofcategoryinterest MapofCategoryInterest

// map of interest and category
func CreateMapOfInterestandCategory() {
	interestandCategories = make(InterestandCategory, 50)
	interestandCategories = InterestandCategory{"music": "entertainment", "film": "entertainment", "pets": "anilmal", "scifi": "entertainment",
		"power lifting": "wellness", "triathlons": "wellness", "football": "sports", "crossfit": "sports", "handball": "sports", "running": "wellness",
		"techno": "entertainment", "drinking": "party", "botany": "learning", "scuba": "outdoor activity", "skydiving": "outdoor activity",
		"eating": "cooking", "family": "family",
		"german": "learning", "reading": "learning", "cocktails": "party", "yoga": "wellness", "sitting comfortably": "wellness",
		"photography": "photography",
	}
	createcategoryInterestMap()
}

func getCategoryfrominterest(interest string) string {

	if category, ok := interestandCategories[interest]; ok {
		return category
	}
	return ""
}

func getAllCategoriesfromInterests() InterestandCategory {
	return interestandCategories
}
func createcategoryInterestMap() {
	interestcategory := getAllCategoriesfromInterests()

	//map of category and slice of interest
	mapofcategoryinterest = make(map[string][]string, 50)

	for interest, category := range interestcategory {
		if mapofcategoryinterest[category] == nil && len(mapofcategoryinterest[category]) == 0 {
			mapofcategoryinterest[category] = append(mapofcategoryinterest[category], interest)
		} else {
			for index, val := range mapofcategoryinterest[category] {
				if val == interest {
					break
				} else if index == len(mapofcategoryinterest[category])-1 {
					mapofcategoryinterest[category] = append(mapofcategoryinterest[category], interest)
				}
			}
		}
	}
}

func GetSliceofSimilarInterestFromCategory(category string) []string {
	if len(mapofcategoryinterest) == 0 {
		return []string{}
	}
	return mapofcategoryinterest[category]
}
