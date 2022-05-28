package golang_united_school_homework

import (
	"fmt"
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) == b.shapesCapacity {
		return fmt.Errorf("it goes out of the shapesCapacity range")
	}
	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i < 0 || i >= len(b.shapes) {
		return nil, fmt.Errorf("shape by index doesn't exist or index went out of the range")
	}
	shape := b.shapes[i]

	return shape, nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	shape, err := b.GetByIndex(i)
	if err != nil {
		return nil, fmt.Errorf("shape by index doesn't exist or index went out of the range")
	}
	b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)

	return shape, nil

}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	oldShape, err := b.GetByIndex(i)
	if err != nil {
		return nil, fmt.Errorf("shape by index doesn't exist or index went out of the range")
	}
	b.shapes[i] = shape

	return oldShape, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	result := 0.0
	for _, v := range b.shapes {
		result += v.CalcPerimeter()
	}

	return result
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	result := 0.0
	for _, v := range b.shapes {
		result += v.CalcArea()
	}

	return result
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	var filteredShapes []Shape
	err := fmt.Errorf("circles are not exist in the list")
	for _, v := range b.shapes {
		if sh, ok := v.(*Circle); ok {
			err = nil
		} else {
			filteredShapes = append(filteredShapes, sh)
		}
	}
	b.shapes = filteredShapes

	return err
}
