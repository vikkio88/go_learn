package lasagna

const (
	MIN_TIME_PER_LAYER    = 2
	PASTA_QUANTITY        = 50
	SAUCE_QUANTITY        = 0.2
	SCALE_STANDARD_FACTOR = 2
)

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgPerLayer int) int {
	if avgPerLayer == 0 {
		avgPerLayer = MIN_TIME_PER_LAYER
	}
	return len(layers) * avgPerLayer
}

// TODO: define the 'Quantities()' function

// this return param names stuff is so cool!
func Quantities(layers []string) (pasta int, sauce float64) {

	for _, layer := range layers {
		if layer == "noodles" {
			pasta += PASTA_QUANTITY
			continue
		}

		if layer == "sauce" {
			sauce += SAUCE_QUANTITY
		}

	}

	return
}

// TODO: define the 'AddSecretIngredient()' function
// func AddSecretIngredient(friendStuff []string, myStuff *[]string) {
func AddSecretIngredient(friendStuff []string, myStuff []string) {
	lastIndex := len(friendStuff) - 1
	secret := friendStuff[lastIndex]

	// (*myStuff)[len(*myStuff)-1] = secret
	myStuff[len(myStuff)-1] = secret
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, scale int) []float64 {
	scaled := make([]float64, len(quantities))
	for i, q := range quantities {
		scaled[i] = q * (float64(scale) / float64(SCALE_STANDARD_FACTOR))
	}

	return scaled
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
