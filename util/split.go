package util

import (
	"github.com/charmbracelet/lipgloss"
)

func SplitVertical(width int, heights []int, m ...string) string {
	// TODO
	if len(heights) != len(m) {
		return ""
	}

	s := lipgloss.NewStyle().Width(width).Height(heights[0]).Render(m[0])
	style := lipgloss.NewStyle().Width(width).Border(lipgloss.NormalBorder(), true, false, false, false)
	for i := 1; i < len(m); i++ {
		s = lipgloss.JoinVertical(0, s, style.Height(heights[i]).Render(m[i]))
	}

	return s
}

func SplitHorizontal(height int, widths []int, m ...string) string {
	// TODO
	if len(widths) != len(m) {
		return ""
	}

	s := lipgloss.NewStyle().Height(height).Width(widths[0]).Render(m[0])
	style := lipgloss.NewStyle().Height(height).Border(lipgloss.NormalBorder(), false, false, false, true)
	for i := 1; i < len(m); i++ {
		s = lipgloss.JoinHorizontal(0, s, style.Width(widths[i]).Render(m[i]))
	}

	return s
}
