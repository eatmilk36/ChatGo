package Interface

type LogServiceInterface interface {
	LogError(value string)

	LogDebug(value string)

	LogInfo(value string)
}
