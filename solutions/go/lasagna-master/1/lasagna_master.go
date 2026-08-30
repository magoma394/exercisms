package lasagnamaster

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgTime int) int{
	if avgTime == 0 {
		avgTime = 2
	}
	return len(layers) * avgTime
}
// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
    noodles := 0
    sauce := 0.0

    for _, layer := range layers {
        if layer == "noodles" {
            noodles += 50
        } else if layer == "sauce" {
            sauce += 0.2
        }
    }

    return noodles, sauce
}
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
	SecretIngredient := friendsList[len(friendsList)-1]
	myList[len(myList)-1] = SecretIngredient
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amounts []float64, portions int) []float64{
	scaled := make([]float64, len(amounts))

	for i, amount := range amounts {
		scaled[i] = (amount / 2.0) * float64(portions)
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
