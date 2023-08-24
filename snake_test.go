package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSnakeUp(t *testing.T) {

	snake_before_moving := snake{
		direction: direction("up"),
		body: []position{
			{x: 2, y: 4},
		},
	}

	snake_after_moving := move(snake_before_moving)
	snake_after_2nd_moving := move(snake_after_moving)

	expected := snake{
		direction: direction("up"),
		body: []position{
			{x: 2, y: 3},
		},
	}

	expected_snake_before_moving := snake{
		direction: direction("up"),
		body:      []position{{x: 2, y: 4}},
	}

	expected_snake_after_2nd_moving := snake{
		direction: direction("up"),
		body:      []position{{x: 2, y: 2}},
	}

	require.Equal(t, expected, snake_after_moving)
	require.Equal(t, expected_snake_before_moving, snake_before_moving)
	require.Equal(t, expected_snake_after_2nd_moving, snake_after_2nd_moving)

}

func TestSnakeDown(t *testing.T) {

	snake_before_moving := snake{
		direction: direction("down"),
		body: []position{
			{x: 2, y: 4},
		},
	}

	snake_after_1stmove := move(snake_before_moving)

	expected_snake_before_moving := snake{
		direction: direction("down"),
		body:      []position{{x: 2, y: 4}},
	}

	expected_snake_after_1stmove := snake{
		direction: direction("down"),
		body:      []position{{x: 2, y: 5}},
	}

	require.Equal(t, expected_snake_before_moving, snake_before_moving)
	require.Equal(t, expected_snake_after_1stmove, snake_after_1stmove)

}

func TestSnakeRight(t *testing.T) {

	snake_before_moving := snake{
		direction: direction("right"),
		body: []position{
			{x: 2, y: 4},
		},
	}

	snake_after_1stmove := move(snake_before_moving)

	expected_snake_before_moving := snake{
		direction: direction("right"),
		body:      []position{{x: 2, y: 4}},
	}

	expected_snake_after_1stmove := snake{
		direction: direction("right"),
		body:      []position{{x: 3, y: 4}},
	}

	require.Equal(t, expected_snake_before_moving, snake_before_moving)
	require.Equal(t, expected_snake_after_1stmove, snake_after_1stmove)
}

func TestSnakeLeft(t *testing.T) {

	snake_before_moving := snake{
		direction: direction("left"),
		body: []position{
			{x: 2, y: 4},
		},
	}

	snake_after_1stmove := move(snake_before_moving)

	expected_snake_before_moving := snake{
		direction: direction("left"),
		body:      []position{{x: 2, y: 4}},
	}

	expected_snake_after_1stmove := snake{
		direction: direction("left"),
		body:      []position{{x: 1, y: 4}},
	}

	require.Equal(t, expected_snake_before_moving, snake_before_moving)
	require.Equal(t, expected_snake_after_1stmove, snake_after_1stmove)
}