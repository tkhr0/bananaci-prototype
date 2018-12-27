package phaseType

type PhaseType int

const (
	Opened PhaseType = iota
	Labeled
	Unlabeled
	Pushed
	ForcePushed
	Closed
	Reopend
)

func (p PhaseType) String() string {
	switch p {
	case Opened:
		return "Opened"
	case Labeled:
		return "Labeled"
	case Unlabeled:
		return "Unlabeled"
	case Pushed:
		return "Pushed"
	case ForcePushed:
		return "ForcePushed"
	case Closed:
		return "Closed"
	case Reopend:
		return "Reopend"
	default:
		return "Unknown"
	}
}
