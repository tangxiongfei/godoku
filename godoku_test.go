package godoku

import (
	// "fmt"
	"testing"
)

const solvable88 string = `0 7 0 0 0 0 0 8 0
0 3 0 7 6 2 0 0 1
0 0 1 9 8 0 0 0 0
1 0 0 0 0 0 0 0 0
8 0 3 0 0 0 0 0 2
0 0 6 0 0 0 0 0 8
0 0 0 0 3 1 6 0 0
5 0 0 2 4 9 0 1 0
0 1 0 0 0 0 0 9 0`

const badRow string = `0 7 0 0 0 0 0 8 0
0 3 0 7 6 2 0 0 1
0 0 1 9 8 0 0 0 0
1 0 0 0 0 0 0 0
8 0 3 0 0 0 0 0 2
0 0 6 0 0 0 0 0 8
0 0 0 0 3 1 6 0 0
5 0 0 2 4 9 0 1 0
0 1 0 0 0 0 0 9 0`

const badCol string = `0 7 0 0 0 0 0 8 0
0 3 0 7 6 2 0 0 1
0 0 1 9 8 0 0 0 0
1 0 0 0 0 0 0 0 0
8 0 3 0 0 0 0 0 2
0 0 6 0 0 0 0 0 8
0 0 0 0 3 1 6 0 0
5 0 0 2 4 9 0 1 0`

const hard string = `0 0 0 0 0 0 0 0 0
0 0 1 0 0 3 0 0 0
8 0 0 0 0 0 0 0 6
0 0 0 2 0 9 0 1 0
6 4 0 0 0 0 9 0 0
0 0 0 0 0 0 0 0 0
0 0 2 0 0 0 7 3 0
7 0 0 6 4 0 0 0 0
0 0 0 8 0 0 0 0 0`

const invalidBoard string = `8 0 0 0 0 0 0 0 0
0 0 1 0 0 3 0 0 0
8 0 0 0 0 0 0 0 6
0 0 0 2 0 9 0 1 0
6 4 0 0 0 0 9 0 0
0 0 0 0 0 0 0 0 0
0 0 2 0 0 0 7 3 0
7 0 0 6 4 0 0 0 0
0 0 0 8 0 0 0 0 0`

const badFormatting = `0 7 0 0 0 0 0 8 0
0 3 0 7 6 2 0 0 1    
0 0 1 9 8 0 0 0 0 6 
1 0 0 0 0 0 0 0 0 
8 0 3 0 0 0 0 0 2 7
0 0 6 0 0 0 0 0 8
0 0 0 0 3 1 6 0 0
5 0 0 2 4 9 0 1 0
0 1 0 0 0 0 0 9 0
`

const noSolution string = `7 0 8 0 0 0 7 0 0
0 0 0 0 6 0 0 5 0
0 0 0 9 0 0 0 2 4
1 6 7 0 0 8 0 0 5
0 0 0 6 3 0 9 0 0
9 3 0 7 1 4 2 0 0
8 0 0 1 5 2 4 6 3
5 0 6 0 4 9 8 1 7
3 1 4 8 7 0 5 9 2`

// solution output
const solution string = `[2 7 9 1 5 3 4 8 6]
[4 3 8 7 6 2 9 5 1]
[6 5 1 9 8 4 3 2 7]
[1 4 5 3 2 8 7 6 9]
[8 9 3 6 1 7 5 4 2]
[7 2 6 4 9 5 1 3 8]
[9 8 2 5 3 1 6 7 4]
[5 6 7 2 4 9 8 1 3]
[3 1 4 8 7 6 2 9 5]
`

// solve only one of the 88 solutions
func TestSolve1(t *testing.T) {
	s, err := NewSudokuFromString(solvable88, 9)
	if err != nil {
		t.Error(err)
	}
	s.Solve()

	if s.GetSolutionsCount() != 1 {
		t.Errorf("Expected 1 != Actual %v", s.GetSolutionsCount())
	}
	if s.String() != solution {
		t.Errorf("Expected\n%v != Actual \n%v", solution, s)
	}
}

// make sure that the solver gets all 88 solutions
func TestSolve88(t *testing.T) {
	s, err := NewSudokuFromString(solvable88, 9)
	if err != nil {
		t.Error(err)
	}
	s.SolveAll()

	if s.GetSolutionsCount() != 88 {
		t.Errorf("Expected 88 != Actual %v", s.GetSolutionsCount())
	}
	// should be the first solution and not the 88th solution
	// - same as testsolve01
	if s.String() != solution {
		t.Errorf("Expected\n%v != Actual \n%v", solution, s)
	}
}

// test that the package is resillient against trailing spaces etc.
func TestBadFormatting(t *testing.T) {
	s, err := NewSudokuFromString(badFormatting, 9)
	if err != nil {
		t.Error(err)
	}
	s.Solve()

	if s.GetSolutionsCount() != 1 {
		t.Errorf("Expected 1 != Actual %v", s.GetSolutionsCount())
	}
}

// test that provided sudoku has at least Dim x Dim dimension
func TestBadDims(t *testing.T) {
	_, err := NewSudokuFromString(badRow, 9)
	if err == nil {
		t.Error(err)
	}
	_, err = NewSudokuFromString(badCol, 9)
	if err == nil {
		t.Error(err)
	}
}

// test that it reports correctly if the sudoky does 
// not have any solutions
func TestFail(t *testing.T) {
	// load sudoku
	s, err := NewSudokuFromString(noSolution, 9)
	if err != nil {
		t.Error(err)
	}
	s.Solve()

	if s.IsSolved() != false {
		t.Errorf("Expected: 'false' != Actual: %v", s.IsSolved())
	}
}

// ones the board has been read, validate it for 
// erroneously positioned starting values
func TestInvalidBoard(t *testing.T) {
	// test that invalid board is invalid
	s, err := NewSudokuFromString(invalidBoard, 9)
	if err != nil {
		t.Error(err)
	}
	if s.IsValidBoard() {
		t.Errorf("Expected: 'false' != Actual: %v", s.IsValidBoard())
	}
	// test that valid board is not invalid
	s, err = NewSudokuFromString(solvable88, 9)
	if err != nil {
		t.Error(err)
	}
	if !s.IsValidBoard() {
		t.Errorf("Expected: 'true' != Actual: %v", s.IsValidBoard())
	}
}

// bench the solving of the hardest sudoku
func BenchmarkSolveHard(b *testing.B) {
	b.StopTimer()
	s, err := NewSudokuFromString(hard, 9)
	if err != nil {
		b.Error(err)
	}
	b.StartTimer()

	s.Solve()
}

// bench how long it takes to get a "no solution" answer
func BenchmarkSolveFail(b *testing.B) {
	b.StopTimer()
	s, err := NewSudokuFromString(noSolution, 9)
	if err != nil {
		b.Error(err)
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		// s.Reset()
		s.Solve()
	}
}
