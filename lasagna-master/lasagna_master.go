package lasagna

const (
	defaultTimePerLayer    = 2
	noodleQuantPerLayer    = 50
	sauceQuantPerLayer     = 0.2
	defaulServingPerRecipe = 2.0
)

// PreparationTime computes the amount of time required from preparation
// based on the number of layers and the average time per layer.
func PreparationTime(layers []string, avgTime int) int {
	if avgTime == 0 {
		return defaultTimePerLayer * len(layers)
	} else {
		return avgTime * len(layers)
	}
}

// Quantities computes the quantity of noodles and sauce needed based on
// a list of layers.
func Quantities(layers []string) (noodles int, sauce float64) {
	for _, layer := range layers {
		if layer == "noodles" {
			noodles += noodleQuantPerLayer
		} else if layer == "sauce" {
			sauce += sauceQuantPerLayer
		}
	}
	return
}

// AddSecretIngredient takes the secret ingredient from a frind's list
// and adds it to your own list.
func AddSecretIngredient(friendList, ownList []string) {
	ownList[len(ownList)-1] = friendList[len(friendList)-1]
}

// ScaleRecipe takes the quantities of ingredients for two portions and
// scales them for n portions.
func ScaleRecipe(portionList []float64, numberPortions int) (scaledList []float64) {
	scaleFactor := float64(numberPortions) / defaulServingPerRecipe
	scaledList = make([]float64, len(portionList))
	for i, v := range portionList {
		scaledList[i] = v * scaleFactor
	}
	return
}
