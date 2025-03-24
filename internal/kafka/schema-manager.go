package kafka

import (
	"codegen-service/internal/config"
	"encoding/json"
	"fmt"
	"github.com/linkedin/goavro/v2"
	"net/http"
	"sync"
)

var schemasForThisService = map[string]*goavro.Codec{
	"NewZip":        nil,
	"ProjectStatus": nil,
}

type SchemaManager struct {
	mu                sync.RWMutex
	schemas           map[string]*goavro.Codec
	schemaRegistryURL string
}

func NewSchemaManager(cfg *config.Config) *SchemaManager {
	manager := &SchemaManager{
		schemas:           schemasForThisService,
		schemaRegistryURL: cfg.SchemaRegistryUrl,
	}

	manager.loadSchemasFromRegistry()

	return manager
}

func (sm *SchemaManager) loadSchemasFromRegistry() {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	for topic := range sm.schemas {
		schemaData, err := sm.fetchSchemaFromRegistry(topic)
		if err != nil {
			panic(fmt.Sprintf("Failed to load schema for topic %s: %v", topic, err))
		}

		codec, err := goavro.NewCodec(schemaData)
		if err != nil {
			panic(fmt.Sprintf("Failed to create codec for topic %s: %v", topic, err))
		}

		sm.schemas[topic] = codec
		fmt.Printf("Schema for topic %s successfully loaded from registry\n", topic)
	}
}

func (sm *SchemaManager) fetchSchemaFromRegistry(topic string) (string, error) {
	schemaURL := fmt.Sprintf("%s/subjects/%s-value/versions/latest", sm.schemaRegistryURL, topic)
	resp, err := http.Get(schemaURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", err
	}

	var schemaResp map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&schemaResp); err != nil {
		return "", err
	}

	schema, ok := schemaResp["schema"].(string)
	if !ok {
		return "", err
	}

	return schema, nil
}

func (sm *SchemaManager) GetCodec(topic string) (*goavro.Codec, error) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	codec, exists := sm.schemas[topic]
	if !exists {
		return nil, fmt.Errorf("schema for topic %s not found", topic)
	}

	return codec, nil
}
