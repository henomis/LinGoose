package cohereembedder

import (
	"context"
	"os"

	coherego "github.com/henomis/cohere-go"
	"github.com/henomis/cohere-go/model"
	"github.com/henomis/cohere-go/request"
	"github.com/henomis/cohere-go/response"
	"github.com/henomis/lingoose/embedder"
)

type EmbedderModel = model.EmbedModel

const (
	defaultEmbedderModel EmbedderModel = model.EmbedModelEnglishV20

	EmbedderModelEnglishV20      EmbedderModel = model.EmbedModelEnglishV20
	EmbedderModelEnglishLightV20 EmbedderModel = model.EmbedModelEnglishLightV20
	EmbedderModelMultilingualV20 EmbedderModel = model.EmbedModelMultilingualV20
)

var EmbedderModelsSize = map[EmbedderModel]int{
	EmbedderModelEnglishLightV20: 1024,
	EmbedderModelEnglishV20:      4096,
	EmbedderModelMultilingualV20: 768,
}

type Embedder struct {
	model  EmbedderModel
	client *coherego.Client
}

func New() *Embedder {
	return &Embedder{
		client: coherego.New(os.Getenv("COHERE_API_KEY")),
		model:  defaultEmbedderModel,
	}
}

func (e *Embedder) WithAPIKey(apiKey string) *Embedder {
	e.client = coherego.New(apiKey)
	return e
}

func (e *Embedder) WithModel(model EmbedderModel) *Embedder {
	e.model = model
	return e
}

func (h *Embedder) Embed(ctx context.Context, texts []string) ([]embedder.Embedding, error) {
	resp := &response.Embed{}
	err := h.client.Embed(
		ctx,
		&request.Embed{
			Texts: texts,
			Model: &h.model,
		},
		resp,
	)
	if err != nil {
		return nil, err
	}

	embeddings := make([]embedder.Embedding, len(resp.Embeddings))

	for i, embedding := range resp.Embeddings {
		embeddings[i] = embedding
	}
	return embeddings, nil
}
