package llm

import (
	"context"
	"encoding/json"

	"github.com/AndrewMysliuk/expath-data-generator/internal/schema"
)

type LLMClient interface {
	CallWithSchema(ctx context.Context, query string, schema []byte) (json.RawMessage, error)

	GetImmigration(ctx context.Context, country schema.CountryInfo) (*schema.ImmigrationInfo, error)
	GetTaxes(ctx context.Context, country schema.CountryInfo) (*schema.TaxInfo, error)
	GetFinance(ctx context.Context, country schema.CountryInfo) (*schema.FinanceInfo, error)
	GetCostOfLiving(ctx context.Context, country schema.CountryInfo) (*schema.CostOfLivingInfo, error)
	GetQualityOfLife(ctx context.Context, country schema.CountryInfo) (*schema.QualityOfLifeInfo, error)
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
