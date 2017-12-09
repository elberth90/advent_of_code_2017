package maze

// FindExit return information about number of steps required to leave a maze
func FindExit(l *List) int {
	current := l.Front()
	steps, jump := 0, 0
	for {
		if current == nil {
			return steps
		}
		jump = current.Value
		current.Value += 1
		current = current.JumpBy(jump)
		steps += 1
	}
}

// FindStrangeExit return information about number of steps required to leave a maze (strange jumps)
func FindStrangeExit(l *List) int {
	current := l.Front()
	steps, jump := 0, 0
	for {
		if current == nil {
			return steps
		}
		jump = current.Value
		if jump >= 3 {
			current.Value -= 1
		} else {
			current.Value += 1
		}
		current = current.JumpBy(jump)
		steps += 1
	}
}
