package llm

import (
	"context"
)

type LLMClient interface {
	SearchAndSummarize(ctx context.Context, query string) (string, error)
}

// Инструменты агента (built-in tools):
// - web_search → поиск в интернете, возвращает свежие данные + ссылки.
// - file_search → поиск по твоим загруженным документам (векторный поиск встроен).
// - code_interpreter → выполнение Python-кода в песочнице (математика, графики, анализ данных).
// - image_generation → генерация изображений (DALL·E 3).
// - computer → экспериментальный доступ к управлению компьютером (аналог ChatGPT «Computer Use»).
// - local_shell → выполнение команд в локальной оболочке (Linux/macOS).

// Аудио (не «tools», но тоже в Responses API):
// - audio.transcriptions (Whisper) → из речи в текст.
// - audio.speech (TTS) → из текста в речь.
