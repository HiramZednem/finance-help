package services

type TelegramServiceImpl struct {
	BaseMessageService
}

func NewTelegramServiceImpl() MessageServiceInterface {
	return &TelegramServiceImpl{}
}

func (t *TelegramServiceImpl) SendMessage(text string) {

}
