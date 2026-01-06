package state

const (
	AntLength           = 30
	AntWidth            = 10
	NumberOfAnts        = 150
	NumberOfFoodSources = 10
	NumberOfObstacles   = 4
	ObstacleMaxLength   = 400
	HomeRadius          = 50.0
	MaxFoodSourceRadius = 40.0

	CameraSpeed = 5.0

	PheromoneDecay       = 0.99
	PheromoneCap         = 10.0
	DiffusionStrength    = 5.0
	InitialScentStrength = 10.0

	MaxFood = 100.0

	AntSpeed         = 2
	AntTurnSpeed     = 0.15
	SensorAngle      = 0.4
	SensorDist       = 35
	SensorThreshold  = 0.05
	DepositStrength  = 5
	ScentDecay       = 0.995
	MovementFoodCost = 0.005
)
