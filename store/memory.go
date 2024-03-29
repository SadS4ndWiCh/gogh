package store

import "context"

type MemoryStore struct{
    data map[string]interface{}
}

func NewMemoryStore() *MemoryStore {
    return &MemoryStore{
        data: make(map[string]interface{}),
    }
}

func (st *MemoryStore) Get(ctx context.Context, key string) (interface{}, error) {
    value, exists := st.data[key]
    if !exists {
        return nil, ERR_NOT_FOUND
    }

    return value, nil
}

func (st *MemoryStore) Set(ctx context.Context, key string, value interface{}) error {
    st.data[key] = value

    return nil
}
