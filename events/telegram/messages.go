package telegram

//ToDO

const (
	msgHelp = `
📂 Link Keeper – твой цифровой сборник ссылок! 

Команды:
/start – начало
/help – помощь    
/rnd – случайная ссылка  

Все ваши ссылки – всегда под рукой!`

	msgHello = `
👋 Привет! Добро пожаловать в Link Keeper!` + "\n\n" + msgHelp

	msgUnknownCommand = `
❌ Неизвестная команда
Введите /help для списка доступных команд`

	msgNoSavedPages = `
📭 У вас пока нет сохранённых ссылок
Добавьте ссылку`

	msgSaved = `
✅ Ссылка успешно сохранена!
Теперь вы можете получить её через /rnd`

	msgAlreadyExists = `
⚠️ Эта ссылка уже сохранена!`
)
