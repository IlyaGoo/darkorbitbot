package algorithm

type state struct {
	name string
}

var notInitedState state = state{"не инициализирован"}
var inMenuState state = state{"В меню"}
var loadingState state = state{"Загрузка"}
var connectingState state = state{"Подключение"}
var interfaceOptionState state = state{"Настройка интерфейса"}
var inGameState state = state{"В игре"}
