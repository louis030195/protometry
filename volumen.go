package protometry


// Volume is a 3-d interface representing volumes like Boxes, Spheres, Capsules ...
type Volume interface {
	Fit(Volume) bool
	Intersects(Volume) bool
	// Average create a new volume averaged on 2 volumes
	Average(Volume) Volume
	// Mutate create a new volume with random mutations
	Mutate(float64) Volume
}