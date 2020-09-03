package enum

type NodeRole int

const(
	Follower  NodeRole = 0
	Candidate NodeRole = 1
	Leader    NodeRole = 2
)
